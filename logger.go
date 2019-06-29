package logger

import (
	"fmt"
	"io"
	"time"
)

type OtusEvent interface {
	GetLogMessage() string
}

type HwAccepted struct {
	Id    int
	Grade int
}

func (e *HwAccepted) GetLogMessage() string {
	return fmt.Sprintf("accepted %d %d", e.Id, e.Grade)
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func (e *HwSubmitted) GetLogMessage() string {
	return fmt.Sprintf("submitted %d \"%s\"", e.Id, e.Comment)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	data := fmt.Sprintf("%s %s\n", time.Now().Format("2006-01-02"), e.GetLogMessage())
	w.Write([]byte(data))
}
