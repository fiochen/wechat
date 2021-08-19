//go:build !linux || !cgo || !msgaudit
// +build !linux !cgo !msgaudit

// Package msgaudit for unsupport platform
package msgaudit

import (
	"fmt"

	"github.com/silenceper/wechat/v2/work/config"
)

// Client 会话存档
type Client struct {
}

// NewClient new
func NewClient(cfg *config.Config) (*Client, error) {
	return nil, fmt.Errorf("会话存档功能目前只支持Linux平台运行，并且打开设置CGO_ENABLED=1")
}

// Free unsupported
func (s *Client) Free() {
	return
}

// GetChatData unsupported
func (s *Client) GetChatData(seq uint64, limit uint64, proxy string, passwd string, timeout int) ([]ChatData, error) {
	return nil, fmt.Errorf("会话存档功能目前只支持Linux平台运行，并且打开设置CGO_ENABLED=1")
}

// GetRawChatData unsupported
func (s *Client) GetRawChatData(seq uint64, limit uint64, proxy string, passwd string, timeout int) (ChatDataResponse, error) {
	return ChatDataResponse{}, fmt.Errorf("会话存档功能目前只支持Linux平台运行，并且打开设置CGO_ENABLED=1")
}

// DecryptData unsupported
func (s *Client) DecryptData(encryptRandomKey string, encryptMsg string) (msg ChatMessage, err error) {
	return msg, fmt.Errorf("会话存档功能目前只支持Linux平台运行，并且打开设置CGO_ENABLED=1")
}

// GetMediaData unsupported
func (s *Client) GetMediaData(indexBuf string, sdkFileID string, proxy string, passwd string, timeout int) (*MediaData, error) {
	return nil, fmt.Errorf("会话存档功能目前只支持Linux平台运行，并且打开设置CGO_ENABLED=1")
}
