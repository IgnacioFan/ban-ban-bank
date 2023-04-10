package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-bank-express/pkg/mq/writer"
)

type BaseHandler struct {
}

func (h *BaseHandler) PubMessage(ctx context.Context, model interface{}) error {
	writer := writer.NewWriter([]string{"localhost:9092"}, "create-user")
	defer writer.Close()
	payload, err := json.Marshal(model)
	if err != nil {
		return err
	}
	fmt.Println("Prep", ctx, payload)
	err = writer.Publish(ctx, []byte("abc"), payload)
	if err != nil {
		fmt.Println("Error producing message:", err)
		return err
	}
	return nil
}

// func (p *Producer) Listen(key string) []byte {
// 	cb := make(chan []byte)
// 	p.MCallback.Set(key, cb)
// 	data := <-cb
// 	p.MCallback.Delete(key)
// 	return data
// }
