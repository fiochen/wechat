package contacts

import (
	"github.com/silenceper/wechat/v2/work/context"
)

//Contacts Contacts
type Contacts struct {
	*context.Context
}

//NewContacts new init contacts
func NewContacts(ctx *context.Context) *Contacts {
	return &Contacts{
		ctx,
	}
}
