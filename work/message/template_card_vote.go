package message

//TemplateCardVote 投票选择型模板卡片消息
type TemplateCardVote struct {
	CommonToken
	TemplateCard TemplateCardVoteContent `xml:"TemplateCard"`
}

//TemplateCardVoteContent 投票选择型内容
type TemplateCardVoteContent struct {
	CommonTemplateCard
	TaskID       CDATA        `xml:"TaskId"`                // 任务id，同一个应用任务id不能重复，只能由数字、字母和“_-@”组成，最长128字节
	CheckBox     CheckBox     `xml:"CheckBox"`              // 选择题样式，仅”vote_interaction”要填该字段
	SubmitButton SubmitButton `xml:"SubmitButton"`          // 提交按钮样式
	ReplaceText  CDATA        `xml:"ReplaceText,omitempty"` // 按钮替换文案，填写本字段后会展现灰色不可点击按钮
}

//CheckBox 选择题样式
type CheckBox struct {
	QuestionKey CDATA        `xml:"QuestionKey,omitempty"` // 选择题key值，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
	Disable     bool         `xml:"Disable,omitempty"`     // 是否可以选择状态
	Mode        int          `xml:"Mode,omitempty"`        // 选择题模式，单选：0，多选：1，不填或错填默认0
	OptionList  []OptionItem `xml:"OptionList"`            // 选项列表，下拉选项不超过 10 个
}

//NewTemplateCardVote 初始化投票选择型模板卡片消息
func NewTemplateCardVote(token CommonToken, template CommonTemplateCard, taskID string, checkBox CheckBox, submitButton SubmitButton, replaceText string) *TemplateCardVote {
	template.CardType = TemplateCardTypeVote
	card := new(TemplateCardVote)
	card.CommonToken = token
	card.TemplateCard = TemplateCardVoteContent{
		CommonTemplateCard: template,
		TaskID:             CDATA(taskID),
		CheckBox:           checkBox,
		SubmitButton:       submitButton,
		ReplaceText:        CDATA(replaceText),
	}
	return card
}
