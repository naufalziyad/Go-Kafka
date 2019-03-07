package producer 
import (
	"fmt"

	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func (p *KafkaProducer) SendMessage(topic, msg string) error{
	fmt.Println("coba saja")
	return nil
}


