package ioc

import (
	"crypto/tls"

	"github.com/singer-stack-lab/emqx-to-kafka/config"

	"github.com/IBM/sarama"
)

func InitSarmaClient(c *config.Config) sarama.Client {
	conf := c.Kafka
	cf := sarama.NewConfig()
	cf.Version = sarama.V3_3_1_0

	if conf.NeedAuth {
		cf.Net.SASL.Enable = true
		cf.Net.SASL.User = conf.Username
		cf.Net.SASL.Password = conf.Password
		cf.Net.SASL.Mechanism = sarama.SASLTypePlaintext
		cf.Net.TLS.Enable = true
		cf.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: true, // 阿里云不要求证书校验
		}
	}
	cf.Producer.Partitioner = sarama.NewHashPartitioner
	cf.Producer.Compression = sarama.CompressionSnappy
	cf.Producer.CompressionLevel = sarama.CompressionLevelDefault
	cf.Producer.RequiredAcks = sarama.NoResponse // Wait for all in-sync replicas to ack the message
	cf.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	cf.Producer.Return.Successes = true          // 必须
	cf.Producer.Return.Errors = true             // 建议
	cf.Producer.RequiredAcks = sarama.WaitForAll
	client, err := sarama.NewClient(conf.Brokers, cf)
	if err != nil {
		panic(err)
	}
	return client
}
