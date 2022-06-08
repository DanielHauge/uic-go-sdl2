package uic_go_sdl2

import (
	"fmt"
	"github.com/DanielHauge/ge-go-sdl2"
	"reflect"
	"strconv"
	"strings"
)

// identifier, code and interfaces.
func constructView(view ge_go_sdl2.View) (string, string, []userEventInterface, []stateChangeEventInterface) {
	defaultColor := ifElseAssign(view.BgColor == 0, "DefaultViewColor", fmt.Sprintf("0x%x", view.BgColor))
	DefaultBorderColor := ifElseAssign(view.BorderColor == 0, "DefaultViewBorderColor", fmt.Sprintf("0x%x", view.BorderColor))
	importSdl := false
	var childElementIdentifiers []string
	var userEventsInterfaces []userEventInterface
	var stateChangeEventInterfaces []stateChangeEventInterface
	childElementConstructions := strings.Builder{}
	elementDeclarations := strings.Builder{}
	for _, cons := range view.Children {
		var identifier, construction, fId string
		fId = ""
		switch t := cons.(type) {
		case ge_go_sdl2.Container:
			identifier, construction = constructContainer(t, view.Id)
			elementDeclarations.WriteString(fmt.Sprintf("%s %s\n", identifier, "ge_go_sdl2.Container"))
			stateChangeEventInterfaces = append(stateChangeEventInterfaces, stateChangeEventInterface{ElementIdIdentifier: fmt.Sprintf("%s.Id", identifier), Property: "ViewId", functionIdentifier: fmt.Sprintf("NotifyViewChange_%s", identifier)})
			break
		case ge_go_sdl2.Text:
			importSdl = true
			identifier, construction = constructText(t, view.Id)
			elementDeclarations.WriteString(fmt.Sprintf("%s %s\n", identifier, "ge_go_sdl2.Text"))
			stateChangeEventInterfaces = append(stateChangeEventInterfaces, stateChangeEventInterface{ElementIdIdentifier: fmt.Sprintf("%s.Id", identifier), Property: "Content", functionIdentifier: fmt.Sprintf("NotifyContentChange_%s", identifier)})
			break
		case ge_go_sdl2.Button:
			identifier, construction, fId = constructButton(t, view.Id, &elementDeclarations)
			elementDeclarations.WriteString(fmt.Sprintf("%s %s\n", identifier, "ge_go_sdl2.Button"))
			userEventsInterfaces = append(userEventsInterfaces, userEventInterface{functionIdentifier: fId, hasArg: false})
			break
		case ge_go_sdl2.TextField:
			identifier, construction, fId = constructTextField(t, view.Id)
			elementDeclarations.WriteString(fmt.Sprintf("%s %s\n", identifier, "ge_go_sdl2.TextField"))
			userEventsInterfaces = append(userEventsInterfaces, userEventInterface{functionIdentifier: fId, hasArg: true})
			stateChangeEventInterfaces = append(stateChangeEventInterfaces, stateChangeEventInterface{ElementIdIdentifier: fmt.Sprintf("%s.Id", identifier), Property: "Value", functionIdentifier: fmt.Sprintf("NotifyValueChange_%s", identifier)})
			break
		default:
			fmt.Errorf("the element type '%s' was not valid for construction within view", reflect.TypeOf(t).Name())
			continue
		}

		childElementIdentifiers = append(childElementIdentifiers, identifier)
		childElementConstructions.WriteString(construction)
	}
	childrenIdentifier := fmt.Sprintf("view_%s_children", view.Id)
	childrenSliceConstruction := constructChildrenSlice(childrenIdentifier, childElementIdentifiers)
	viewIdentifier := fmt.Sprintf("View_%s", view.Id)
	viewConstruction := fmt.Sprintf(
		"%s = ge_go_sdl2.View{\nX:%d,\nY: %d,\nH:%d,\nW:%d,\nId:\"%s\",\nBorderColor:%s,\nBgColor:%s,\n\tChildren: %s,}\n",
		viewIdentifier, view.X, view.Y, view.H, view.W, view.Id, DefaultBorderColor, defaultColor, childrenIdentifier)
	elementDeclarations.WriteString(fmt.Sprintf("%s %s\n", viewIdentifier, "ge_go_sdl2.View"))
	importSdlString := ifElseAssign(importSdl, "\"github.com/veandco/go-sdl2/sdl\"", "")
	viewCodeFileContent := fmt.Sprintf("package gui\n\nimport (\n \"github.com/DanielHauge/ge-go-sdl2\"\n %s \n)\n // UI element declaration \n var (\n%s) \n\n// %s view and elements construction \nfunc construct_%s(){\n %s \n// %s view element construction \n %s %s \n} \n", importSdlString, elementDeclarations.String(), view.Id, viewIdentifier, childElementConstructions.String(), view.Id, childrenSliceConstruction, viewConstruction)

	return viewIdentifier, viewCodeFileContent, userEventsInterfaces, stateChangeEventInterfaces
}

func constructChildrenSlice(identifier string, identifiers []string) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("var %s []interface{}\n", identifier))
	for _, id := range identifiers {
		builder.WriteString(fmt.Sprintf("%s = append(%s, %s)\n", identifier, identifier, id))
	}
	return builder.String()
}

func constructContainer(container ge_go_sdl2.Container, viewId string) (string, string) {
	identifier := fmt.Sprintf("%s_container_%s", viewId, container.Id)
	construction := fmt.Sprintf("\n// %s container construction \n%s = ge_go_sdl2.Container{\n X: %d, \nY:%d, \n Id: \"%s\", \n ViewId: \"%s\",\n}\n", container.Id, identifier, container.X, container.Y, container.Id, container.ViewId)
	return identifier, construction
}

