package main

import (
	"fmt"
	"pi-hole"
)

func main() {
	config := Configuration()
	p := pihole.New(config.Pihole.Url, config.Pihole.Pwhash)

	t, err := p.GetType()
	fmt.Printf("Type: %v, %v\n", t, err)

	v, err := p.GetVersion()
	fmt.Printf("Version: %v, %v\n", v, err)
	s, err := p.GetStatus()
	fmt.Printf("Status: %v, %v\n", s, err)
	su, err := p.GetSummary()
	fmt.Printf("Summary: %v, %v\n", su, err)
	sur, err := p.GetSummaryRaw()
	fmt.Printf("SummaryRaw: %v, %v\n", sur, err)
	o, err := p.GetOverTimeData10mins()
	fmt.Printf("OverTimeData10mins: %v, %v\n", o, err)
	ti, err := p.GetTopItems()
	fmt.Printf("TopItems: %v, %v\n", ti, err)
	tc, err := p.GetTopClients()
	fmt.Printf("TopClients: %v, %v\n", tc, err)
	f, err := p.GetForwardDestinations()
	fmt.Printf("ForwardDestinations: %v, %v\n", f, err)
	q, err := p.GetQueryTypes()
	fmt.Printf("QueryTypes: %v, %v\n", q, err)
	q2, err := p.GetQueries()
	fmt.Printf("Queries: %v, %v\n", q2, err)
	se, err := p.Enable()
	fmt.Printf("Enable: %v, %v\n", se, err)
	sd, err := p.Disable(300)
	fmt.Printf("Disable: %v, %v\n", sd, err)
}
