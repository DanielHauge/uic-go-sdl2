package uic_go_sdl2

import (
	"fmt"
	"strings"
)

func defaultsDeclaration(fontPath string) string {
	fontPath = strings.Replace(fontPath, "\\", "/", -1)
	fontPath = strings.Replace(fontPath, ":", "", -1)
	return fmt.Sprintf("\nvar (\n\tDefaultFont = \"%s\"\n\tDefaultButtonBorderColor uint32 = 0xffff0000\n\tDefaultButtonColor  uint32        = 0xffffff00\n\tDefaultButtonTextSize       = 12\n\tDefaultLabelSize            = 14\n\tDefaultTextSize             = 14\n\tDefaultTextFieldColor  uint32     = 0xffffffff\n\tDefaultTextFieldBorderColor uint32 = 0x00000000\n\tDefaultViewColor      uint32      = 0xeeeeeeee\n\tDefaultViewBorderColor  uint32    = 0xffffff00\n)\n\n", fontPath)
}

func constructMain(fontDir string, viewIdentifiers []string) string {
	builder := strings.Builder{}
	builder.WriteString("package gui\n")
	builder.WriteString("\nimport ge_go_sdl2 \"github.com/DanielHauge/ge-go-sdl2\"\n\n")
	builder.WriteString(defaultsDeclaration(fmt.Sprintf("%s/default.ttf", fontDir)))
	builder.WriteString("func RunGui(){\n")
	builder.WriteString("var views []ge_go_sdl2.View\n")
	for _, viewIdentifier := range viewIdentifiers {
		builder.WriteString(fmt.Sprintf("construct_%s()\n", viewIdentifier))
		builder.WriteString(fmt.Sprintf("views = append(views, %s)\n", viewIdentifier))
	}
	builder.WriteString(fmt.Sprintf("ge_go_sdl2.Run(views)\n}"))

	return builder.String()
}
