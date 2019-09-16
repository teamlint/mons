package event

import (
	"time"

	"github.com/teamlint/mons/shared/lib"
)

const (
	ClusterID = "test-cluster"
)

// Event 事件结构
type Event struct {
	ID            string    `json:"id"`             // 事件唯一ID
	AggregateID   string    `json:"aggregate_id"`   // 聚合ID, 相关业务聚合根ID
	AggregateType string    `json:"aggregate_type"` // 聚合类型
	Subject       string    `json:"subject"`        // 事件主题
	Data          []byte    `json:"data"`           // 事件数据
	CreatedAt     time.Time `json:"created_at"`     // 事件发生时间
}

func New(aggID string, aggType string, subject string, data []byte) Event {
	return Event{
		ID:            lib.UUID(),
		AggregateID:   aggID,
		AggregateType: aggType,
		Subject:       subject,
		Data:          data,
		CreatedAt:     time.Now(),
	}
}
