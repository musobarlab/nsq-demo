package pub

import (
	"github.com/nsqio/go-nsq"
)

//PublisherImpl struct
type PublisherImpl struct {
	producer *nsq.Producer
}

//NewPublisher
func NewPublisher(address string, config *nsq.Config) (*PublisherImpl, error) {
	producer, err := nsq.NewProducer(address, config)
	if err != nil {
		return nil, err
	}
	return &PublisherImpl{producer: producer}, nil
}

//Publish
func (pub *PublisherImpl) Publish(topic string, message []byte) error {

	//defer pub.producer.Stop()

	err := pub.producer.Publish(topic, message)

	if err != nil {
		return err
	}

	return nil
}
