package res

//回傳實體
type Entity struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
	Data      interface{} `json:"data"`
}
