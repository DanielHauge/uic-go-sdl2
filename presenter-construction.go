package uic_go_sdl2

import (
	"fmt"
	"strings"
)

type stateChangeEventInterface struct {
	functionIdentifier  string
	ElementIdIdentifier string
	Property            string
}

type userEventInterface struct {
	functionIdentifier string
	hasArg             bool
}

func constructPresenter(ueInterfaces []userEventInterface, sceInterfaces []stateChangeEventInterface) string {
	builder := strings.Builder{}
	builder.WriteString("package gui\n")

	if len(sceInterfaces) > 0 {
		builder.WriteString("\nimport \"github.com/DanielHauge/ge-go-sdl2\"\n\n")
	}

	for _, ueInterface := range ueInterfaces {
		builder.WriteString(constructUserInputCallback(ueInterface))
	}
	for _, sceInterface := range sceInterfaces {
		builder.WriteString(constructStateChangeNotifyFunction(sceInterface))
	}
	return builder.String()
}

func constructUserInputCallback(eventInterface userEventInterface) string {
	if eventInterface.hasArg {
		return fmt.Sprintf("func %s(value string) {\n// Update model\n}\n\n", eventInterface.functionIdentifier)
	} else {
		return fmt.Sprintf("func %s() {\n// Update model\n}\n\n", eventInterface.functionIdentifier)
	}
}

func constructStateChangeNotifyFunction(eventInterface stateChangeEventInterface) string {
	return fmt.Sprintf("func %s(value string){\n ge_go_sdl2.NotifyPropertyChangeAsync(%s, \"%s\", value) }\n", eventInterface.functionIdentifier, eventInterface.ElementIdIdentifier, eventInterface.Property)
}
