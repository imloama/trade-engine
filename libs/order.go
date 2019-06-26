package libs

// LimitOrder 表示为限价单
const LimitOrder uint8 = 0

// MARKET_ORDER表示 市价单
const MARKET_ORDER uint8 = 1

// BUY_ORDER 表示买单
const BUY_ORDER uint8 = 0

// SELL_ORDER 表示卖单
const SELL_ORDER uint8 = 1

type Order struct {
	// 订单类型,市价单，限价单
	t uint8
	// 订单方向,买单还是卖单，
	d uint8
	// 订单id
	id int64
	// 价格
	p float32
	// 数量，未成交数量
	m float32
}
