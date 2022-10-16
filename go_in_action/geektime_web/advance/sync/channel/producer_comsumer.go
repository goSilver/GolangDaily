package channel

type Broker struct {
	consumers []*Consumer
}

func (b *Broker) Produce(msg string) {
	for _, c := range b.consumers {
		c.ch <- msg
	}
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumers = append(b.consumers, c)
}

type Consumer struct {
	ch chan string
}
