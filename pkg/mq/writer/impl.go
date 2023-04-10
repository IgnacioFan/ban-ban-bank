package writer

import (
	"context"
	"fmt"
	"go-bank-express/pkg/mq"

	"github.com/segmentio/kafka-go"
)

type Impl struct {
	Writer *kafka.Writer
}

func NewWriter(brokers []string, topic string) mq.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		RequiredAcks: 1,
	})
	return &Impl{
		Writer: writer,
	}
}

func (i *Impl) Publish(ctx context.Context, key, val []byte) error {
	message := kafka.Message{
		Key:   key,
		Value: val,
	}
	if err := i.Writer.WriteMessages(ctx, message); err != nil {
		fmt.Println("Error producing message:", err)
		return err
	}
	return nil
}

func (i *Impl) Close() error {
	if err := i.Writer.Close(); err != nil {
		return err
	}
	return nil
}
