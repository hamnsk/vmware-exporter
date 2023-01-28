package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"reflect"
)

//var _ Client = &vaultClient{}

type VaultClient struct {
	Ttl    int
	Client *api.Client
}

type Client interface {
	GetTokenTTL() int
	GetClient() *api.Client
}

func NewClient(cfg interface{}) (*VaultClient, error) {
	conf := api.DefaultConfig()
	client, err := api.NewClient(conf)

	if err != nil {
		return nil, err
	}

	appCfg := reflect.ValueOf(cfg).Elem()

	vaultAddr := appCfg.FieldByName("VaultAddr").Interface().(string)
	authName := appCfg.FieldByName("VaultAuthName").Interface().(string)
	roleId := appCfg.FieldByName("VaultRoleID").Interface().(string)

	client.SetAddress(vaultAddr)

	resp, err := client.Logical().Write(fmt.Sprintf("auth/%s/login", authName), map[string]interface{}{
		//resp, err := client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id": roleId,
	})
	if err != nil {
		return nil, err
	}

	client.SetToken(resp.Auth.ClientToken)
	return &VaultClient{
		Ttl:    resp.Auth.LeaseDuration,
		Client: client,
	}, nil
}

//func (c *vaultClient) GetTokenTTL() int {
//	return c.Ttl
//}
//
//func (c *vaultClient) GetClient() *api.Client {
//	return c.Client
//}
