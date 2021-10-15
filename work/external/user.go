package external

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/contacts"
)

const (
	listUsersURL    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	getUserURL      = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	batchGetUserURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user"
	remarkUserURL   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark"
)

//UsersRes UsersRes
type UsersRes struct {
	util.CommonError
	ExternalUserIDs []string `json:"external_userid,omitempty"`
}

//UserRes UserRes
type UserRes struct {
	util.CommonError
	ExternalContact User         `json:"external_contact,omitempty"`
	FollowUsers     []FollowUser `json:"follow_user,omitempty"`
	NextCursor      string       `json:"next_cursor,omitempty"`
}

//BatchUserRes BatchUserRes
type BatchUserRes struct {
	util.CommonError
	ExternalContactList []struct {
		ExternalContact User       `json:"external_contact,omitempty"`
		FollowInfo      FollowUser `json:"follow_info,omitempty"`
	} ` json:"external_contact_list"`
	NextCursor string `json:"next_cursor,omitempty"`
}

//User User
type User struct {
	ExternalUserID  string                    `json:"external_userid,omitempty"`
	Name            string                    `json:"name,omitempty"`
	Position        string                    `json:"position,omitempty"`
	Avatar          string                    `json:"avatar,omitempty"`
	CorpName        string                    `json:"corp_name,omitempty"`
	CorpFulName     string                    `json:"corp_full_name,omitempty"`
	Type            int                       `json:"type,omitempty"`
	Gender          int                       `json:"gender,omitempty"`
	UnionID         string                    `json:"unionid,omitempty"`
	ExternalProfile *contacts.ExternalProfile `json:"external_profile,omitempty"`
}

//FollowUser Follow User
type FollowUser struct {
	UserID         string          `json:"userid,omitempty"`
	Remark         string          `json:"remark,omitempty"`
	Description    string          `json:"description,omitempty"`
	CreateTime     int             `json:"createtime,omitempty"`
	Tags           []FollowUserTag `json:"tags,omitempty"`
	RemarkCorpName string          `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string        `json:"remark_mobiles,omitempty"`
	OperUserID     string          `json:"oper_userid,omitempty"`
	AddWay         int             `json:"add_way"`
	State          string          `json:"state,omitempty"`
}

//FollowUserTag Follow User Tag
type FollowUserTag struct {
	GroupName       string `json:"group_name,omitempty"`
	ExternalTagName string `json:"tag_name,omitempty"`
	ExternalTagID   string `json:"tag_id,omitempty"`
	Type            int    `json:"type,omitempty"`
}

//ListUsers 获取客户列表
func (externalContacts *External) ListUsers(userID string) (externalUserIDs []string, err error) {
	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&userid=%s", listUsersURL, accessToken, userID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res UsersRes
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("ListUsers Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	externalUserIDs = res.ExternalUserIDs
	return
}

//GetUser 获取客户详情
func (externalContacts *External) GetUser(externalUserID string, cursorOption ...string) (res UserRes, err error) {

	cursor := ""
	if len(cursorOption) >= 1 {
		cursor = cursorOption[0]
	}

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&external_userid=%s&cursor=%s", getUserURL, accessToken, externalUserID, cursor)
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
		err = fmt.Errorf("GetUser Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//BatchGetExUser 请求内容
type BatchGetExUser struct {
	UserIDList []string `json:"userid_list"`
	Cursor     string   `json:"cursor,omitempty"`
	Limit      int      `json:"limit,omitempty"`
}

//BatchGetUsers 批量获取客户详情
func (externalContacts *External) BatchGetUsers(data BatchGetExUser) (res BatchUserRes, err error) {

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", batchGetUserURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, &data)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("BatchGetUsers Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//RemarkUserData RemarkUserData
type RemarkUserData struct {
	UserID           string   `json:"userid"`
	ExternalUserID   string   `json:"external_userid,omitempty"`
	Remark           string   `json:"remark,omitempty"`
	Description      string   `json:"description,omitempty"`
	RemarkCompany    string   `json:"remark_company,omitempty"`
	RemarkMobiles    []string `json:"remark_mobiles,omitempty"`
	RemarkPicMediaID string   `json:"remark_pic_mediaid,omitempty"`
}

//RemarkUser 修改客户备注信息
func (externalContacts *External) RemarkUser(data RemarkUserData) (err error) {

	var accessToken string
	accessToken, err = externalContacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", remarkUserURL, accessToken)
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
		err = fmt.Errorf("RemarkUser Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
