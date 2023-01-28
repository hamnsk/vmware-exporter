package config

import (
	"os"
	"strconv"
)

type cfg struct {
	BindAddr             string
	VmwareUser           string
	VmwarePass           string
	ScrapeTimeout        string
	LogLevel             string
	HTTPWriteTimeout     string
	HTTPReadTimeout      string
	UseVault             bool
	VaultRoleID          string
	VaultAddr            string
	VaultAuthName        string
	VaultSecretStoreName string
	VaultSecretStorePath string
}

func GetConfig() *cfg {
	flag, _ := strconv.ParseBool(os.Getenv("VMWARE_EXPORTER_USE_VAULT"))
	return &cfg{
		BindAddr:             os.Getenv("VMWARE_EXPORTER_BIND_ADDR"),
		VmwareUser:           os.Getenv("VMWARE_EXPORTER_VMWARE_USER"),
		VmwarePass:           os.Getenv("VMWARE_EXPORTER_VMWARE_PASSWORD"),
		ScrapeTimeout:        os.Getenv("VMWARE_EXPORTER_SCRAPE_TIMEOUT"),
		LogLevel:             os.Getenv("VMWARE_EXPORTER_LOG_LEVEL"),
		HTTPWriteTimeout:     os.Getenv("VMWARE_EXPORTER_HTTP_WRITE_TIMEOUT"),
		HTTPReadTimeout:      os.Getenv("VMWARE_EXPORTER_HTTP_READ_TIMEOUT"),
		UseVault:             flag,
		VaultAddr:            os.Getenv("VMWARE_EXPORTER_VAULT_ADDR"),
		VaultAuthName:        os.Getenv("VMWARE_EXPORTER_VAULT_AUTH_NAME"),
		VaultRoleID:          os.Getenv("VMWARE_EXPORTER_VAULT_ROLE_ID"),
		VaultSecretStoreName: os.Getenv("VMWARE_EXPORTER_VAULT_SECRET_STORE_NAME"),
		VaultSecretStorePath: os.Getenv("VMWARE_EXPORTER_VAULT_SECRET_STORE_PATH"),
	}
}
