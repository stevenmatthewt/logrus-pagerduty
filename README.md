# logrus-pagerduty

## Setup

In your `main.go` file (or whenever you want to enable PagerDuty)
```go
package main

import "github.com/stevenmatthewt/logrus-pagerduty"
func main() {
    //...
    // This will enable the pagerduty integration if the PAGERDUTY_KEY
    // is found in the environment.
    if key := os.Getenv("PAGERDUTY_KEY"); key != "" {
        log.AddHook(pagerduty.NewHook(key))
    }
    //...
}
```


## How it works

Anytime a `Logrus.Error()`, `Logrus.Fatal()`, or `Logrus.Panic()` log occurs, the contents of the log will automatically be sent to PagerDuty. The level of alerting can be configured within PagerDuty itself.

The actual message that is logged with `logrus` will be sent as the PagerDuty summary. Any additionaly fields that are sent using `logrus.WithField()` will be sent as additional data to PagerDuty.

Example:
```go
// The following log will send an event to PagerDuty with a
// summary of "everything is broken!". Additionally,
// it will send an extra field of "area":"The Whole World" as well.
// These additionaly fields can be used to send arbitrary information
// thatn may be useful for debugging.
logrus.
    WithField("area", "The Whole World").
    Error("everything is broken!")
```
