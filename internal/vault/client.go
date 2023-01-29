package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"reflect"
	"vmware-exporter/pkg/logging"
)

var _ Client = &vaultClient{}

type vaultClient struct {
	VaultAddr     string
	VaultAuthName string
	VaultRoleID   string
	client        *api.Client
	authInfo      *api.Secret
	logger        *logging.Logger
}

type Client interface {
	GetClient() *api.Client
	KeepAlive()
}

func NewClient(cfg interface{}, l *logging.Logger) (Client, error) {
	appCfg := reflect.ValueOf(cfg).Elem()

	vaultAddr := appCfg.FieldByName("VaultAddr").Interface().(string)
	authName := appCfg.FieldByName("VaultAuthName").Interface().(string)
	roleId := appCfg.FieldByName("VaultRoleID").Interface().(string)

	client, resp, err := login(vaultAddr, authName, roleId)
	if err != nil {
		return nil, err
	}
	return &vaultClient{
		VaultAddr:     vaultAddr,
		VaultAuthName: authName,
		VaultRoleID:   roleId,
		client:        client,
		authInfo:      resp,
		logger:        l,
	}, nil
}

func login(vaultAddr string, authName string, roleId string) (*api.Client, *api.Secret, error) {
	conf := api.DefaultConfig()
	client, err := api.NewClient(conf)

	if err != nil {
		return nil, nil, err
	}

	client.SetAddress(vaultAddr)

	resp, err := auth(client, authName, roleId)
	if err != nil {
		return nil, nil, err
	}

	client.SetToken(resp.Auth.ClientToken)
	return client, resp, nil
}

func auth(client *api.Client, authName string, roleId string) (*api.Secret, error) {
	resp, err := client.Logical().Write(fmt.Sprintf("auth/%s/login", authName), map[string]interface{}{
		"role_id": roleId,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vaultClient) GetClient() *api.Client {
	return c.client
}

func (c *vaultClient) renewToken() error {
	renew := c.authInfo.Auth.Renewable
	if !renew {
		c.logger.Error("Token is not configured to be renewable. Re-attempting login.")
		return nil
	}
	watcher, err := c.client.NewLifetimeWatcher(&api.LifetimeWatcherInput{
		Secret:    c.authInfo,
		Increment: c.authInfo.Auth.LeaseDuration,
	})

	if err != nil {
		c.logger.Error(fmt.Sprintf("unable to initialize new lifetime watcher for renewing auth token: %w", err))
	}

	go watcher.Start()
	defer watcher.Stop()

	for {
		select {
		case err := <-watcher.DoneCh():
			if err != nil {
				c.logger.Error(fmt.Sprintf("Failed to renew token: %v. Re-attempting login.", err))
			}
			c.logger.Info("Token can no longer be renewed. Re-attempting login.")
			resp, err := auth(c.client, c.VaultAuthName, c.VaultRoleID)
			if err != nil {
				return err
			}
			c.authInfo = resp
			c.client.SetToken(resp.Auth.ClientToken)
			c.logger.Info("Login to Hashicorp Vault Success.")
			return nil

		case renewal := <-watcher.RenewCh():
			c.logger.Info(fmt.Sprintf("Token successfully renewed at %s", renewal.RenewedAt))
		}
	}
}

func (c *vaultClient) KeepAlive() {
	err := c.renewToken()
	if err != nil {
		c.logger.Error(err.Error())
	}
}
