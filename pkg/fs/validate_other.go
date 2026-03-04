#//go:build !windows
package fs

import "path/filepath"

// ValidateAndNormalizePath is a non-windows shim that performs minimal normalization so
// cross-platform builds remain unaffected. It returns a cleaned path.
func ValidateAndNormalizePath(p string) (string, error) {
if p == "" {
return p, nil
}
// Normalize path (convert windows backslashes -> slashes on non-windows)
p = filepath.Clean(p)
return p, nil
}
