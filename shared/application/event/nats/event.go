package nats

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/teamlint/mons/shared/application/event"

	"github.com/nats-io/stan.go"
	"github.com/teamlint/golog"
)

var _ = event.Eventer(&NATSEventer{})

// NATSEventer
type NATSEventer struct {
	conn      stan.Conn
	clusterID string
	clientID  string
}

func NewNATSEventer(conf *Config) event.Eventer {
	e := NATSEventer{}
	var opts []stan.Option
	// cluster id
	if conf.ClusterID != "" {
		e.clusterID = conf.ClusterID
	}
	// client id
	hostname, _ := os.Hostname()
	// pid := os.Getpid()
	clientID := fmt.Sprintf("%v_%v", conf.ClientID, hostname)
	e.clientID = clientID
	// }
	golog.Infof("[NATSEventer] client id: %v", e.clientID)
	// url
	if conf.URL != "" {
		opts = append(opts, stan.NatsURL(conf.URL))
	} else {
		golog.Warn("[NATSEventer] nats no url setting")
	}
	// conn
	conn, err := stan.Connect(e.clusterID, e.clientID, opts...)
	if err != nil {
		golog.Fatalf("nats streaming connect err: %v", err)
	}
	e.conn = conn
	runtime.SetFinalizer(&e, close)
	return &e
}

func (e *NATSEventer) Publish(evt event.Event) error {
	if e == nil {
		golog.Error("[Publish] NATSEventer is nil")
		return nil
	}
	data, err := json.Marshal(evt)
	if err != nil {
		return nil
	}
	return e.conn.Publish(evt.Subject, data)

}

func (e *NATSEventer) Subscribe(subject string, durableID string, hdl event.EventHandler, opts ...event.SubscriptionOptions) (event.Subscription, error) {
	if e == nil {
		golog.Error("[Subscribe] NATSEventer is nil")
		return nil, nil
	}
	isDurabled := true
	//options
	options := []stan.SubscriptionOption{
		stan.SetManualAckMode(),
	}
	if len(opts) > 0 {
		opt := opts[0]
		if opt.MaxInflight > 0 {
			golog.Infof("[Subscribe] MaxInflight: %v", opt.MaxInflight)
			options = append(options, stan.MaxInflight(opt.MaxInflight))
		}
		if opt.AckWait > 0 {
			golog.Infof("[Subscribe] AckWait: %v", opt.AckWait)
			options = append(options, stan.AckWait(opt.AckWait))
		}
		if opt.StartSequence > 0 {
			golog.Infof("[Subscribe] StartSequence: %v", opt.StartSequence)
			isDurabled = false
			options = append(options, stan.StartAtSequence(opt.StartSequence))
		}
		if !opt.StartTime.IsZero() {
			golog.Infof("[Subscribe] StartTime: %v", opt.StartTime)
			isDurabled = false
			options = append(options, stan.StartAtTime(opt.StartTime))
		}
	}
	// durable
	if isDurabled && durableID != "" {
		options = append(options, stan.DurableName(durableID))
	}
	// subscribe
	sub, err := e.conn.QueueSubscribe(subject, durableID, msgHandler(hdl), options...)
	if err != nil {
		golog.Errorf("[NATSEventer] Subscribe error: %v", err)
	}
	return sub, err
}

func (e *NATSEventer) Close() error {
	return close(e)
}

func close(e *NATSEventer) error {
	if e == nil {
		golog.Error("[Close] NATSEventer is nil")
		return nil
	}
	golog.Info("nats streaming connect closed")
	return e.conn.Close()
}

// msg handler
func msgHandler(hdl event.EventHandler) stan.MsgHandler {
	return func(msg *stan.Msg) {
		evt := event.Event{}
		err := json.Unmarshal(msg.Data, &evt)
		if err != nil {
			golog.Errorf("[Subscribe] unmarshal error: %v", err)
			return
		}
		err = hdl(evt)
		if err != nil {
			golog.Errorf("[Subscribe] event handler error: %v", err)
			return
		}
		msg.Ack()
	}
}
