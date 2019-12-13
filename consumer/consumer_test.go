package consumer_test

import (
	"os"
	"testing"
	"time"

	"github.com/Shopify/sarama"

	"github.com/Shopify/sarama/mocks"
)

func TestConsume(t *testing.T) {
	consumer := mocks.NewConsumer(t, nil)
	defer func() {
		if err := consumer.Close(); err != nil {
			t.Error(err)
		}
	}()

	consumer.SetTopicMetadata(map[string][]int32{
		"Test_topic": {0},
	})

	kafka := &consumer.KafkaConsumer{
		Consumer: consumers,
	}

	consumer.ExpectConsumePartition("Test_topic", 0, sarama.OffsetNewest).YieldMessage(&sarama.ConsumerMessage{
		Value: []byte("haii hai")})
	signals := make(chan os.Signal, 1)

	go kafka.Consume([]string{"Test_topic"}, signals)
	timeout := time.After(2 * time.Second)
	for {
		select {
		case <-timeout:
			signals <- os.Interrupt
			return
		}
	}
}
