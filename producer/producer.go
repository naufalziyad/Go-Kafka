package producer

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func (p *KafkaProducer) SendMessage(topic, msg string) error {
	fmt.Println("coba saja")
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	paritition, offset, err := p.Producer.SendMessage(kafkaMessage)
	if err != nil {
		logrus.Errorf("Send Message Error : %v", err)
		return err

	}

	logrus.Infof("Send Message Success, Topic %v, Paritition %v, Offset %d", topic, paritition, offset)
	return nil
}
