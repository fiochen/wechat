package message

//TemplateCardMultipleInteraction 多项选择型模板卡片消息
type TemplateCardMultipleInteraction struct {
	CommonToken
	TemplateCard TemplateCardMultipleInteractionContent `xml:"TemplateCard"`
}

//TemplateCardMultipleInteractionContent 多项选择型内容
type TemplateCardMultipleInteractionContent struct {
	CommonTemplateCard
	TaskID       CDATA        `xml:"TaskId"`                // 任务id，同一个应用任务id不能重复，只能由数字、字母和“_-@”组成，最长128字节
	SelectList   []SelectItem `xml:"SelectList"`            // 下拉式的选择器列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，一个消息最多支持 3 个选择器
	SubmitButton SubmitButton `xml:"SubmitButton"`          // 提交按钮样式
	ReplaceText  CDATA        `xml:"ReplaceText,omitempty"` // 按钮替换文案，填写本字段后会展现灰色不可点击按钮
}

//NewTemplateCardMultipleInteraction 初始化多项选择型模板卡片消息
func NewTemplateCardMultipleInteraction(token CommonToken, template CommonTemplateCard, taskID string, selectList []SelectItem, submitButton SubmitButton, replaceText string) *TemplateCardMultipleInteraction {
	template.CardType = TemplateCardTypeMultipleInteraction
	card := new(TemplateCardMultipleInteraction)
	card.CommonToken = token
	card.TemplateCard = TemplateCardMultipleInteractionContent{
		CommonTemplateCard: template,
		TaskID:             CDATA(taskID),
		SelectList:         selectList,
		SubmitButton:       submitButton,
		ReplaceText:        CDATA(replaceText),
	}
	return card
}
