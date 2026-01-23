package api

import "fmt"

type Widget struct {
	PrettyName         string   `json:"prettyName"`
	Id                 string   `json:"id"`
	WidgetType         string   `json:"widgetType"`
	WidgetValueType    string   `json:"widgetValueType"`
	MinValue           string   `json:"minValue"`
	MaxValue           string   `json:"maxValue"`
	DefaultValue       string   `json:"defaultValue"`
	StepAmount         string   `json:"stepAmount"`
	CurrentValue       string   `json:"currentValue"`
	Target             func()   `json:"-"` // this is a function that will be called when the widget is interacted with if the type is a button
	ValuePointerInt    *int     `json:"-"`
	ValuePointerFloat  *float64 `json:"-"`
	ValuePointerString *string  `json:"-"`
	ValuePointerBool   *bool    `json:"-"`
}

func NewFloatSliderWidget(prettyName, id, minValue, maxValue, defaultValue, stepAmount string, valuePointer *float64) Widget {
	return Widget{
		PrettyName:        prettyName,
		Id:                id,
		WidgetType:        "slider",
		WidgetValueType:   "float",
		MinValue:          minValue,
		MaxValue:          maxValue,
		DefaultValue:      defaultValue,
		StepAmount:        stepAmount,
		ValuePointerFloat: valuePointer,
	}
}

func NewIntSliderWidget(prettyName, id, minValue, maxValue, defaultValue, stepAmount string, valuePointer *int) Widget {
	return Widget{
		PrettyName:      prettyName,
		Id:              id,
		WidgetType:      "slider",
		WidgetValueType: "int",
		MinValue:        minValue,
		MaxValue:        maxValue,
		DefaultValue:    defaultValue,
		StepAmount:      stepAmount,
		ValuePointerInt: valuePointer,
	}
}

func NewButtonWidget(prettyName, id string, target func()) Widget {
	return Widget{
		PrettyName: prettyName,
		Id:         id,
		WidgetType: "button",
		Target:     target,
	}
}

func NewMouseXClickedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-x-clicked",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseYClickedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-y-clicked",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseClickedHook(valuePointer *bool) Widget {
	return Widget{
		WidgetType:       "background",
		Id:               "mouse-clicked",
		WidgetValueType:  "bool",
		ValuePointerBool: valuePointer,
	}
}

func NewMouseXMovedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-x-moved",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseYMovedHook(valuePointer *float64) Widget {
	return Widget{
		WidgetType:        "background",
		Id:                "mouse-y-moved",
		WidgetValueType:   "float",
		ValuePointerFloat: valuePointer,
	}
}

func NewMouseMovedHook(valuePointer *bool) Widget {
	return Widget{
		WidgetType:       "background",
		Id:               "mouse-moved",
		WidgetValueType:  "bool",
		ValuePointerBool: valuePointer,
	}
}

func (w *Widget) getCurrentValue() string {
	if w.ValuePointerInt != nil {
		return fmt.Sprintf("%d", *w.ValuePointerInt)
	}
	if w.ValuePointerFloat != nil {
		value := fmt.Sprintf("%f", *w.ValuePointerFloat)
		// remove trailing zeros
		for len(value) > 0 && value[len(value)-1] == '0' {
			value = value[:len(value)-1]
		}
		// remove trailing decimal point
		if len(value) > 0 && value[len(value)-1] == '.' {
			value = value[:len(value)-1]
		}
		return value
	}
	if w.ValuePointerString != nil {
		return *w.ValuePointerString
	}
	if w.ValuePointerBool != nil {
		return fmt.Sprintf("%t", *w.ValuePointerBool)
	}
	return w.DefaultValue
}
