package reader

import (
	"context"
	"fmt"
	"go-bank-express/pkg/mq"

	"github.com/segmentio/kafka-go"
)

type Impl struct {
	Reader *kafka.Reader
}

func NewReader(groupId, topic string) mq.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &Impl{
		Reader: reader,
	}
}

func (i *Impl) Subscribe(ctx context.Context, process func(key string, data []byte) (bool, error)) error {
	var err error
	for {
		m, err := i.Reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}
		fmt.Printf("Received message from %s-%d [%d]: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		p := func(key string, data []byte) (ack bool, err error) {
			return process(key, data)
		}
		ack, err := p(string(m.Key), m.Value)
		fmt.Println("ACK:", ack, err)
		if err = i.Reader.CommitMessages(ctx, m); err != nil {
			fmt.Println("Error committing message to Kafka:", err)
			break
		}
		fmt.Println("Acknowledgement message sent to Kafka")
	}
	return err
}

func (i *Impl) Close() error {
	if err := i.Reader.Close(); err != nil {
		return err
	}
	return nil
}
