package testcontainers

import (
	"github.com/testcontainers/testcontainers-go/internal/config"
)

// TestcontainersConfig represents the configuration for Testcontainers
type TestcontainersConfig struct {
	Host           string `properties:"docker.host,default="`                    // Deprecated: use Config.Host instead
	TLSVerify      int    `properties:"docker.tls.verify,default=0"`             // Deprecated: use Config.TLSVerify instead
	CertPath       string `properties:"docker.cert.path,default="`               // Deprecated: use Config.CertPath instead
	RyukDisabled   bool   `properties:"ryuk.disabled,default=false"`             // Deprecated: use Config.RyukDisabled instead
	RyukPrivileged bool   `properties:"ryuk.container.privileged,default=false"` // Deprecated: use Config.RyukPrivileged instead
	Config         config.Config
}

// ReadConfig reads from testcontainers properties file, storing the result in a singleton instance
// of the TestcontainersConfig struct
func ReadConfig() TestcontainersConfig {
	cfg := config.Read()

	if cfg.RyukDisabled {
		ryukDisabledMessage := `
**********************************************************************************************
Ryuk has been disabled for the current execution. This can cause unexpected behavior in your environment.
More on this: https://golang.testcontainers.org/features/garbage_collector/
**********************************************************************************************`

		Logger.Printf(ryukDisabledMessage)
	}

	return TestcontainersConfig{
		Host:           cfg.Host,
		TLSVerify:      cfg.TLSVerify,
		CertPath:       cfg.CertPath,
		RyukDisabled:   cfg.RyukDisabled,
		RyukPrivileged: cfg.RyukPrivileged,
		Config:         cfg,
	}
}
