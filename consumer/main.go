package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()

	q, _ := nsq.NewConsumer("wurys", "ch", config)

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", string(message.Body))
		wg.Done()
		return nil
	}))

	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	wg.Wait()

}
