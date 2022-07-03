package consumer

import (
	"github.com/nsqio/go-nsq"
	"time"
)

func NewConsumer(topic, channel string) (*nsq.Consumer, error) {
	nsqCfg := nsq.NewConfig()
	nsqCfg.MaxRequeueDelay = time.Second * 900
	nsqCfg.DefaultRequeueDelay = time.Second * 0

	consumer, err := nsq.NewConsumer(topic, channel, nsqCfg)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
