package consumer

import (
	"clean-gin-template/pkg/logger"
	"context"
	"errors"
	"fmt"
	"strings"
)

type (
	jsonObj = map[string]interface{}
	jsonArr = []interface{}
)

// consumer interface
type Consumer interface {
	//초기 작업
	Initializer
	//읽어 오기
	Consume(ctx context.Context) error
}

// 컨슈머 팩토리 함수
type ConsumerFactory func(config jsonObj) Consumer

// 컨슈머 팩토리 저장소
var consumerFactories = make(map[string]ConsumerFactory)

// 컨슈머 등록 함수
func Register(name string, factory ConsumerFactory) {
	logger.Debugf("Registering consumer factory for %s", name)
	if factory == nil {
		logger.Panicf("Consumer factory %s does not exist.", name)
	}
	_, registered := consumerFactories[name]
	if registered {
		logger.Errorf("Consumer factory %s already registered. Ignoring.", name)
	}
	consumerFactories[name] = factory
}

// 컨슈머 생성 함수
func CreateConsumer(name string, config jsonObj) (Consumer, error) {

	factory, ok := consumerFactories[name]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableConsumers := make([]string, 0)
		for k := range consumerFactories {
			availableConsumers = append(availableConsumers, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid Consumer name. Must be one of: %s", strings.Join(availableConsumers, ", ")))
	}

	// Run the factory with the configuration.
	return factory(config), nil
}
