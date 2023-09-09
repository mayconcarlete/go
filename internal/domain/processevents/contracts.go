package processevents

type IGetEvent interface {
	GetEvent(Input Input) Output
}
