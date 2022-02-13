package hook

import "github.com/sirupsen/logrus"

type ExecClose interface {
	exec(*logrus.Entry) error
	close()
}

type options struct {
	maxQueues int
	maxwork   int
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
		o.maxwork = maxWorks
	}
}
