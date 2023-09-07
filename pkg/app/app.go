// Package app 应用信息
package app

import "gohub/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}
