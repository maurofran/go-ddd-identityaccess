package model

import "time"

var noEvents = make([]DomainEvent, 0)

func anEvent(event DomainEvent) []DomainEvent {
	return someEvents(event)
}

func someEvents(events ...DomainEvent) []DomainEvent {
	return events
}

// DomainEvent is the interface exposed by domain events.
type DomainEvent interface {
	// The timestamp of event.
	Timestamp() time.Time
	// The version of the event.
	Version() int
}

type domainEvent struct {
	ts time.Time
	v  int
}

func (ev domainEvent) Timestamp() time.Time {
	return ev.ts
}

func (ev domainEvent) Version() int {
	return ev.v
}

func newDomainEvent(version int) domainEvent {
	return domainEvent{ts: time.Now(), v: version}
}
