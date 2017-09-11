package pagerduty

import (
	"testing"

	"github.com/sirupsen/logrus"
	"gopkg.in/h2non/gock.v1"
)

func TestHook(t *testing.T) {
	defer setupGock()()

	logrus.AddHook(NewHook("test_key"))
	logrus.Error("test_message")

	if !gock.IsDone() {
		t.Error("not all gock requests were met")
	}
}

func setupGock() func() {
	gock.New("http").Post("").BodyString(".*test_key.*").Reply(200).JSON(nil)
	return gock.Off
}
