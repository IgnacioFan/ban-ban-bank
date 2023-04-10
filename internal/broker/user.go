package broker

import (
	"context"
	"fmt"
	"go-bank-express/pkg/mq"
)

type UserConsumer struct {
	Reader mq.Reader
}

func NewUserConsumer(reader mq.Reader) *UserConsumer {
	return &UserConsumer{
		Reader: reader,
	}
}

func (c *UserConsumer) CreateUser(ctx context.Context) {
	fmt.Println("Sub message", ctx)
	c.Reader.Subscribe(ctx, func(key string, data []byte) (bool, error) {
		fmt.Println("In the process", key, data)
		// bind user with user entity
		// register user
		// publish to response topic and send data
		return true, nil
	})
}
