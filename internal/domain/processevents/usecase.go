package processevents

type ProcessEventsUC struct {
	EventRepository IGetEvent
}

type Input struct {
	Uid   string
	Value int
}

type Output struct {
	Uid   string
	Value int
}

func (t *ProcessEventsUC) Perfom(Input) Output {
	evento := t.EventRepository.GetEvent(Input{Uid: "Maycon", Value: 100})
	return evento
}

func (t *ProcessEventsUC) Setup(e interface{ IGetEvent }) {
	t.EventRepository = e
}
