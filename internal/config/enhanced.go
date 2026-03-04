package config

import (
	"os"
	"strings"
)

// WindowsEnhancedEnabled returns true if either the persistent config flag is set
// or the DECYPHARR_WINDOWS_ENHANCED environment variable is set to a truthy value.
func WindowsEnhancedEnabled(cfg *Config) bool {
	// If caller passed nil, fallback to the global config singleton
	if cfg == nil {
		cfg = Get()
	}
	if cfg != nil && cfg.WindowsEnhancedMode {
		return true
	}
	v := os.Getenv("DECYPHARR_WINDOWS_ENHANCED")
	if v == "" {
		return false
	}
	v = strings.ToLower(strings.TrimSpace(v))
	return v == "1" || v == "true" || v == "yes"
}
