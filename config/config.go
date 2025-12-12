package config

type KafkaConfig struct {
	Brokers  []string `mapstructure:"brokers" yaml:"brokers"`
	Username string   `mapstructure:"username" yaml:"userName"`
	Password string   `mapstructure:"password" yaml:"password"`
	ClientID string   `mapstructure:"clientid" yaml:"clientId"`
	NeedAuth bool     `mapstructure:"needauth" yaml:"needAuth"`
}

type MappingRule struct {
	EmqxTopicPrefix string `mapstructure:"emqxtopicprefix" yaml:"EmqxTopicPrefix"`
	KafkaTopic      string `mapstructure:"kafkatopic" yaml:"KafkaTopic"`
}

type Config struct {
	Kafka KafkaConfig   `mapstructure:"kafka" yaml:"Kafka"`
	Rules []MappingRule `mapstructure:"rules" yaml:"Rules"`
}
