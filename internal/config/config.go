package config

import (
	"os"
)

type cfg struct {
	BindAddr         string
	VmwareUser       string
	VmwarePass       string
	LogLevel         string
	HTTPWriteTimeout string
	HTTPReadTimeout  string
}

func GetConfig() *cfg {
	return &cfg{
		BindAddr:         os.Getenv("VMWARE_EXPORTER_BIND_ADDR"),
		VmwareUser:       os.Getenv("VMWARE_EXPORTER_VMWARE_USER"),
		VmwarePass:       os.Getenv("VMWARE_EXPORTER_VMWARE_PASSWORD"),
		LogLevel:         os.Getenv("VMWARE_EXPORTER_LOG_LEVEL"),
		HTTPWriteTimeout: os.Getenv("VMWARE_EXPORTER_HTTP_WRITE_TIMEOUT"),
		HTTPReadTimeout:  os.Getenv("VMWARE_EXPORTER_HTTP_READ_TIMEOUT"),
	}
}
