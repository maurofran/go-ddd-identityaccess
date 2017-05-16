package model

import "time"

// DomainEvent is the interface exposed by domain events.
type DomainEvent interface {
	// The timestamp of event.
	Timestamp() time.Time
	// The version of the event.
	Version() int
}

// DomainEvents represents an ordered sequence of events.
type DomainEvents []DomainEvent

// noEvents will return an empty domain events
func noEvents() DomainEvents {
	return make(DomainEvents, 0)
}

// anEvent will return a new domain events with provided event
func anEvent(event DomainEvent) DomainEvents {
	return DomainEvents{event}
}

// and will append the provided domain event
func (d DomainEvents) and(event... DomainEvent) DomainEvents {
	return append(d, event...)
}

// DomainEventPublisher is the interface used to publish events.
type DomainEventPublisher interface {
	Publish(events DomainEvents)
}

// Base class for domain events
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
