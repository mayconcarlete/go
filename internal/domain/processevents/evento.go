package processevents

// type Input struct {
// 	Uid   string
// 	Value int
// }

// type Output struct {
// 	Uid   string
// 	Value int
// }

// type IGetEvent interface {
// 	GetEvent(Input Input) Output
// }

type EventRepository struct {
	// Config string
	// GetEvent processevents.IGetEvent
}

// func (r *EventRepository) GetEvent() {
// 	value := processevents.Input{Uid: "Maycon", Value: 100}
// 	processevents.IGetEvent
// 	return
// }

//	func (t *EventRepository) GetEvent(Input processevents.Input) processevents.Output {
//		return processevents.Output{
//			Uid:   "Maycao Ufes",
//			Value: 1000,
//		}
//	}
func (t *EventRepository) GetEvent(Input Input) Output {
	return Output{
		Uid:   "Maycao Ufes",
		Value: 1000,
	}
}
