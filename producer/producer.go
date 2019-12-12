package producer

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func (p *KafkaProducer) SendMessage(topic, msg string) error {
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMessage)
	if err != nil {
		logrus.Errorf("Send Message Error : %v", err)
		return err

	}

	logrus.Infof("Send Message Success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}
