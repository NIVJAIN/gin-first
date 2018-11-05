package system

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if strings.Contains(path, "swagger") {
			return
		}
		token := context.Request.Header.Get("ACCESS_TOKEN")
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": "请求未携带token,无访问权限！",
			})
			context.Abort()
			return
		}
		j := NewJWT()
		// 解析token包含的信息
		claims, err := j.ResolveToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": err.Error(),
			})
			context.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		context.Set("claims", claims)
	}
}

// jwt签名结构
type JWT struct {
	SigningKey []byte
}

// 定义一些常量
var (
	TokenExpired     error  = errors.New("Token 已经过期")
	TokenNotValidYet error  = errors.New("Token 尚未激活")
	TokenMalformed   error  = errors.New("Token 格式错误")
	TokenInvalid     error  = errors.New("Token 无法解析")
	SignKey          string = "82040620FEFAC4511FC65000ADAB0F77"
)

// 载荷，加一些系统需要的信息
type CustomClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// 新建一个 jwt 实例
func NewJWT() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// 获取 signKey
func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// 生成 tokenConfig
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 tokenConfig
func (j *JWT) ResolveToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
