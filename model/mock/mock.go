package mock

import "github.com/srdjanimperator/genui/model"

func GetMockForms() []model.GuiForm {
	forms := make([]model.GuiForm, 0)

	forms = append(forms, model.GuiForm{
		Code:  "#CaseCreate",
		Title: "Create New Case",
	})

	forms = append(forms, model.GuiForm{
		Code:  "#CaseFile",
		Title: "Case Filing",
	})

	forms = append(forms, model.GuiForm{
		Code:  "#BillEdit",
		Title: "Edit Bill",
	})

	return forms
}

func GetMockCtrls() []model.GuiFormCtrl {
	ctrls := make([]model.GuiFormCtrl, 0)

	ctrls = append(ctrls, model.GuiFormCtrl{
		AttrName: "fileNumber",
		Type:     model.TextInput,
	})

	ctrls = append(ctrls, model.GuiFormCtrl{
		AttrName: "title",
		Type:     model.TextInput,
	})

	ctrls = append(ctrls, model.GuiFormCtrl{
		AttrName: "processor",
		Type:     model.TextInput,
	})

	ctrls = append(ctrls, model.GuiFormCtrl{
		AttrName: "filed",
		Type:     model.Checkbox,
	})

	ctrls = append(ctrls, model.GuiFormCtrl{
		AttrName: "caseTypeId",
		Type:     model.Dropdown,
	})

	return ctrls
}
