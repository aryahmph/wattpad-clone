package main

import (
	"github.com/aryahmph/wattpad-clone/src/ishigami/common/env"
	mailCommon "github.com/aryahmph/wattpad-clone/src/ishigami/common/mail"
	nsqConsumer "github.com/aryahmph/wattpad-clone/src/ishigami/common/nsq/consumer"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := env.LoadConfig()
	mailManager := mailCommon.NewMailManager(cfg.MailEmail, cfg.MailPassword, cfg.MailSMTPHost, cfg.MailSMTPPort)
	consumer, err := nsqConsumer.NewConsumer(cfg.NSQTopic, cfg.NSQChannel)
	if err != nil {
		log.Fatalln("failed to create nsq consumer:", err)
	}
	mailConsumerMessageHandler := nsqConsumer.NewMailMessageHandler(mailManager)
	consumer.AddHandler(mailConsumerMessageHandler)
	err = consumer.ConnectToNSQLookupd(cfg.NSQLookupAddr)
	if err != nil {
		log.Fatalln("failed to connect to nsqlookupd:", err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
