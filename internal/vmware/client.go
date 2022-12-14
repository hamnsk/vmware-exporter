package vmware

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"net/url"
)

func vmwareUrl(username, password string, u *url.URL) {
	if username != "" {
		var password string
		var ok bool
		if u.User != nil {
			password, ok = u.User.Password()
		}
		if ok {
			u.User = url.UserPassword(username, password)
		} else {
			u.User = url.User(username)
		}
	}
	if password != "" {
		var username string
		if u.User != nil {
			username = u.User.Username()
		}
		u.User = url.UserPassword(username, password)
	}
}

func NewClient(ctx context.Context, host, username, password string) (*govmomi.Client, error) {
	var baseurl = fmt.Sprintf("https://username:password@%s"+vim25.Path, host)
	u, err := soap.ParseURL(baseurl)
	if err != nil {
		return nil, err
	}
	vmwareUrl(username, password, u)
	return govmomi.NewClient(ctx, u, true)
}
