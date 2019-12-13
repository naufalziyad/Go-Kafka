package main

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/naufalziyad/Go-Kafka/consumer"
	"github.com/sirupsen/logrus"
)

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}

func main() {
	//logger setup
	logFormat := new(logrus.TextFormatter)
	logFormat.TimestampFormat = "2006-01002 15:04:05"
	logFormat.FullTimestamp = true
	logrus.SetFormatter(logFormat)
	//end logger setup

	kafkaConfig := getKafkaConfig("", "")
	consumers, err := sarama.NewConsumer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka consumer got error: %v", err)
		return
	}

	defer func() {
		if err := consumers.Close(); err != nil {
			logrus.Fatal(err)
			return
		}
	}()

	kafkaConsumer := &consumer.KafkaConsumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)
	kafkaConsumer.Consume([]string{"Test_topic"}, signals)

}
