package nats

// Config NATS配置
type Config struct {
	ClusterID   string // 集群ID
	ClientID    string // 客户端ID, 只能包含字母,-,_字符
	URL         string // nats 连接url
	MaxInflight int    // 未确认消息最大数量,默认1024
	AckWait     string // 消息确认最大超时时间,超过时间重复发送消息,默认30s
}
