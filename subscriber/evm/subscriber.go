package evm

import (
	"context"
	"errors"

	"github.com/mitchellh/mapstructure"

	"github.com/imxyb/contract-indexer/subscriber"
)

// Subscriber subscribes to the chain
type Subscriber struct {
	conf *Config
}

// LoadConfig loads the config
func (s *Subscriber) LoadConfig(conf interface{}) error {
	var ok bool
	cm, ok := conf.(map[string]any)
	if !ok {
		return errors.New("invalid config")
	}
	if err := mapstructure.Decode(cm, &s.conf); err != nil {
		return err
	}
	return nil
}

// Subscribe subscribes to the chain
func (s *Subscriber) Subscribe(ctx context.Context) (*subscriber.Output, error) {
	return &subscriber.Output{
		ContractName: "test",
		Method:       "test",
		Params: map[string]any{
			"test": "test",
		},
	}, nil
}

func init() {
	subscriber.Register("evm", &Subscriber{})
}
