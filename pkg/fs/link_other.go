#//go:build !windows
package fs

// CreateLinkWithRetries is a no-op on non-Windows platforms (shim).
func CreateLinkWithRetries(target, linkPath string, attempts int) error { return nil }
