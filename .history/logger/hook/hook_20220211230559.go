package hook

import "github.com/sirupsen/logrus"

type ExecClose interface {
	exec(*logrus.Entry) error
	close()
}
