package initializers

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func NewLogger() {
	Logger = logrus.New()
}
