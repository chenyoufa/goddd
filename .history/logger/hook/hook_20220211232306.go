package hook

import "github.com/sirupsen/logrus"

type ExecClose interface {
	exec(*logrus.Entry) error
	close()
}

type options struct {
	maxquese int
	maxwork  int
	extra    map[string]interface{}
	filter   *logrus.Entry
	levels   []logrus.Level
}

func SetMaxQueues(maxQueues int) func(*options) {
	return func(o *options) {
		o.maxquese = maxQueues
	}
}
