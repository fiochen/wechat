package contacts

//Attr 对外成员的属性，支持文本、网页、小程序三种类型
type Attr struct {
	Type        int              `json:"type" xml:"Type"` // 属性类型: 0-文本 1-网页 2-小程序
	Name        string           `json:"name" xml:"Name"` // 属性名称： 需要先确保在管理端有创建该属性，否则会忽略
	Text        *AttrText        `json:"text,omitempty" xml:"Text"`
	Web         *AttrWeb         `json:"web,omitempty" xml:"Web"`
	MiniProgram *AttrMiniProgram `json:"miniprogram,omitempty" xml:"MiniProgram"`
}

//AttrText 文本类型的属性
type AttrText struct {
	Value string `json:"value,omitempty" xml:"Value"` // 文本属性内容,长度限制12个UTF8字符
}

//AttrWeb 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空
type AttrWeb struct {
	URL   string `json:"url,omitempty" xml:"Url"`     // 网页的url,必须包含http或者https头
	Title string `json:"title,omitempty" xml:"Title"` // 网页的展示标题,长度限制12个UTF8字符
}

//AttrMiniProgram 小程序类型的属性，appid和title字段要么同时为空表示清除该属性，要么同时不为空
type AttrMiniProgram struct {
	AppID    string `json:"appid,omitempty" xml:"AppID"`       // 小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
	Title    string `json:"title,omitempty" xml:"Title"`       // 小程序的展示标题,长度限制12个UTF8字符
	PagePath string `json:"pagepath,omitempty" xml:"PagePath"` // 小程序的页面路径
}

//ExternalAttr 属性列表
type ExternalAttr struct {
	Attrs []Attr `json:"attrs,omitempty"`
}

//ExternalProfile 成员对外属性
type ExternalProfile struct {
	ExternalCorpName string          `json:"external_corp_name,omitempty"` // 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	WechatChannels   *WechatChannels `json:"wechat_channels,omitempty"`    // 视频号名字（设置后，成员将对外展示该视频号）。须从企业绑定到企业微信的视频号中选择，可在“我的企业”页中查看绑定的视频号。
	ExternalAttrs    []Attr          `json:"external_attr,omitempty"`
}

//WechatChannels 视频号名字
type WechatChannels struct {
	Nickname string `json:"nickname,omitempty"`
}