func AlignmentToString(alignment ge_go_sdl2.HAlign) string {
	switch alignment {
	case ge_go_sdl2.Left:
		return "ge_go_sdl2.Left"
	case ge_go_sdl2.Center:
		return "ge_go_sdl2.Center"
	case ge_go_sdl2.Right:
		return "ge_go_sdl2.Right"
	default:
		panic("No valid alignment found")
	}
}

func constructText(txt ge_go_sdl2.Text, viewId string) (string, string) {
	identifier := fmt.Sprintf("%s_text_%s", viewId, txt.Id)

	defaultSize := ifElseAssign(txt.Size == 0, "DefaultLabelSize", strconv.Itoa(txt.Size))
	construction := fmt.Sprintf("\n// %s text construction \n%s = ge_go_sdl2.Text{\nId:\"%s\",\nContent:\"%s\",\nFont: DefaultFont,\n\t\tSize:  %s,\n\t\tX: %d,\n\t\tY: %d,W:%d,\nH: %d,\n\nAlignment: %s,\nTextColor: sdl.Color{R:%d, G: %d, B: %d, A: %d},\n   \t}\n", txt.Id, identifier, txt.Id, txt.Content, defaultSize, txt.X, txt.Y, txt.W, txt.H, AlignmentToString(txt.Alignment), txt.TextColor.R, txt.TextColor.G, txt.TextColor.B, txt.TextColor.A)
	return identifier, construction
}

func constructButton(btn ge_go_sdl2.Button, viewId string, declarations *strings.Builder) (string, string, string) {
	btnLabelId := btn.Id + "_label"
	btn.ContentLabel.Id = btnLabelId
	btn.ContentLabel.Alignment = ge_go_sdl2.Center
	btn.ContentLabel.W = btn.W
	btnTextIdentifier, btnTextConstruction := constructText(btn.ContentLabel, viewId)
	declarations.WriteString(fmt.Sprintf("%s %s\n", btnTextIdentifier, "ge_go_sdl2.Text"))
	onClickIdentifier := fmt.Sprintf("%s_button_%s_onclick", viewId, btn.Id)
	onClickChannelConstruction := fmt.Sprintf("%s := make(chan string)\n", onClickIdentifier)
	defaultColor := ifElseAssign(btn.BgColor == 0, "DefaultButtonColor", fmt.Sprintf("0x%x", btn.BgColor))
	DefaultBorderColor := ifElseAssign(btn.BorderColor == 0, "DefaultButtonBorderColor", fmt.Sprintf("0x%x", btn.BorderColor))
	identifier := fmt.Sprintf("%s_button_%s", viewId, btn.Id)
	btnDeclaration := fmt.Sprintf(
		"%s = ge_go_sdl2.Button{\nId: \"%s\",\nX: %d,\nY: %d,\nH: %d,\nW: %d,\nContent: \"%s\",\nBgColor: %s,\nBorderColor: %s,\nContentLabel: %s ,\n OnClick: %s,\n\t}\n",
		identifier, btn.Id, btn.X, btn.Y, btn.H, btn.W, btn.Content, defaultColor, DefaultBorderColor, btnTextIdentifier, onClickIdentifier)
	onClickFunctionIdentifier := fmt.Sprintf("OnClick_%s", identifier)
	handleClickConstruction := fmt.Sprintf("ge_go_sdl2.HandleAsCallback(%s, %s)\n", onClickIdentifier, onClickFunctionIdentifier)
	construction := fmt.Sprintf("// %s button construction \n%s %s %s %s\n", btn.Id, btnTextConstruction, onClickChannelConstruction, btnDeclaration, handleClickConstruction)

	return identifier, construction, onClickFunctionIdentifier
}

func constructTextField(textField ge_go_sdl2.TextField, viewId string) (string, string, string) {

	defaultSize := ifElseAssign(textField.Size == 0, "DefaultTextSize", strconv.Itoa(textField.Size))
	defaultColor := ifElseAssign(textField.BgColor == 0, "DefaultTextFieldColor", fmt.Sprintf("0x%x", textField.BgColor))
	DefaultBorderColor := ifElseAssign(textField.BorderColor == 0, "DefaultTextFieldBorderColor", fmt.Sprintf("0x%x", textField.BorderColor))
	onChangedChannelIdentifier := fmt.Sprintf("textfield_%s_onchange", textField.Id)
	onChangedChannelConstruction := fmt.Sprintf("%s := make(chan string)\n", onChangedChannelIdentifier)
	identifier := fmt.Sprintf("%s_textField_%s", viewId, textField.Id)
	textFieldConstruction := fmt.Sprintf(
		"%s = ge_go_sdl2.TextField{\nX: %d,\nY:%d,\nH: %d,\nW:%d,\nId:\"%s\",\nSize:%s,\nValue:\"%s\",\nBgColor: %s,\nBorderColor: %s,\nFont:DefaultFont,\nOnChanged:   textfield_%s_onchange,\n\t}\n",
		identifier, textField.X, textField.Y, textField.H, textField.W, textField.Id, defaultSize, textField.Value, defaultColor, DefaultBorderColor, textField.Id)

	onNewValueFunctionIdentifier := fmt.Sprintf("OnNewValue_%s", identifier)
	handleNewValueConstruction := fmt.Sprintf("ge_go_sdl2.HandleAsCallbackArg(%s, %s)\n", onChangedChannelIdentifier, onNewValueFunctionIdentifier)

	construction := fmt.Sprintf("// %s text field construction \n\n%s %s %s\n", textField.Id, onChangedChannelConstruction, textFieldConstruction, handleNewValueConstruction)

	return identifier, construction, onNewValueFunctionIdentifier
}
