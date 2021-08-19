package contacts

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	listTagURL     = "https://qyapi.weixin.qq.com/cgi-bin/tag/list"
	createTagURL   = "https://qyapi.weixin.qq.com/cgi-bin/tag/create"
	updateTagURL   = "https://qyapi.weixin.qq.com/cgi-bin/tag/update"
	deleteTagURL   = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete"
	getTagURL      = "https://qyapi.weixin.qq.com/cgi-bin/tag/get"
	addTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers"
	delTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers"
)

//Tag 标签信息
type Tag struct {
	ID   int    `json:"tagid,omitempty"`
	Name string `json:"tagname,omitempty"`
}

//TagUsers TagUsers
type TagUsers struct {
	TagID     int      `json:"tagid"`
	UserList  []string `json:"userlist,omitempty"`
	PartyList []int    `json:"partylist,omitempty"`
}

type resTag struct {
	util.CommonError
	TagID        int    `json:"tagid,omitempty"`
	Tags         []Tag  `json:"taglist,omitempty"`
	InvalidList  string `json:"invalidlist,omitempty"`
	InvalidParty []int  `json:"invalidparty,omitempty"`
}

//ResTagUsers ResTagUsers
type ResTagUsers struct {
	util.CommonError
	TagName   string       `json:"tagname,omitempty"`
	UserList  []UserSimple `json:"userlist,omitempty"`
	PartyList []int        `json:"partylist,omitempty"`
}

//GetTags 获取标签列表
func (contacts *Contacts) GetTags() (tagList []Tag, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", listTagURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}

	var res resTag
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("GetTags Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}

	tagList = res.Tags
	return
}

//CreateTag 创建标签
func (contacts *Contacts) CreateTag(tag Tag) (tagID int, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", createTagURL, accessToken)
	response, err := util.PostJSON(uri, &tag)
	if err != nil {
		return
	}
	var res resTag
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	tagID = res.TagID
	return
}

//UpdateTag 更新标签
func (contacts *Contacts) UpdateTag(tag Tag) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", updateTagURL, accessToken)
	response, err := util.PostJSON(uri, &tag)
	if err != nil {
		return
	}
	var res resTag
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("UpdateTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//DeleteTag 删除标签
func (contacts *Contacts) DeleteTag(tagID int) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&tagid=%d", deleteTagURL, accessToken, tagID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resTag
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DeleteTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//AddTagUsers 增加标签成员
func (contacts *Contacts) AddTagUsers(tagUsers TagUsers) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", addTagUsersURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &tagUsers)
	if err != nil {
		return
	}
	var res resTag
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("AddTagUsers Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	if res.InvalidList != "" || res.InvalidParty != nil {
		err = fmt.Errorf("AddTagUsers Error, invalidlist=%s, invalidparty=%v", res.InvalidList, res.InvalidParty)
		return
	}
	return
}

//GetTagUsers 获取标签成员
func (contacts *Contacts) GetTagUsers(tagID int) (res ResTagUsers, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&tagid=%d", getTagURL, accessToken, tagID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DeleteTag Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//DelTagUsers 删除标签成员
func (contacts *Contacts) DelTagUsers(tagUsers TagUsers) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", delTagUsersURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &tagUsers)
	if err != nil {
		return
	}
	var res resTag
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DelTagUsers Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	if res.InvalidList != "" || res.InvalidParty != nil {
		err = fmt.Errorf("DelTagUsers Error, invalidlist=%s, invalidparty=%v", res.InvalidList, res.InvalidParty)
		return
	}
	return
}
