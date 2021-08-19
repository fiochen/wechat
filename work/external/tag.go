package external

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	getCorpTagListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list"
	addCorpTagURL     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag"
	editCorpTagURL    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag"
	delCorpTagURL     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag"
	markTagURL        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag"
)

//Tag 企业标签信息
type Tag struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CreateTime int    `json:"create_time,omitempty"`
	Order      int    `json:"order,omitempty"`
	AgentID    int    `json:"agentid,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
}

//TagGroup Tag Group
type TagGroup struct {
	GroupID    string `json:"group_id,omitempty"`
	GroupName  string `json:"group_name,omitempty"`
	CreateTime int    `json:"create_time,omitempty"`
	Order      int    `json:"order,omitempty"`
	Tags       []Tag  `json:"tag,omitempty"`
	AgentID    int    `json:"agentid,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
}

//AddTagRes AddTagRes
type AddTagRes struct {
	util.CommonError
	TagGroup TagGroup `json:"tag_group,omitempty"`
}

//GetTagRes GetTagRes
type GetTagRes struct {
	util.CommonError
	TagGroups []TagGroup `json:"tag_group,omitempty"`
}

//AddCorpTag 添加企业客户标签
func (externalContacts *External) AddCorpTag(data TagGroup) (group TagGroup, err error) {

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", addCorpTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &data)
	if err != nil {
		return
	}
	var res AddTagRes
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("AddCorpTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	group = res.TagGroup
	return
}

//TagRequestData TagRequestData
type TagRequestData struct {
	TagIDList   []string `json:"tag_id,omitempty"`
	GroupIDList []string `json:"group_id,omitempty"`
	AgentID     string   `json:"agentid,omitempty"`
}

//GetCorpTagList 获取企业标签库
func (externalContacts *External) GetCorpTagList(data TagRequestData) (groups []TagGroup, err error) {

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getCorpTagListURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &data)
	if err != nil {
		return
	}
	var res GetTagRes
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("GetCorpTagList Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	groups = res.TagGroups
	return
}

//EditCorpTag 编辑企业客户标签
func (externalContacts *External) EditCorpTag(data Tag) (err error) {

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", editCorpTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &data)
	if err != nil {
		return
	}
	var res util.CommonError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("EditCorpTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//DelCorpTag 删除企业客户标签
func (externalContacts *External) DelCorpTag(data TagRequestData) (err error) {

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", delCorpTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &data)
	if err != nil {
		return
	}
	var res util.CommonError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DelCorpTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//MarkTagRequestData MarkTagRequestData
type MarkTagRequestData struct {
	UserID          string   `json:"userid"`
	ExternalUserID  string   `json:"external_userid"`
	AddTagIDList    []string `json:"add_tag,omitempty"`
	RemoveTagIDList []string `json:"remove_tag,omitempty"`
}

//MarkTag 编辑客户企业标签
func (externalContacts *External) MarkTag(data MarkTagRequestData) (err error) {

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", markTagURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &data)
	if err != nil {
		return
	}
	var res util.CommonError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("MarkTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
