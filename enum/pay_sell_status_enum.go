package enum

type SellStatus int

const (
	Selling  SellStatus = 0
	StopSell SellStatus = 1
)

func (p SellStatus) String() string {
	switch p {
	case Selling:
		return "銷售中"
	case StopSell:
		return "停止銷售"
	default:
		return "Unkown Sell Status"
	}
}
