package tools

import (
	"github.com/sirupsen/logrus"
)

func Check(err error) {
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}
