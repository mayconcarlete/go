package main

import (
	"fmt"

	"app.com/domain/internal/domain/processevents"
)

func main() {
	repository := processevents.EventRepository{}
	eventUseCase := processevents.ProcessEventsUC{}
	eventUseCase.Setup(&repository)

	result := eventUseCase.Perfom(processevents.Input{Uid: "asdas", Value: 122})
	fmt.Println(result)

}
