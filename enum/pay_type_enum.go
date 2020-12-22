package enum

type PayType int

const (
	Bank       PayType = 0
	CreditCard PayType = 1
	PayPal     PayType = 2
)

func (p PayType) String() string {
	switch p {
	case Bank:
		return "銀行支付"
	case CreditCard:
		return "信用卡支付"
	case PayPal:
		return "Paypal"
	default:
		return "Known Pay"
	}
}
