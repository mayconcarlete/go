package processevents

type Input struct {
	Uid   string
	Value int
}

type Output struct {
	Uid   string
	Value int
}

type ProcessEventsUC struct {
	EventRepository IGetEvent
}

func (t *ProcessEventsUC) Perfom(Input) Output {
	evento := t.EventRepository.GetEvent(Input{Uid: "Maycon", Value: 100})
	return evento
}

func (t *ProcessEventsUC) Setup(e interface{ IGetEvent }) {
	t.EventRepository = e
}

// func Perfom(t ProcessEventsUC,  Input) Output {
// 	t.EventRepository.GetEvent()
// 	&t.EventRepository.GetEvent()
// 	return Output{Uid: "1111", Value: 10}
// }

// func Setup() {
// 	input := Perfom(Input{Value: 2, Uid: "123-abc"})
// 	fmt.Println(input)
// }
