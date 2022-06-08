package uic_go_sdl2

import (
	"fmt"
	"os"
	"strings"
)

func defaultsDeclaration(fontPath string) string {

	return fmt.Sprintf("\nvar (\n\tDefaultFont = \"%s\"\n\tDefaultButtonBorderColor uint32 = 0x00000000\n\tDefaultButtonColor  uint32        = 0xD1D1D1D1\n\tDefaultButtonTextSize       = 12\n\tDefaultLabelSize            = 14\n\tDefaultTextSize             = 14\n\tDefaultTextFieldColor  uint32     = 0xffffffff\n\tDefaultTextFieldBorderColor uint32 = 0x00000000\n\tDefaultViewColor      uint32      = 0xeeeeeeee\n\tDefaultViewBorderColor  uint32    = 0x08080808\n)\n\n", fontPath)
}

func constructMain(fontDir string, viewIdentifiers []string) string {
	builder := strings.Builder{}
	builder.WriteString("package gui\n")
	builder.WriteString("\nimport \"github.com/DanielHauge/ge-go-sdl2\"\n\n")
	fontDirSafe := strings.Replace(fontDir, "\\", "\\\\", -1)
	builder.WriteString(defaultsDeclaration(fmt.Sprintf("%s%s%sdefault.ttf", fontDirSafe, string(os.PathSeparator), string(os.PathSeparator))))
	builder.WriteString("func RunGui(){\n")
	builder.WriteString("var views []ge_go_sdl2.View\n")
	for _, viewIdentifier := range viewIdentifiers {
		builder.WriteString(fmt.Sprintf("construct_%s()\n", viewIdentifier))
		builder.WriteString(fmt.Sprintf("views = append(views, %s)\n", viewIdentifier))
	}
	builder.WriteString(fmt.Sprintf("ge_go_sdl2.Run(views)\n}"))

	return builder.String()
}
