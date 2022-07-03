package mail

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MailProducer struct {
	producer *nsq.Producer
	topic    string
}

func NewMailProducer(producer *nsq.Producer, topic string) *MailProducer {
	return &MailProducer{producer: producer, topic: topic}
}

func (p MailProducer) Publish(mail Receiver) error {
	payload, err := json.Marshal(mail)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	err = p.producer.Publish(p.topic, payload)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
