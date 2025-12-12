package server

import (
	"github.com/singer-stack-lab/emqx-to-kafka/config"
)

type TopicRouter struct {
	rules []config.MappingRule
}

func NewTopicRouter(rules []config.MappingRule) *TopicRouter {
	return &TopicRouter{
		rules: rules,
	}
}

func (r *TopicRouter) Map(emqxTopic string) (bool, string) {
	for _, rule := range r.rules {
		if len(emqxTopic) >= len(rule.EmqxTopicPrefix) && emqxTopic[:len(rule.EmqxTopicPrefix)] == rule.EmqxTopicPrefix {
			return true, rule.KafkaTopic
		}
	}
	return false, ""
}
