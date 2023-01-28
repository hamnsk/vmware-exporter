package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"reflect"
)

var _ Client = &vaultClient{}

type vaultClient struct {
	Ttl      int
	Client   *api.Client
	AuthInfo *api.Secret
}

type Client interface {
	GetTokenTTL() int
	GetClient() *api.Client
	GetAuthInfo() *api.Secret
}

func NewClient(cfg interface{}) (Client, error) {
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
		"role_id": roleId,
	})
	if err != nil {
		return nil, err
	}

	client.SetToken(resp.Auth.ClientToken)
	return &vaultClient{
		Ttl:      resp.Auth.LeaseDuration,
		Client:   client,
		AuthInfo: resp,
	}, nil
}

func (c *vaultClient) GetTokenTTL() int {
	return c.Ttl
}

func (c *vaultClient) GetClient() *api.Client {
	return c.Client
}

func (c *vaultClient) GetAuthInfo() *api.Secret {
	return c.AuthInfo
}
