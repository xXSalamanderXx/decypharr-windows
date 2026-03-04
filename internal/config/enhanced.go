package config

import (
"os"
"strings"
)

const envWindowsEnhanced = "DECYPHARR_WINDOWS_ENHANCED"

// WindowsEnhancedEnabled returns true if either the persistent config flag is set
// or the DECYPHARR_WINDOWS_ENHANCED environment variable is set to a truthy value.
func WindowsEnhancedEnabled(cfg *Config) bool {
if cfg != nil && cfg.WindowsEnhancedMode {
return true
}
v := os.Getenv(envWindowsEnhanced)
if v == "" {
return false
}
v = strings.ToLower(strings.TrimSpace(v))
return v == "1" || v == "true" || v == "yes"
}
