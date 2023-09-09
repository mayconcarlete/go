package main

import (
	"fmt"

	"app.com/domain/internal/domain/processevents"
	repository "app.com/domain/internal/infra/EventRepository"
)

func main() {
	repository := repository.EventRepository{}

	eventUseCase := processevents.ProcessEventsUC{}
	eventUseCase.Setup(&repository)

	result := eventUseCase.Perfom(processevents.Input{Uid: "asdas", Value: 122})
	fmt.Println(result)

}
