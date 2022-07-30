package rabbitmq

import (
	"encoding/json"
	"testing"
)

const host = "amqp://root:root@localhost:5672"

func TestPublish(t *testing.T) {
	q := New(host)
	defer q.Close()
	q.Bind("apiServers")

	q2 := New(host)
	defer q2.Close()
	q2.Bind("apiServers")

	q3 := New(host)
	defer q3.Close()

	var actual interface{}
	expect := "test"
	q3.Publish("apiServers", expect)

	c := q.Consume()
	msg := <-c
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
}

func TestSend(t *testing.T) {
	q := New(host)
	defer q.Close()

	q2 := New(host)
	defer q2.Close()

	expect := "test"
	expect2 := "test2"
	var actual interface{}
	q2.Send(q.Name, expect)
	q2.Send(q2.Name, expect2)

	c := q.Consume()
	msg := <-c
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
