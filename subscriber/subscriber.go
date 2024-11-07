package subscriber

import (
	"context"
)

// Map SubscriberMap is the map that stores the subscriber.
var Map = map[string]Subscriber{}

// Register registers the subscriber with the given name.
func Register(name string, subscriber Subscriber) {
	Map[name] = subscriber
}

// GetSubscriber returns the subscriber with the given name.
func GetSubscriber(name string) Subscriber {
	return Map[name]
}

// Subscriber is the interface that wraps the basic Subscribe method.
type Subscriber interface {
	// LoadConfig loads the configuration from the given interface.
	LoadConfig(conf interface{}) error
	// Subscribe subscribes to the given context.
	Subscribe(ctx context.Context) (*Output, error)
}
