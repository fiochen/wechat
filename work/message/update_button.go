package message

//UpdateButton 回复更新按钮文案消息
type UpdateButton struct {
	CommonToken

	Button UpdateButtonContent `xml:"Button"`
}

//UpdateButtonContent 回复更新按钮内容
type UpdateButtonContent struct {
	ReplaceName CDATA `xml:"ReplaceName"` // 点击卡片按钮后显示的按钮名称
}

//NewUpdateButton 初始化回复更新按钮文案消息
func NewUpdateButton(replaceName string) *UpdateButton {
	button := new(UpdateButton)
	button.Button = UpdateButtonContent{
		ReplaceName: CDATA(replaceName),
	}
	return button
}
