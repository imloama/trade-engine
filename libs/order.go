package libs

// LimitOrder 表示为限价单
const LimitOrder int = 0

// MARKET_ORDER表示 市价单
const MARKET_ORDER int = 1

// BUY_ORDER 表示买单
const BUY_ORDER = 0

// SELL_ORDER 表示卖单
const SELL_ORDER = 1

type Order struct {
	// 订单类型
	t int
	// 订单方向,买单还是卖间
	d int
	// 订单id
	id int64

}
