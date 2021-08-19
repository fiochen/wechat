package contacts

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	listSimpleUserURL  = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist"
	listUserURL        = "https://qyapi.weixin.qq.com/cgi-bin/user/list"
	getUserURL         = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	createUserURL      = "https://qyapi.weixin.qq.com/cgi-bin/user/create"
	updateUserURL      = "https://qyapi.weixin.qq.com/cgi-bin/user/update"
	deleteUserURL      = "https://qyapi.weixin.qq.com/cgi-bin/user/delete"
	batchDeleteUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete"
	convertOpenIDURL   = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid"
	convertUserIDURL   = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid"
	authSuccURL        = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc"
	batchInviteURL     = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite"
	getJoinQRCode      = "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode"
	getActiveStat      = "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat"
)

//是否单独获取该部门成员
const (
	FetchChildSingly      = 0
	FetchChildRecursively = 1
)

//UserSimple UserSimple
type UserSimple struct {
	UserID     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department,omitempty"`
	OpenUserID string `json:"open_userid,omitempty"`
}

//UserBasic UserBasic
type UserBasic struct {
	UserID           string           `json:"userid"`
	Name             string           `json:"name"`
	Department       []int            `json:"department,omitempty"`
	Alias            string           `json:"alias,omitempty"`
	Mobile           string           `json:"mobile,omitempty"`
	Order            []int            `json:"order,omitempty"`
	Position         string           `json:"position,omitempty"`
	Gender           string           `json:"gender,omitempty"`
	Email            string           `json:"email,omitempty"`
	Telephone        string           `json:"telephone,omitempty"`
	IsLeaderInDept   []int            `json:"is_leader_in_dept,omitempty"`
	ExternalPosition string           `json:"external_position,omitempty"`
	Nickname         string           `json:"nickname,omitempty"`
	Address          string           `json:"address,omitempty"`
	MainDepartment   int              `json:"main_department,omitempty"`
	QRCode           string           `json:"qr_code,omitempty"`
	ExtAttr          *ExternalAttr    `json:"extattr,omitempty"`
	ExternalProfile  *ExternalProfile `json:"external_profile,omitempty"`
}

//UserExt UserExt
type UserExt struct {
	OpenUserID  string `json:"open_userid,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	ThumbAvatar string `json:"thumb_avatar,omitempty"`
	Status      int    `json:"status,omitempty"`
	HideMobile  int    `json:"hide_mobile,omitempty"`
	EnglishName string `json:"english_name,omitempty"`
}

//UserCreate UserCreate
type UserCreate struct {
	AvatarMediaID string `json:"avatar_mediaid,omitempty"`
	Enable        *int   `json:"enable,omitempty"`
	ToInvite      *bool  `json:"to_invite,omitempty"`
}

//UserUpdate UserUpdate
type UserUpdate struct {
	AvatarMediaID string `json:"avatar_mediaid,omitempty"`
	Enable        *int   `json:"enable,omitempty"`
	NewUserID     string `json:"new_userid,omitempty"`
}

//UserInfo User Info
type UserInfo struct {
	util.CommonError
	UserBasic
	UserExt
}

//UserCreateInfo User Create Info
type UserCreateInfo struct {
	UserBasic
	UserCreate
}

//UserUpdateInfo UserUpdateInfo
type UserUpdateInfo struct {
	UserBasic
	UserUpdate
}

//resListUserSimple 返回数据
type resListUserSimple struct {
	util.CommonError
	UserList []UserSimple `json:"userlist,omitempty"`
}

//resListUser 返回数据
type resListUser struct {
	util.CommonError
	UserList []UserInfo `json:"userlist,omitempty"`
}

type userIDList struct {
	UserIDList []string `json:"useridlist"`
}

//ListSimpleUsers 获取部门成员
func (contacts *Contacts) ListSimpleUsers(departmentID int, fetchChildOptions ...int) (userList []UserSimple, err error) {
	var fetchChild int
	if len(fetchChildOptions) >= 1 {
		fetchChild = fetchChildOptions[0]
	}

	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&department_id=%d&fetch_child=%d", listSimpleUserURL, accessToken, departmentID, fetchChild)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resListUserSimple
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("ListSimpleUsers Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	userList = res.UserList
	return
}

//ListUsers 获取部门成员详情
func (contacts *Contacts) ListUsers(departmentID int, fetchChildOptions ...int) (userList []UserInfo, err error) {
	var fetchChild int
	if len(fetchChildOptions) >= 1 {
		fetchChild = fetchChildOptions[0]
	}

	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&department_id=%d&fetch_child=%d", listUserURL, accessToken, departmentID, fetchChild)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resListUser
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("ListUsers Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	userList = res.UserList
	return
}

//GetUser 读取成员
func (contacts *Contacts) GetUser(userID string) (res UserInfo, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&userid=%s", getUserURL, accessToken, userID)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("ListUsers Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//CreateUser 创建成员
func (contacts *Contacts) CreateUser(user UserCreateInfo) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", createUserURL, accessToken)

	response, err := util.PostJSON(uri, &user)
	if err != nil {
		return
	}
	var res util.CommonError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateUser Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//UpdateUser 更新成员
func (contacts *Contacts) UpdateUser(user UserUpdateInfo) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", updateUserURL, accessToken)

	response, err := util.PostJSON(uri, &user)
	if err != nil {
		return
	}
	var res util.CommonError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("UpdateUser Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//DeleteUser 删除成员
func (contacts *Contacts) DeleteUser(userID string) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&userid=%s", deleteUserURL, accessToken, userID)

	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res util.CommonError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DeleteUser Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//BatchDeleteUser 批量删除成员
func (contacts *Contacts) BatchDeleteUser(userIDs []string) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", batchDeleteUserURL, accessToken)

	response, err := util.PostJSON(uri, &userIDList{UserIDList: userIDs})
	if err != nil {
		return
	}
	var res util.CommonError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateUser Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//resUser 返回数据
type resUser struct {
	util.CommonError
	OpenID    string `json:"openid,omitempty"`
	UserID    string `json:"userid,omitempty"`
	ActiveCnt int    `json:"active_cnt,omitempty"`
}

//ConvertOpenID userid转openid
func (contacts *Contacts) ConvertOpenID(userID string) (openID string, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", convertOpenIDURL, accessToken)

	response, err := util.PostJSON(uri, map[string]interface{}{"userid": userID})
	if err != nil {
		return
	}
	var res resUser
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("ConvertOpenID Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	openID = res.OpenID
	return
}

//ConvertUserID openid转userid
func (contacts *Contacts) ConvertUserID(openID string) (userID string, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", convertUserIDURL, accessToken)

	response, err := util.PostJSON(uri, map[string]interface{}{"openid": openID})
	if err != nil {
		return
	}
	var res resUser
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("ConvertUserID Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	userID = res.UserID
	return
}

//todo: 二次验证
// https://work.weixin.qq.com/api/doc/90000/90135/90203

//todo: 邀请成员
// https://work.weixin.qq.com/api/doc/90000/90135/90975

//todo: 获取加入企业二维码
// https://work.weixin.qq.com/api/doc/90000/90135/91714

//GetActiveStat 获取企业活跃成员数
func (contacts *Contacts) GetActiveStat(date string) (activeCnt int, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getActiveStat, accessToken)

	response, err := util.PostJSON(uri, map[string]interface{}{"date": date})
	if err != nil {
		return
	}
	var res resUser
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("GetActiveStat Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	activeCnt = res.ActiveCnt
	return
}
