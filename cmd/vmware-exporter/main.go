package main

import (
	"os"
	"vmware-exporter/internal/app"
)

func main() {
	os.Exit(app.Run())
}
