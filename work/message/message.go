package message

import (
	"encoding/xml"
	"github.com/silenceper/wechat/v2/work/contacts"
)

// MsgType 基本消息类型
type MsgType string

const (
	//MsgTypeText 表示文本消息
	MsgTypeText MsgType = "text"
	//MsgTypeImage 表示图片消息
	MsgTypeImage = "image"
	//MsgTypeVoice 表示语音消息
	MsgTypeVoice = "voice"
	//MsgTypeVideo 表示视频消息
	MsgTypeVideo = "video"
	//MsgTypeFile 表示文件消息
	MsgTypeFile = "file"
	//MsgTypeLocation 表示坐标消息[限接收]
	MsgTypeLocation = "location"
	//MsgTypeLink 表示链接消息[限接收]
	MsgTypeLink = "link"
	//MsgTypeTextCard 表示文本卡片消息
	MsgTypeTextCard = "textcard"
	//MsgTypeNews 表示图文消息
	MsgTypeNews = "news"
	//MsgTypeMPNews 表示mpnews类型的图文消息
	MsgTypeMPNews = "mpnews"
	//MsgTypeMarkdown 表示markdown消息
	MsgTypeMarkdown = "markdown"
	//MsgTypeMiniProgramNotice 表小程序通知消息
	MsgTypeMiniProgramNotice = "miniprogram_notice"
	//MsgTypeTemplateCard 表模板卡片消息
	MsgTypeTemplateCard = "template_card"

	//MsgTypeEvent 表示事件推送消息
	MsgTypeEvent = "event"

	//MsgTypeUpdateButton 表示更新按钮点击文案
	MsgTypeUpdateButton = "update_button"
	//MsgTypeUpdateTemplateCard 表示更新整张卡片
	MsgTypeUpdateTemplateCard = "update_template_card"
)

const (
	//EventChangeContact 表示通讯录变更事件(成员变更/部门变更/标签变更)
	EventChangeContact = "change_contact"
	//EventBatchJobResult 表示异步任务完成通知
	EventBatchJobResult = "batch_job_result"
	//EventChangeExternalContact 表示客户变更事件
	EventChangeExternalContact = "change_external_contact"
	//EventChangeExternalChat 客户群事件
	EventChangeExternalChat = "change_external_chat"
)

const (
	//ChangeTypeCreateUser 新增成员事件
	ChangeTypeCreateUser = "create_user"
	//ChangeTypeUpdateUser 更新成员事件
	ChangeTypeUpdateUser = "update_user"
	//ChangeTypeDeleteUser 删除成员事件
	ChangeTypeDeleteUser = "delete_user"
	//ChangeTypeCreateParty 新增部门事件
	ChangeTypeCreateParty = "create_party"
	//ChangeTypeUpdateParty 更新部门事件
	ChangeTypeUpdateParty = "update_party"
	//ChangeTypeDeleteParty 删除部门事件
	ChangeTypeDeleteParty = "delete_party"
	//ChangeTypeUpdateTag 标签成员变更事件
	ChangeTypeUpdateTag = "update_tag"

	//ChangeTypeAddExternalContact 添加企业客户事件
	ChangeTypeAddExternalContact = "add_external_contact"
	//ChangeTypeEditExternalContact 编辑企业客户事件
	ChangeTypeEditExternalContact = "edit_external_contact"
	//ChangeTypeAddHalfExternalContact 外部联系人免验证添加成员事件
	ChangeTypeAddHalfExternalContact = "add_half_external_contact"
	//ChangeTypeDelExternalContact 删除企业客户事件
	ChangeTypeDelExternalContact = "del_external_contact"
	//ChangeTypeDelFollowUser 删除跟进成员事件
	ChangeTypeDelFollowUser = "del_follow_user"
	//ChangeTypeTransferFail 客户接替失败事件
	ChangeTypeTransferFail = "transfer_fail"
	//ChangeTypeCreate 客户群创建事件
	ChangeTypeCreate = "create"
	//ChangeTypeUpdate 客户群变更事件
	ChangeTypeUpdate = "update"
	//ChangeTypeDismiss 客户群解散事件
	ChangeTypeDismiss = "dismiss"
)

const (
	//JobTypeSyncUser 增量更新成员
	JobTypeSyncUser = "sync_user"
	//JobTypeReplaceUser 全量覆盖成员
	JobTypeReplaceUser = "replace_user"
	//JobTypeInviteUser 邀请成员关注
	JobTypeInviteUser = "invite_user"
	//JobTypeReplaceParty 全量覆盖部门
	JobTypeReplaceParty = "replace_party"
)

//TemplateCardType 模板卡片类型
const (
	// TemplateCardTypeText 文本通知型
	TemplateCardTypeText = "text_notice"
	//TemplateCardTypeNews 图文展示型
	TemplateCardTypeNews = "news_notice"
	//TemplateCardTypeButton 按钮交互型
	TemplateCardTypeButton = "button_interaction"
	//TemplateCardTypeVote 投票选择型
	TemplateCardTypeVote = "vote_interaction"
	//TemplateCardTypeMultipleInteraction 多项选择型
	TemplateCardTypeMultipleInteraction = "multiple_interaction"
)

//MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
	CommonToken

	Event      string `xml:"Event"`      // 事件的类型
	EventKey   string `xml:"EventKey"`   // 事件KEY值
	ChangeType string `xml:"ChangeType"` // 变更类型
	AgentID    int    `xml:"AgentID"`    // 企业应用的id，整型。可在应用的设置页面查看

	// 通讯录 - 成员变更相关
	UserID         string          `xml:"UserID"`         // 成员UserID
	NewUserID      string          `xml:"NewUserID"`      // 新的UserID，变更时推送（userid由系统生成时可更改一次）
	Name           string          `xml:"Name"`           // 名称
	DepartmentIDs  string          `xml:"Department"`     // 成员部门列表，仅返回该应用有查看权限的部门id
	MainDepartment int             `xml:"MainDepartment"` // 主部门
	IsLeaderInDept string          `xml:"IsLeaderInDept"` // 表示所在部门是否为上级，0-否，1-是，顺序与Department字段的部门逐一对应
	Position       string          `xml:"Position"`       // 职位信息。长度为0~64个字节
	Mobile         string          `xml:"Mobile"`         // 手机号码
	Gender         int             `xml:"Gender"`         // 性别，1表示男性，2表示女性
	Email          string          `xml:"Email"`          // 邮箱
	Status         int             `xml:"Status"`         // 激活状态：1=已激活 2=已禁用 4=未激活 已激活代表已激活企业微信或已关注微工作台（原企业号）5=成员退出
	Avatar         string          `xml:"Avatar"`         // 头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可。
	Alias          string          `xml:"Alias"`          // 成员别名
	Telephone      string          `xml:"Telephone"`      // 座机
	Address        string          `xml:"Address"`        // 地址
	ExtAttr        []contacts.Attr `xml:"ExtAttr>Item"`   // 扩展属性

	// 通讯录 - 部门变更相关
	ID       interface{} `xml:"Id"`       // 部门Id 或 标签或标签组的ID
	ParentID string      `xml:"ParentId"` // 父部门id
	Order    int         `xml:"Order"`    // 部门排序

	// 通讯录 - 标签变更相关
	TagID         int    `xml:"TagId"`         // 标签Id
	AddUserItems  string `xml:"AddUserItems"`  // 标签中新增的成员userid列表，用逗号分隔
	DelUserItems  string `xml:"DelUserItems"`  // 标签中删除的成员userid列表，用逗号分隔
	AddPartyItems string `xml:"AddPartyItems"` // 标签中新增的部门id列表，用逗号分隔
	DelPartyItems string `xml:"DelPartyItems"` // 标签中删除的部门id列表，用逗号分隔

	// 异步任务结果通知相关
	JobID   string `xml:"JobId"`   // 异步任务id，最大长度为64字符
	JobType string `xml:"JobType"` // 操作类型，字符串，目前分别有：sync_user(增量更新成员)、 replace_user(全量覆盖成员）、invite_user(邀请成员关注）、replace_party(全量覆盖部门)
	ErrCode int    `xml:"ErrCode"` // 返回码
	ErrMsg  string `xml:"ErrMsg"`  // 对返回码的文本描述内容

	// 客户变更事件相关
	ExternalUserID string `xml:"ExternalUserID"` // 外部联系人的userid，注意不是企业成员的帐号
	State          string `xml:"State"`          // 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	WelcomeCode    string `xml:"WelcomeCode"`    // 欢迎语code，可用于发送欢迎语
	Source         string `xml:"Source"`         // 删除客户的操作来源，DELETE_BY_TRANSFER表示此客户是因在职继承自动被转接成员删除
	FailReason     string `xml:"FailReason"`     // 接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
	ChatID         string `xml:"ChatId"`         // 群ID
	UpdateDetail   string `xml:"UpdateDetail"`   // 变更详情。目前有以下几种：<add_member : 成员入群>, <del_member : 成员退群>, <change_owner : 群主变更>, <change_name : 群名变更>, <change_notice : 群公告变更>
	JoinScene      int    `xml:"JoinScene"`      // 当是成员入群时有值。表示成员的入群方式: 0 - 由成员邀请入群（包括直接邀请入群和通过邀请链接入群）, 3 - 通过扫描群二维码入群
	QuitScene      int    `xml:"QuitScene"`      // 当是成员退群时有值。表示成员的退群方式: 0 - 自己退群, 1 - 群主/群管理员移出
	MemChangeCnt   int    `xml:"MemChangeCnt"`   // 当是成员入群或退群时有值。表示成员变更数量
	TagType        string `xml:"TagType"`        // 创建标签时，此项为tag，创建标签组时，此项为tag_group
	StrategyID     string `xml:"StrategyId"`     // 规则组id，如果修改了规则组标签的顺序，则回调时会带上此标签所属规则组的id

	// 应用消息相关

}

//EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string   `xml:"Encrypt"    json:"Encrypt"`
}

//ResponseEncryptedXMLMsg 需要返回的消息体
type ResponseEncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	EncryptedMsg string   `xml:"Encrypt"      json:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature" json:"MsgSignature"`
	Timestamp    int64    `xml:"TimeStamp"    json:"TimeStamp"`
	Nonce        string   `xml:"Nonce"        json:"Nonce"`
}

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
}

//SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName CDATA) {
	msg.ToUserName = toUserName
}

//SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName CDATA) {
	msg.FromUserName = fromUserName
}

//SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

//SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}

//GetOpenID get the FromUserName value
func (msg *CommonToken) GetOpenID() string {
	return string(msg.FromUserName)
}
