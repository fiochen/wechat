package media

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
)

//Material 素材管理
type Material struct {
	*context.Context
}

//NewMaterial init
func NewMaterial(context *context.Context) *Material {
	material := new(Material)
	material.Context = context
	return material
}

//Type 媒体文件类型
type Type string

const (
	//TypeImage 媒体文件:图片
	TypeImage Type = "image"
	//TypeVoice 媒体文件:声音
	TypeVoice Type = "voice"
	//TypeVideo 媒体文件:视频
	TypeVideo Type = "video"
	//TypeFile 媒体文件:普通文件
	TypeFile Type = "file"
)

const (
	mediaUploadURL      = "https://qyapi.weixin.qq.com/cgi-bin/media/upload"
	mediaGetURL         = "https://qyapi.weixin.qq.com/cgi-bin/media/get"
	mediaUploadImageURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg"
	mediaGetVoiceURL    = "https://qyapi.weixin.qq.com/cgi-bin/media/get/jssdk"
)

//Media 临时素材上传返回信息
type Media struct {
	util.CommonError

	Type      Type   `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

//MediaUpload 上传临时素材
func (material *Material) MediaUpload(mediaType Type, filename string) (media Media, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=%s", mediaUploadURL, accessToken, mediaType)
	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &media)
	if err != nil {
		return
	}
	if media.ErrCode != 0 {
		err = fmt.Errorf("MediaUpload error : errcode=%v , errmsg=%v", media.ErrCode, media.ErrMsg)
		return
	}
	return
}

//GetMediaURL 返回临时素材的下载地址供用户自己处理
//NOTICE: URL 不可公开，因为含access_token 需要立即另存文件
func (material *Material) GetMediaURL(mediaID string) (mediaURL string, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}
	mediaURL = fmt.Sprintf("%s?access_token=%s&media_id=%s", mediaGetURL, accessToken, mediaID)
	return
}

//resMediaImage 图片上传返回结果
type resMediaImage struct {
	util.CommonError

	URL string `json:"url"`
}

//ImageUpload 图片上传
func (material *Material) ImageUpload(filename string) (url string, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", mediaUploadImageURL, accessToken)
	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	var image resMediaImage
	err = json.Unmarshal(response, &image)
	if err != nil {
		return
	}
	if image.ErrCode != 0 {
		err = fmt.Errorf("UploadImage error : errcode=%v , errmsg=%v", image.ErrCode, image.ErrMsg)
		return
	}
	url = image.URL
	return
}

//GetVoiceURL 获取高清语音素材
//NOTICE: URL 不可公开，因为含access_token 需要立即另存文件
func (material *Material) GetVoiceURL(mediaID string) (mediaURL string, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}
	mediaURL = fmt.Sprintf("%s?access_token=%s&media_id=%s", mediaGetVoiceURL, accessToken, mediaID)
	return
}
