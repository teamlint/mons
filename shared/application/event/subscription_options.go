package event

import (
	"time"
)

// Enum for start position type.
type StartPosition int32

const (
	StartPosition_NewOnly        StartPosition = 0
	StartPosition_LastReceived   StartPosition = 1
	StartPosition_TimeDeltaStart StartPosition = 2
	StartPosition_SequenceStart  StartPosition = 3
	StartPosition_First          StartPosition = 4
)

// SubscriptionOptions 订阅选项
type SubscriptionOptions struct {
	MaxInflight   int           // 最大未确认消息堆积数量
	AckWait       time.Duration // 等待最大消息确认时间
	StartAt       StartPosition // 消息开始位置
	StartSequence uint64        // 消息开始序列号
	StartTime     time.Time     // 消息开始时间
}
