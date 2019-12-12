package producer_test

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama/mocks"
	"github.com/naufalziyad/Go-Kafka/producer"
)

func TestKafkaProducer_SendMessage(t *testing.T) {
	t.Run("send ok message", func(t *testing.T) {
		mockedProducer := mocks.NewSyncProducer(t, nil)

		mockedProducer.ExpectSendMessageAndSucceed()
		kafka := &producer.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("Test_Topic", msg)
		if err != nil {
			t.Errorf("Send message should not be error but have: %v", err)
		}

	})

	t.Run("send message not ok", func(t *testing.T) {
		mockerProducer := mocks.NewSyncProducer(t, nil)

		mockerProducer.ExpectSendMessageAndFail(fmt.Errorf("Error nih"))
		kafka := &producer.KafkaProducer{
			Producer: mockerProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("Test_Topic", msg)
		if err == nil {
			t.Error("This should be error")
		}
	})

}
