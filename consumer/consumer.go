package consumer

import (
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	message, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)

	if err != nil {
		logrus.Errorf("Unable to consume partition %v got error %v", partition, err)
		return
	}

	defer func() {
		if err := message.Close(); err != nil {
			logrus.Errorf("Unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		message := <-message.Messages()
		c <- message
	}

}

func (c *KafkaConsumer) Consume(topics []string, signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)

	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Unable to get partition got error %v", err)
			continue
		}

		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}
	logrus.Infof("Kafka is Consuming ....")

ConsumerLoop:
	for {
		select {
		case message := <-chanMessage:
			logrus.Infof("New Message from Kafka, message : %v", string(message.Value))

		case signal := <-signals:
			if signal == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}
