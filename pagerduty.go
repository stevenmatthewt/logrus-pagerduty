package pagerduty

import (
	"fmt"

	pd "github.com/PagerDuty/go-pagerduty"
	"github.com/sirupsen/logrus"
)

type hook struct {
	serviceKey string
	levels     []logrus.Level
}

func NewHook(serviceKey string) *hook {
	return &hook{
		serviceKey: serviceKey,
		levels: []logrus.Level{
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	}
}

func (hook *hook) Fire(entry *logrus.Entry) error {
	event := pd.Event{
		ServiceKey:  hook.serviceKey,
		Type:        "trigger",
		Description: entry.Message,
		Details:     entry.Data,
	}
	_, err := pd.CreateEvent(event)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Print("TESTING -- Fire")
	return nil
}

func (hook *hook) sendToPagerduty() {

}

func (hook *hook) Levels() []logrus.Level {
	return hook.levels
}
