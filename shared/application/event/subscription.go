package event

// Subscription 订阅者
type Subscription interface {
	Unsubscribe() error // 取消订阅,持久化订阅也被删除,客户端重启后持久化订阅也不能消费消息
	Close() error       // 从服务端移除订阅,持久化订阅不删除,如果客户端已经连接,调用此方法将会出错
}
