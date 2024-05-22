package config

// RMQ exchange
const (
	RabbitEventsExchange       = "e.events.forward"
	RabbitOrderExchange        = "e.orders.forward"
	RabbitExchangeRateExchange = "e.exchange.forward"
	RabbitBalanceExchange      = "e.balances.forward"
)

// RMQ rk
const (
	GetExchangeRateRequestRoutingKey  = "r.quote-manager.rate.#.GetExchangeRateRequest"
	GetExchangeRateResponseRoutingKey = "r.quote-manager.rate.GetExchangeRateResponse"

	UpdatedOrderEventRoutingKey = "r.event.order.OrderUpdateEvent"
)

// Queue name
const (
	UpdatedOrderEventQueueName  = "q.quote-manager.order.event"
	CreateOrderRequestQueueName = "q.quote-manager.exchange.rate.get"
)

// Errors
const (
	ErrLoadConfig = "Error load configuration"
	ErrParseLog   = "Error parse log level"
	ErrConnectDb  = "Error connect db"
)

// Url pattern
const (
	RmqUrlConnectionPattern = "amqp://%s:%s@%s:%d/"
	RedisConnectionPattern  = "%s:%d"
)
