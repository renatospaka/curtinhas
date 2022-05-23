package event_test

import (
	"testing"
	assert "github.com/stretchr/testify/assert"

	"github.com/renatospaka/event-dispatcher/event"
)

type testListener struct {
	data interface{}
	called bool
}

func (l *testListener) Handle() error {
	l.called = true
	return nil
}

func (l *testListener) SetData(data interface{}) {
	l.data = data
}

type testEvent struct {
	data interface{}
}

func (t *testEvent) GetKey() string {
	return "test"
}

func (t * testEvent) GetData() interface{} {
	t.data = "test"
	return t.data
}

func TestEventDispatcher_AddListener(t *testing.T) {
	ed := event.NewEventDispatcher()
	testListener := &testListener{}
	ed.AddListener("test", testListener)

	assert.Equal(t, 1, len(ed.Listeners["test"]))
	assert.Equal(t, testListener, ed.Listeners["test"][0])
}

func TestEventDispatcher_Dispatch(t *testing.T) {
	ed := event.NewEventDispatcher()
	testListener := &testListener{}
	ed.AddListener("test", testListener)
	
	event := &testEvent{}
	ed.Dispatch(event)

	assert.True(t, testListener.called)
}