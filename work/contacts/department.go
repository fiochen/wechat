package contacts

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	listDepartmentURL   = "https://qyapi.weixin.qq.com/cgi-bin/department/list"
	createDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/create"
	updateDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/update"
	deleteDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/delete"
)

//Department 部门信息
type Department struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	NameEn   string `json:"name_en,omitempty"`
	ParentID int    `json:"parentid,omitempty"`
	Order    int    `json:"order,omitempty"`
}

//resDepartment 返回数据
type resDepartment struct {
	util.CommonError
	ID         int          `json:"id,omitempty"`
	Department []Department `json:"department,omitempty"`
}

//GetDepartments 获取部门列表
func (contacts *Contacts) GetDepartments() (departmentInfoList []Department, err error) {
	return contacts.GetDepartmentsByID(0)
}

//GetDepartmentsByID 获取指定部门列表
func (contacts *Contacts) GetDepartmentsByID(id int) (departmentInfoList []Department, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&id=%d", listDepartmentURL, accessToken, id)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}

	var res resDepartment
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("GetDepartments Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}

	departmentInfoList = res.Department
	return
}

//CreateDepartment 创建部门
func (contacts *Contacts) CreateDepartment(departmentInfo Department) (id int, err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", createDepartmentURL, accessToken)
	response, err := util.PostJSON(uri, &departmentInfo)
	if err != nil {
		return
	}
	var res resDepartment
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("CreateDepartment Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	id = res.ID
	return
}

//UpdateDepartment 更新部门
func (contacts *Contacts) UpdateDepartment(departmentInfo Department) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", updateDepartmentURL, accessToken)
	response, err := util.PostJSON(uri, &departmentInfo)
	if err != nil {
		return
	}
	var res resDepartment
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("UpdateDepartment Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

//DeleteDepartment 删除部门
func (contacts *Contacts) DeleteDepartment(id int) (err error) {
	var accessToken string
	accessToken, err = contacts.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&id=%d", deleteDepartmentURL, accessToken, id)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resDepartment
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DeleteDepartment Error, errcode=%d, errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
