package uic_go_sdl2

import (
	"fmt"
	ge_go_sdl2 "github.com/DanielHauge/ge-go-sdl2"
	"strconv"
)

func parseFromQtFormat(qtUi QtUi) ge_go_sdl2.View {
	return parseView(qtUi.Widget)
}

func parseWidget(qtWidget Widget) interface{} {
	switch qtWidget.Class {
	case "QPushButton":
		return parseButton(qtWidget)
	case "QTextEdit":
		return parseTextField(qtWidget)
	case "QLabel":
		return parseLabel(qtWidget)
	case "QGroupBox":
		return parseContainer(qtWidget)
	default:
		fmt.Println("Does not support parsing of class: " + qtWidget.Class)
		return nil
	}
}

func parseButton(btn Widget) ge_go_sdl2.Button {
	button := ge_go_sdl2.Button{}
	for _, prop := range btn.Property {
		switch prop.Name {
		case "geometry":
			button.X = parseInt32(prop.Rect.X)
			button.Y = parseInt32(prop.Rect.Y)
			button.W = parseInt32(prop.Rect.Width)
			button.H = parseInt32(prop.Rect.Height)
			break
		case "text":
			button.Content = prop.String
			break
		default:
			fmt.Println("Does not support property: " + prop.Name)
		}
	}
	button.Id = btn.Name
	return button
}

func parseTextField(txtField Widget) ge_go_sdl2.TextField {
	tf := ge_go_sdl2.TextField{}
	for _, prop := range txtField.Property {
		switch prop.Name {
		case "geometry":
			tf.X = parseInt32(prop.Rect.X)
			tf.Y = parseInt32(prop.Rect.Y)
			tf.W = parseInt32(prop.Rect.Width)
			tf.H = parseInt32(prop.Rect.Height)
			break
		case "text":
			tf.Value = prop.String
			break
		default:
			fmt.Println("Does not support property: " + prop.Name)
		}
	}
	tf.Id = txtField.Name
	return tf
}

func parseLabel(label Widget) ge_go_sdl2.Text {
	t := ge_go_sdl2.Text{}
	for _, prop := range label.Property {
		switch prop.Name {
		case "geometry":
			t.X = parseInt32(prop.Rect.X)
			t.Y = parseInt32(prop.Rect.Y)
			break
		case "text":
			t.Content = prop.String
			break
		default:
			fmt.Println("Does not support property: " + prop.Name)
		}
	}
	t.Id = label.Name
	return t
}

func parseContainer(cont Widget) ge_go_sdl2.Container {
	c := ge_go_sdl2.Container{}
	for _, prop := range cont.Property {
		switch prop.Name {
		case "geometry":
			c.X = parseInt32(prop.Rect.X)
			c.Y = parseInt32(prop.Rect.Y)
			break
		case "title":
			c.ViewId = prop.String
			break
		default:
			fmt.Println("Does not support property: " + prop.Name)
		}
	}
	c.Id = cont.Name
	return c
}

func parseWidgets(qtWidgets []Widget) []interface{} {
	var parsedWidgets []interface{}
	for _, w := range qtWidgets {
		widget := parseWidget(w)
		if widget != nil {
			parsedWidgets = append(parsedWidgets, widget)
		}
	}
	return parsedWidgets
}

func parseView(qtWidget Widget) ge_go_sdl2.View {
	view := ge_go_sdl2.View{}
	for _, prop := range qtWidget.Property {
		switch prop.Name {
		case "geometry":
			view.X = parseInt32(prop.Rect.X)
			view.Y = parseInt32(prop.Rect.Y)
			view.W = parseInt32(prop.Rect.Width)
			view.H = parseInt32(prop.Rect.Height)
			break
		case "windowTitle":
			view.Id = prop.String
			break
		default:
			fmt.Println("Does not support property: " + prop.Name)
		}
	}
	var children []interface{}
	for _, child := range qtWidget.Widget {
		switch child.Class {
		case "QWidget":
			children = parseWidgets(child.Widget)
		default:
			fmt.Sprintf("Does not support parsing of class: " + child.Class)
		}
	}
	view.Children = children
	return view
}

func parseInt32(s string) int32 {
	n, err := strconv.ParseInt(s, 10, 64)
	nilOrPanic(err)
	return int32(n)
}
