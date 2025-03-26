package messagingmq

type MessageConsumer interface {
	StartConsuming() error
	ProcessMessage(message string) error
}