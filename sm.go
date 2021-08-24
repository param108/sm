package main

import (
	"fmt"
	sm "github.com/qmuntal/stateless"
)

//AlertState has the different Alert States a node can be in
type AlertState string

//AlertTrigger has the different Triggers
type AlertTrigger string

const (
	//Boring Alert State: No alerts currently active and all alerts have resolved
	Boring = "Boring"

	//Alerted Alert State: Has one or more active alerts
	Alerted = "Alerted"

	//NotAlerted Alert State: No Alerts in the last period
	NotAlerted = "NotAlerted"

	//AlertSeen Trigger: An alert is see in this period
	AlertSeen = "AlertSeen"

	//NoAlerts1Interval Trigger: This interval no alerts seen
	NoAlerts1Interval = "NoAlerts1Interval"

	//NoAlertWaitInterval Trigger: No Alerts were seen for the complete cool off period, this alert is resolved
	NoAlertWaitInterval = "NoAlertWaitInterval"
)

func main() {

	bm := sm.NewStateMachine(Boring)
	bm.Configure(Boring).
		Permit(AlertSeen, Alerted).
		PermitReentry(NoAlerts1Interval).
		PermitReentry(NoAlertWaitInterval)
	bm.Configure(Alerted).
		PermitReentry(AlertSeen).
		Permit(NoAlerts1Interval, NotAlerted)
	bm.Configure(NotAlerted).
		Permit(AlertSeen, Alerted).
		Permit(NoAlertWaitInterval, Boring).
		PermitReentry(NoAlerts1Interval)
	graph := bm.ToGraph()
	fmt.Println(graph)

}
