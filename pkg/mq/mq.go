package mq

import "context"

type Reader interface {
	Subscribe(ctx context.Context, process func(key string, data []byte) (bool, error)) error
	Close() error
}

type Writer interface {
	Publish(ctx context.Context, key, message []byte) error
	Close() error
}
