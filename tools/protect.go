package tools

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

func ProtectRun(Entry func()) {
	defer func() {
		err := recover()
		if err != nil {
			switch err.(type) {
			case runtime.Error:
				logrus.Errorln(err)
			default:
				logrus.Errorln(err)
			}
		}
	}()
	Entry()
}
