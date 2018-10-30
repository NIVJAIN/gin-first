package helper

type JsonObject struct {

	/** 状态码 */
	Code string

	/** 内容体 */
	Content interface{}

	/** 消息 */
	Message string

}

type PageBean struct {

	/** 当前页  */
	Page      int

	/** 每页显示的最大行数 */
	PageSize  int

	/** 总记录数 */
	Total     int

	/** 每行的数据 */
	Rows      interface{}
}


