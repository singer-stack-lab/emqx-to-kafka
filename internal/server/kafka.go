package server

import (
	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(client sarama.Client) (*KafkaProducer, error) {
	p, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	return &KafkaProducer{p}, nil
}

func (p *KafkaProducer) Send(topic string, key string, msg []byte) error {
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(msg),
	})
	return err
}
