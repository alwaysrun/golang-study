package learn

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

// publisher
type Publisher struct {
	m           sync.RWMutex
	buffLen     int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(publishTimeout time.Duration, bufLen int) *Publisher {
	return &Publisher{
		buffLen:     bufLen,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffLen)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) Unsubscribe(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) SendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		p.SendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func PubSub()  {
	p:=NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{})bool{
		if s,ok := v.(string); ok{
			return strings.Contains(s, "golang")
		}
		return false
	})

	go func() {
		for i:=0; i<100; i++{
			p.Publish(fmt.Sprint(i, "-hello, world"))
			time.Sleep(time.Millisecond*10)
		}
	}()

	go func() {
		for i:=0; i<100; i++{
			p.Publish(fmt.Sprint(i, "-golang"))
			time.Sleep(time.Millisecond*10)
		}
	}()

	go func() {
		for msg :=range all{
			fmt.Println("all:", msg)
		}
	}()
	go func() {
		for msg:=range golang{
			fmt.Println("goland:", msg)
		}
	}()

	time.Sleep(3*time.Second)
}