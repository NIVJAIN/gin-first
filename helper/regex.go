package helper

import "regexp"

// Regular expression patterns
const (

	/** 匹配日期 */
	DateRegex = `(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))-02-29)`

	/** 匹配时间 */
	TimeRegex = `(?i)\d{1,2}:\d{2} ?(?:[ap]\.?m\.?)?|\d[ap]\.?m\.?`

	/** 匹配电话号码 */
	PhoneRegex = `^((1[3,5,8][0-9])|(14[5,7])|(17[0,6,7,8])|(19[6,7,8,9]))\\d{8}$`

	/** 匹配网络连接地址 */
	UrlRegex = `^((ht|f)tps?):\\/\\/[\\w\\-]+(\\.[\\w\\-]+)+([\\w\\-\\.,@?^=%&:\\/~\\+#]*[\\w\\-\\@?^=%&\\/~\\+#])?$`

	/** 匹配邮箱地址 */
	EmailRegex = `^[a-zA-Z0-9_.-]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$`

	/** 匹配Ip4 */
	IPv4Regex = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`

	/** 匹配Ip6 */
	Ipv6Regex = `(?:(?:(?:[0-9A-Fa-f]{1,4}:){7}(?:[0-9A-Fa-f]{1,4}|:))|(?:(?:[0-9A-Fa-f]{1,4}:){6}(?::[0-9A-Fa-f]{1,4}|(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){5}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,2})|:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){4}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,3})|(?:(?::[0-9A-Fa-f]{1,4})?:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){3}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,4})|(?:(?::[0-9A-Fa-f]{1,4}){0,2}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){2}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,5})|(?:(?::[0-9A-Fa-f]{1,4}){0,3}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){1}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,6})|(?:(?::[0-9A-Fa-f]{1,4}){0,4}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?::(?:(?:(?::[0-9A-Fa-f]{1,4}){1,7})|(?:(?::[0-9A-Fa-f]{1,4}){0,5}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(?:%.+)?\s*`

	/** 匹配Ip地址 */
	IPRegex = IPv4Regex + `|` + Ipv6Regex

	/** 匹配端口 */
	PortRegex = `6[0-5]{2}[0-3][0-5]|[1-5][\d]{4}|[2-9][\d]{3}|1[1-9][\d]{2}|10[3-9][\d]|102[4-9]`

	/** 身份证号 */
	IDCardRegex = `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`

	/** 匹配Mac地址 */
	MacRegex = `(([a-fA-F0-9]{2}[:-]){5}([a-fA-F0-9]{2}))`

	/** 匹配带表单符号的 描述信息 */
	DescRegex = `^[，|。|！|：|“|”|？|；|《|》|（|）|\u4e00-\u9fa5A-Za-z0-9_/./,]+$`

	/** 匹配中文 */
	ChineseRegex = `^[\u4e00-\u9fa5]+$`

	/** 匹配强密码   字母+数字+特殊字符 */
	StrongPassword = `^(?![a-zA-z]+$)(?!\d+$)(?![!@#$%^&*]+$)(?![a-zA-z\d]+$)(?![a-zA-z!@#$%^&*]+$)(?![\d!@#$%^&*]+$)[a-zA-Z\d!@#$%^&*]+$`

	/** 匹配中度密码 字母+数字，字母+特殊字符，数字+特殊字符*/
	MediumPassword = `^(?![a-zA-z]+$)(?!\d+$)(?![!@#$%^&*]+$)[a-zA-Z\d!@#$%^&*]+$`

	/** 匹配数字字母 */
	LetNumRegex = `^[A-Za-z0-9]+$`
)

// 匹配是否为日期格式
func IsDate(text string) bool {
	matched, _ := regexp.MatchString(DateRegex, text)
	return matched
}

// 匹配是否是时间格式
func IsTime(text string) bool {
	matched, _ := regexp.MatchString(TimeRegex, text)
	return matched
}

// 匹配是否是电话号码
func IsPhone(text string) bool {
	matched, _ := regexp.MatchString(PhoneRegex, text)
	return matched
}

// 匹配网络连接地址
func IsUrl(text string) bool {
	matched, _ := regexp.MatchString(UrlRegex, text)
	return matched
}

// 匹配邮箱地址
func IsEmail(text string) bool {
	matched, _ := regexp.MatchString(EmailRegex, text)
	return matched
}

// 匹配Ip
func IsIPAddress(text string) bool {
	matched, _ := regexp.MatchString(IPRegex, text)
	return matched
}

// 匹配 mac 地址
func IsMacAddress(text string) bool {
	matched, _ := regexp.MatchString(MacRegex, text)
	return matched
}

// 匹配端口
func IsPort(text string) bool {
	matched, _ := regexp.MatchString(PortRegex, text)
	return matched
}

// 匹配身份证号码
func IsIDCard(text string) bool {
	matched, _ := regexp.MatchString(IDCardRegex, text)
	return matched
}

// 匹配带最大，最小长度限制的描述信息
func DescMatchMinAndMax(text string, min int, max int) bool {
	if min > max {
		return false
	}
	if len(text) >= min && len(text) <= max {
		matched, _ := regexp.MatchString(DescRegex, text)
		return matched
	}
	return false
}

// 匹配带最大长度的
func DescMatchMax(text string, max int) bool {
	if max < 0 {
		return false
	}
	if len(text) <= max {
		matched, _ := regexp.MatchString(DescRegex, text)
		return matched
	}
	return false
}

// 匹配带最大，最小长度的数字字母
func MatchLetterNumMinAndMax(text string, min int, max int) bool {
	if min > max {
		return false
	}
	if len(text) >= min && len(text) <= max {
		matched, _ := regexp.MatchString(LetNumRegex, text)
		return matched
	}
	return false
}

// 匹配带最大长度的数字字母
func MatchLetterNumMax(text string, max int) bool {
	if max < 0 {
		return false
	}
	if len(text) <= max {
		matched, _ := regexp.MatchString(LetNumRegex, text)
		return matched
	}
	return false
}

// 匹配带最大最小长度的中文
func MatchChineseMinAndMax(text string, min int, max int) bool {
	if min > max {
		return false
	}
	if len(text) >= min && len(text) <= max {
		matched, _ := regexp.MatchString(ChineseRegex, text)
		return matched
	}
	return false
}

// 匹配带最大长度的中文
func MatchChineseMax(text string, max int) bool {
	if max < 0 {
		return false
	}
	if len(text) <= max {
		matched, _ := regexp.MatchString(ChineseRegex, text)
		return matched
	}
	return false
}

// 强密码
func MatchStrongPassword(text string, min int, max int) bool {
	if min > max {
		return false
	}
	if len(text) >= min && len(text) <= max {
		matched, _ := regexp.MatchString(StrongPassword, text)
		return matched
	}
	return false
}

// 匹配中度密码
func MatchMediumPassword(text string, min int, max int) bool {
	if min > max {
		return false
	}
	if len(text) >= min && len(text) <= max {
		matched, _ := regexp.MatchString(MediumPassword, text)
		return matched
	}
	return false
}
