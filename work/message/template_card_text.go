package message

//TemplateCardText 文本通知型模板卡片消息
type TemplateCardText struct {
	CommonToken
	TemplateCard TemplateCardTextContent `xml:"TemplateCard"`
}

//TemplateCardTextContent 文本通知型内容
type TemplateCardTextContent struct {
	CommonTemplateCard
	CardAction      CardAction      `xml:"CardAction"`      // 整体卡片的点击跳转事件，必填
	EmphasisContent EmphasisContent `xml:"EmphasisContent"` // 关键数据样式
}

//EmphasisContent 关键数据样式
type EmphasisContent struct {
	Title CDATA `xml:"Title"` // 关键数据样式的数据内容
	Desc  CDATA `xml:"Desc"`  // 关键数据样式的数据描述内容
}

//NewTemplateCardText 初始化文本通知型模板卡片消息
func NewTemplateCardText(token CommonToken, template CommonTemplateCard, cardAction CardAction, emphasisContent EmphasisContent) *TemplateCardText {
	template.CardType = TemplateCardTypeText
	card := new(TemplateCardText)
	card.CommonToken = token
	card.TemplateCard = TemplateCardTextContent{
		CommonTemplateCard: template,
		CardAction:         cardAction,
		EmphasisContent:    emphasisContent,
	}
	return card
}
