package enum

//僅給自己看訂單狀態，不可寫入API中
type PayStatus int

const (
	Unpay PayStatus = 0
	Payed PayStatus = 1
)

func (p PayStatus) String() string {
	switch p {
	case Unpay:
		return "尚未支付"
	case Payed:
		return "已支付"
	default:
		return "Unknown Status"
	}
}
