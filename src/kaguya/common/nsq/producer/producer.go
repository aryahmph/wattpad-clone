package producer

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func NewProducer(addr string) *nsq.Producer {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		log.Fatalln("failed to create nsq producer:", err)
	}
	err = producer.Ping()
	if err != nil {
		log.Fatalln("Could not connect to nsqd:", err)
	}
	return producer
}
