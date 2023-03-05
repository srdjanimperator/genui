package model

type GuiFormCtrlType string

const (
	// TextInput is the most common type of form control: simple text field.
	TextInput GuiFormCtrlType = "TEXT_FIELD"
	// Dropdown is a type of form control - a simple dropdown for small option collections.
	Dropdown GuiFormCtrlType = "DROPDOWN"
	// Checkbox is a type of form control with simple true/false or 1/0 output value.
	Checkbox GuiFormCtrlType = "CHECK_BOX"
)

// GuiForm
type GuiForm struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

type GuiFormCtrl struct {
	AttrName string          `json:"attrName"`
	Type     GuiFormCtrlType `json:"type"`
}
