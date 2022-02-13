package hook

import (
	"github.com/LyricTian/queue"
	"github.com/sirupsen/logrus"
)

type ExecClose interface {
	exec(*logrus.Entry) error
	close()
}

type options struct {
	maxQueues int
	maxWorks  int
	extra     map[string]interface{}
	filter    *logrus.Entry
	levels    []logrus.Level
}

func SetMaxQueues(maxQueues int) func(*options) {
	return func(o *options) {
		o.maxQueues = maxQueues
	}
}
func SetMaxWorks(maxWorks int) func(*options) {
	return func(o *options) {
		o.maxWorks = maxWorks
	}
}
func SetExtra(extra map[string]interface{}) func(*options) {
	return func(o *options) {
		o.extra = extra
	}
}
func SetFilter(filter *logrus.Entry) func(*options) {
	return func(o *options) {
		o.filter = filter
	}
}
func SetLevels(levels ...logrus.Level) func(*options) {
	return func(o *options) {
		o.levels = levels
	}
}

type Hook struct {
	o options
	e ExecClose
	q *queue.Queue
}

func (h *Hook) Leves() []logrus.Level {
	return h.o.levels
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	h.q.Push(queue.NewJob(entry.Dup(), func(i interface{}) {
		v := i.(*logrus.Entry)
		h.Exec(v)
	}))
	return nil
}

func (h *Hook) Exec(entry *logrus.Entry) {

}
func (h *Hook) Flush() {
	h.q.Terminate()
	h.e.close()
}

func New() {

}
