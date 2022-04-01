package hook

import (
	"fmt"
	"os"

	"github.com/LyricTian/queue"
	"github.com/sirupsen/logrus"
)

type ExecClose interface {
	Exec(*logrus.Entry) error
	Close() error
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

func (h *Hook) Levels() []logrus.Level {
	return h.o.levels
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	h.q.Push(queue.NewJob(entry.Dup(), func(i interface{}) {
		v := i.(*logrus.Entry)
		fmt.Println("cf fire")
		h.Exec(v)
	}))
	return nil
}

func (h *Hook) Exec(entry *logrus.Entry) {
	fmt.Println("cf Exec hook")
	for k, v := range h.o.extra {
		if _, ok := entry.Data[k]; !ok {
			entry.Data[k] = v
		}
	}
	if filter := h.o.filter; filter != nil {
		entry = filter
	}
	err := h.e.Exec(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[logrus-hook] execution error:%s", err.Error())
	}
}
func (h *Hook) Flush() {
	h.q.Terminate()
	h.e.Close()
}

var defaultOptions = options{
	maxQueues: 512,
	maxWorks:  1,
	levels: []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	},
}

func New(exceclose ExecClose, opt ...func(opt *options)) *Hook {
	dopts := defaultOptions
	for _, o := range opt {
		o(&dopts)
	}
	q := queue.NewQueue(dopts.maxQueues, dopts.maxWorks)
	q.Run()
	fmt.Println("cf New hook")
	return &Hook{
		o: dopts,
		q: q,
		e: exceclose,
	}
}