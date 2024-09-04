package main

import (
	"github.com/sirupsen/logrus"
)

type Hook struct {
	Name string
}

func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["name"] = h.Name
	return nil
}

func main() {
	h := &Hook{Name: "hook"}
	logrus.AddHook(h)
	logrus.Info("info msg")
}
