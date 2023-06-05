package slack

import "github.com/sirupsen/logrus"

type InMemoryClient struct{}

func NewInMemoryClient() Client {

	return &InMemoryClient{}
}

func (c *InMemoryClient) Info(message string) error {

	logrus.Infof("sent info message to slack '%s'", message)
	return nil
}
