package gforms

import (
	"bytes"
)

type TextWidget struct {
	Attrs map[string]string
	Widget
}

type parameter struct {
	Name  string
	Value string
	Attrs map[string]string
}

func (self *TextWidget) html(field Field, vs ...string) string {
	var buffer bytes.Buffer
	var v string
	if len(vs) > 0 {
		v = vs[0]
	}
	Template.ExecuteTemplate(&buffer, "TextWidget", parameter{
		Attrs: self.Attrs,
		Name:  field.GetName(),
		Value: v,
	})
	return buffer.String()
}

func (self *TextWidget) Validate(value interface{}) error {
	return nil
}

func NewTextWidget(attrs map[string]string) *TextWidget {
	w := new(TextWidget)
	w.Attrs = attrs
	return w
}
