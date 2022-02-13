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
}
