package external

import "github.com/silenceper/wechat/v2/work/context"

//External 外部联系人管理
type External struct {
	*context.Context
}

//NewExternal new init external contacts
func NewExternal(ctx *context.Context) *External {
	return &External{
		ctx,
	}
}
