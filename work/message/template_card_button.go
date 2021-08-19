package message

//TemplateCardButton 按钮交互型模板卡片消息
type TemplateCardButton struct {
	CommonToken
	TemplateCard TemplateCardButtonContent `xml:"TemplateCard"`
}

//TemplateCardButtonContent 按钮交互型内容
type TemplateCardButtonContent struct {
	CommonTemplateCard
	CardAction   CardAction   `xml:"CardAction,omitempty"`   // 整体卡片的点击跳转事件
	TaskID       CDATA        `xml:"TaskId"`                 // 任务id，同一个应用任务id不能重复，只能由数字、字母和“_-@”组成，最长128字节
	ButtonList   []ButtonItem `xml:"ButtonList,omitempty"`   // 按钮列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
	ButtonColumn int          `xml:"ButtonColumn,omitempty"` // todo: 文档说明没有，但是样例有
	ReplaceText  CDATA        `xml:"ReplaceText,omitempty"`  // 按钮替换文案，填写本字段后会展现灰色不可点击按钮
}

//NewTemplateCardButton 初始化按钮交互型模板卡片消息
func NewTemplateCardButton(token CommonToken, template CommonTemplateCard, cardAction CardAction, taskID string, buttonLs []ButtonItem, buttonColumn int, replaceText string) *TemplateCardButton {
	template.CardType = TemplateCardTypeButton
	card := new(TemplateCardButton)
	card.CommonToken = token
	card.TemplateCard = TemplateCardButtonContent{
		CommonTemplateCard: template,
		CardAction:         cardAction,
		TaskID:             CDATA(taskID),
		ButtonList:         buttonLs,
		ButtonColumn:       buttonColumn,
		ReplaceText:        CDATA(replaceText),
	}
	return card
}
