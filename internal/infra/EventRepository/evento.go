package repository

import "app.com/domain/internal/domain/processevents"

type EventRepository struct{}

func (t *EventRepository) GetEvent(Input processevents.Input) processevents.Output {
	return processevents.Output{
		Uid:   "Maycao Ufes",
		Value: 1000,
	}
}
