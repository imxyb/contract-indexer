package tasker

import (
	"context"
	"time"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"

	"github.com/imxyb/contract-indexer/subscriber"
	_ "github.com/imxyb/contract-indexer/subscriber/evm"
)

// Start tasker
func Start(ctx context.Context, confFile string) error {
	if err := configor.Load(&Conf, confFile); err != nil {
		return err
	}

	sub := subscriber.GetSubscriber(Conf.Subscriber.Protocol)
	if err := sub.LoadConfig(Conf.Subscriber.Conf); err != nil {
		return err
	}

	for {
		output, err := sub.Subscribe(ctx)
		if err != nil {
			return err
		}
		logrus.Debugf("Received message: %s", output)
		time.Sleep(1 * time.Second)
	}
}
