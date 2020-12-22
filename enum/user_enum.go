package enum

type ResponseType int

//HTTP 200是請求成功，但不代表業務一定會成功，所以回傳狀態碼
const (
	ResOK   ResponseType = 200
	ResFail ResponseType = 500
)

func (p ResponseType) String() string {
	switch p {
	case ResOK:
		return "Response Ok"
	case ResFail:
		return "Response Failed"
	default:
		return "Unkown"
	}
}
