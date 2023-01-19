package log

import "github.com/sirupsen/logrus"

type LogCtxKeyType string

const LogCtxKey LogCtxKeyType = "log"

type Entry struct {
	l *logrus.Entry
}

type IFields interface {
	map[string]interface{}
}

type Fields map[string]interface{}
