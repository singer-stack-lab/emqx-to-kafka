package server

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/singer-stack-lab/emqx-to-kafka/config"
)

type compiledRule struct {
	pattern    *regexp.Regexp
	kafkaTopic string
}

type TopicRouter struct {
	rules []compiledRule
}

func NewTopicRouter(rules []config.MappingRule) *TopicRouter {
	compiledRules := make([]compiledRule, 0, len(rules))
	for i, rule := range rules {
		pattern := strings.TrimSpace(rule.EmqxTopicPrefix)
		if pattern == "" {
			panic(fmt.Errorf("empty EmqxTopicPrefix for rule %d", i))
		}
		if !strings.HasPrefix(pattern, "^") {
			pattern = "^" + pattern
		}

		re, err := regexp.Compile(pattern)
		if err != nil {
			panic(fmt.Errorf("invalid EmqxTopicPrefix regex for rule %d (%q): %w", i, rule.EmqxTopicPrefix, err))
		}

		compiledRules = append(compiledRules, compiledRule{
			pattern:    re,
			kafkaTopic: rule.KafkaTopic,
		})
	}

	return &TopicRouter{
		rules: compiledRules,
	}
}

func (r *TopicRouter) Map(emqxTopic string) (bool, string) {
	for _, rule := range r.rules {
		if rule.pattern.MatchString(emqxTopic) {
			return true, rule.kafkaTopic
		}
	}
	return false, ""
}
