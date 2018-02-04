package pub

//Publisher interface
type Publisher interface {
	Publish(topic string, message []byte) error
}
