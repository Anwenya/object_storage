package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

const host = "amqp://wlq:wlq@localhost:5672"

func TestPublish(t *testing.T) {
	q := New(host)
	defer q.Close()
	q.Bind("test")
	fmt.Println(q.Name)
	q2 := New(host)
	defer q2.Close()
	q2.Bind("test")
	fmt.Println(q2.Name)
	q3 := New(host)
	defer q3.Close()
	fmt.Println(q3.Name)
	expect := "test"
	q3.Publish("test2", "any")
	q3.Publish("test", expect)

	fmt.Println(2)
	c := q.Consume()
	msg := <-c
	var actual interface{}
	err := json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("expected %s, actual %s", expect, actual)
	}
	if msg.ReplyTo != q3.Name {
		t.Error(msg)
	}
	fmt.Println(3)
	c2 := q2.Consume()
	msg = <-c2
	err = json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("expected %s, actual %s", expect, actual)
	}
	if msg.ReplyTo != q3.Name {
		t.Error(msg)
	}
	fmt.Println(4)
	fmt.Println(msg.ReplyTo)
	q2.Send(msg.ReplyTo, "test3")
	fmt.Println(5)
	c3 := q3.Consume()
	fmt.Println(6)
	msg = <-c3
	fmt.Println(7)
	if string(msg.Body) != `"test3"` {
		t.Error(string(msg.Body))
	}
}

func TestSend(t *testing.T) {
	q := New(host)
	defer q.Close()

	q2 := New(host)
	defer q2.Close()

	expect := "test"
	expect2 := "test2"
	q.Send(q.Name, expect)
	q2.Send(q2.Name, expect2)

	c := q.Consume()
	msg := <-c
	var actual interface{}
	err := json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}

	if actual != expect {
		t.Errorf("expected %s, actual %s", expect, actual)
	}

	c2 := q2.Consume()
	msg = <-c2
	err = json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect2 {
		t.Errorf("expected %s, actual %s", expect2, actual)
	}
}

func TestT(t *testing.T) {
	q := New(host)
	defer q.Close()

	q.Send(q.Name, "123123")

	msgs, err := q.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
