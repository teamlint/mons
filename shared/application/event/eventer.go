package event

// EventHandler 事件处理器
type EventHandler func(Event) error

// Eventer 事件处理接口
type Eventer interface {
	Publish(e Event) error                                                                                           // 发布事件
	Subscribe(subject string, durableID string, hdl EventHandler, opts ...SubscriptionOptions) (Subscription, error) // 订阅事件: 事件主题, 事件处理器,订单客户端持久化ID
	Close() error                                                                                                    // 关闭
}
