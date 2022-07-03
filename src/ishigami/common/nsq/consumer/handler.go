package consumer

import (
	"encoding/json"
	mailCommon "github.com/aryahmph/wattpad-clone/src/ishigami/common/mail"
	"github.com/nsqio/go-nsq"
)

type mailMessageHandler struct {
	mailManager *mailCommon.MailManager
}

func NewMailMessageHandler(mailManager *mailCommon.MailManager) *mailMessageHandler {
	return &mailMessageHandler{mailManager: mailManager}
}

func (m *mailMessageHandler) HandleMessage(message *nsq.Message) error {
	var request mailCommon.Receiver
	err := json.Unmarshal(message.Body, &request)
	if err != nil {
		return err
	}

	err = m.mailManager.SendMail(request)
	if err != nil {
		return err
	}
	return nil
}
