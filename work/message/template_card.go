package message

//CommonTemplateCard 模板卡片消息
type CommonTemplateCard struct {
	CardType              CDATA               `xml:"CardType"`                        // 模板卡片类型
	Source                Source              `xml:"Source,omitempty"`                // 卡片来源样式信息，不需要来源样式可不填写
	MainTitle             CDATA               `xml:"MainTitle"`                       // 一级标题
	SubTitleText          CDATA               `xml:"SubTitleText,omitempty"`          // 二级普通文本
	HorizontalContentList []HorizontalContent `xml:"HorizontalContentList,omitempty"` // 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
	JumpList              []Jump              `xml:"JumpList,omitempty"`              // 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
}

//NewCommonTemplateCard 初始化模板卡片消息
func NewCommonTemplateCard(mainTitle string, source Source, subTitleText string, horizontalContentList []HorizontalContent, jumpList []Jump) *CommonTemplateCard {
	card := new(CommonTemplateCard)
	card.Source = source
	card.MainTitle = CDATA(mainTitle)
	card.SubTitleText = CDATA(subTitleText)
	card.HorizontalContentList = horizontalContentList
	card.JumpList = jumpList
	return card
}

//Source source
type Source struct {
	IconURL CDATA `xml:"IconUrl"` // 来源图片的url
	Desc    CDATA `xml:"Desc"`    // 来源图片的描述
}

//MainTitle MainTitle
type MainTitle struct {
	Title CDATA `xml:"Title"` // 一级标题
	Desc  CDATA `xml:"Desc"`  // 标题辅助信息
}

//HorizontalContent HorizontalContent
type HorizontalContent struct {
	KeyName CDATA `xml:"KeyName"`           // 二级标题，必填
	Value   CDATA `xml:"Value"`             // 二级文本，如果HorizontalContentList.Type是2，该字段代表文件名称（要包含文件类型）
	Type    int   `xml:"Type"`              // 链接类型，0或不填或错填代表不是链接，1 代表跳转url，2 代表下载附件
	URL     CDATA `xml:"Url,omitempty"`     // 链接跳转的url，HorizontalContentList.Type是1时必填
	MediaID CDATA `xml:"MediaId,omitempty"` // 附件的media_id，HorizontalContentList.Type是2时必填
}

//VerticalContent VerticalContent
type VerticalContent struct {
	Title CDATA `xml:"Title"` // 卡片二级标题
	Desc  CDATA `xml:"Desc"`  // 卡片二级内容
}

//Jump jump
type Jump struct {
	Title    CDATA `xml:"Title"`              // 跳转链接样式的文案内容，必填
	Type     int   `xml:"Type,omitempty"`     // 跳转链接类型，0或不填或错填代表不是链接，1 代表跳转url，2 代表跳转小程序
	URL      CDATA `xml:"Url,omitempty"`      // 跳转链接的url，JumpList.Type是1时必填
	AppID    CDATA `xml:"AppId,omitempty"`    // 跳转链接的小程序的appid，JumpList.Type是2时必填
	PagePath CDATA `xml:"PagePath,omitempty"` // 跳转链接的小程序的pagepath，JumpList.Type是2时选填
}

//CardAction Card Action
type CardAction struct {
	Type     int   `xml:"Type"`               // 跳转事件类型，0或不填或错填代表不是链接，1 代表跳转url，2 代表下载附件
	URL      CDATA `xml:"Url,omitempty"`      // 跳转事件的url，CardAction.Type是1时必填
	AppID    CDATA `xml:"AppId,omitempty"`    // 跳转事件的小程序的appid，CardAction.Type是2时必填
	PagePath CDATA `xml:"PagePath,omitempty"` // 跳转事件的小程序的pagepath，CardAction.Type是2时选填
}

//SubmitButton SubmitButton
type SubmitButton struct {
	Title CDATA `xml:"Title,omitempty"` // 按钮文案，建议不超过10个字，不填默认为提交
	Key   CDATA `xml:"Key"`             // 提交按钮的key，会产生回调事件将本参数作为EventKey返回，最长支持1024字节，必填
}

//ButtonItem ButtonItem
type ButtonItem struct {
	Title CDATA `xml:"Title"`           // 按钮文案
	Style int   `xml:"Style,omitempty"` // 按钮样式，目前可填1~4，不填或错填默认1
	Key   CDATA `xml:"Key"`             // 按钮key值，用户点击后，会产生回调事件，回调事件会带上该key值，最长支持1024字节
}

//OptionItem Option Item
type OptionItem struct {
	ID   CDATA `xml:"Id"`   // 下拉式的选择器选项的id，用户提交选项后，会产生回调事件，回调事件会带上该id值表示该选项，最长支持128字节
	Text CDATA `xml:"Text"` // 下拉式的选择器选项的文案
}

//SelectItem Select Item
type SelectItem struct {
	QuestionKey CDATA        `xml:"QuestionKey"`       // 下拉式的选择器题目的key，用户提交选项后，会产生回调事件，回调事件会带上该key值表示该题，最长支持1024字节
	Title       CDATA        `xml:"Title"`             // 下拉式的选择器上面的Title
	SelectedID  CDATA        `xml:"SelectedId"`        // 下拉式的选择器默认选定的选项
	Disable     bool         `xml:"Disable,omitempty"` // 是否可以选择状态
	OptionList  []OptionItem `xml:"OptionList"`        // 选项列表，下拉选项不超过 10 个
}
