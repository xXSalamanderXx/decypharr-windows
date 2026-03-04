#//go:build windows
package fs

import (
"errors"
"os"
"path/filepath"
"strings"
)

// ValidateAndNormalizePath verifies the path exists and is a directory, and normalizes
// separators to Windows-style backslashes.
func ValidateAndNormalizePath(p string) (string, error) {
if p == "" {
return "", errors.New("empty path")
}
// Normalize forward slashes to backslashes and clean
p = filepath.Clean(strings.ReplaceAll(p, "/", `\`))

fi, err := os.Stat(p)
if err != nil {
return "", err
}
if !fi.IsDir() {
return "", errors.New("path is not a directory")
}
return p, nil
}
