package main

import (
	"fmt"
	"time"

	"github.com/naufalziyad/Go-Kafka/producer"

	"github.com/Shopify/sarama"
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
	producers, err := sarama.NewSyncProducer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error: %v", err)
		return
	}

	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	logrus.Infof("Success create kafka sync-producer")

	kafka := &producer.KafkaProducer{
		Producer: producers,
	}

	for i := 1; i <= 50000; i++ {
		// time.Sleep(500 * time.Millisecond)
		message := fmt.Sprintf("message received %v", i)
		err := kafka.SendMessage("Test_topic", message)
		if err != nil {
			logrus.Error("Error nih")
		}
	}
}
