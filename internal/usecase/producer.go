package usecase

// type Producer struct {
// 	MQ        mq.MQ
// 	MCallback MessageCallback
// }

// func NewProducer(mq mq.MQ) *Producer {
// 	return &Producer{
// 		MQ: mq,
// 	}
// }

// func (p *Producer) PubMessage(ctx context.Context, topic string, model interface{}) error {
// 	fmt.Println("Pub message", ctx, topic, model)
// 	payload, err := json.Marshal(model)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("create payload", payload)
// 	return p.MQ.Publish(ctx, payload)
// }

// func (p *Producer) Listen(key string) []byte {
// 	cb := make(chan []byte)
// 	p.MCallback.Set(key, cb)
// 	data := <-cb
// 	p.MCallback.Delete(key)
// 	return data
// }
