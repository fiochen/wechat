package message

//TemplateCardNews 图文展示型模板卡片消息
type TemplateCardNews struct {
	CommonToken
	TemplateCard TemplateCardNewsContent `xml:"TemplateCard"`
}

//TemplateCardNewsContent 图文展示型内容
type TemplateCardNewsContent struct {
	CommonTemplateCard
	CardAction          CardAction        `xml:"CardAction"`                    // 整体卡片的点击跳转事件，必填
	CardImage           CardImage         `xml:"CardImage"`                     // 图片展示样式
	VerticalContentList []VerticalContent `xml:"VerticalContentList,omitempty"` // 卡片二级内容，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过4
}

//CardImage 图片展示样式
type CardImage struct {
	URL         CDATA   `xml:"Url"`                   // 图片的url
	AspectRatio float32 `xml:"AspectRatio,omitempty"` // 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
}

//NewTemplateCardNews 初始化图文展示型模板卡片消息
func NewTemplateCardNews(token CommonToken, template CommonTemplateCard, cardAction CardAction, cardImage CardImage, verticalContents []VerticalContent) *TemplateCardNews {
	template.CardType = TemplateCardTypeNews
	card := new(TemplateCardNews)
	card.CommonToken = token
	card.TemplateCard = TemplateCardNewsContent{
		CommonTemplateCard:  template,
		CardAction:          cardAction,
		CardImage:           cardImage,
		VerticalContentList: verticalContents,
	}
	return card
}
