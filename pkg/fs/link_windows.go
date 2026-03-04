//go:build windows

package fs

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// CreateLinkWithRetries creates a link (symlink for files, junction for directories),
// verifies accessibility of the created link, and retries with backoff on transient failures.
func CreateLinkWithRetries(target, linkPath string, attempts int) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		if err := createLinkOnce(target, linkPath); err != nil {
			lastErr = err
		} else {
			// verify readability
			if _, err := os.Stat(linkPath); err == nil {
				return nil
			} else {
				lastErr = err
			}
		}
		// backoff
		time.Sleep(time.Duration(200*(i+1)) * time.Millisecond)
	}
	if lastErr == nil {
		lastErr = errors.New("link creation failed")
	}
	return lastErr
}

func createLinkOnce(target, linkPath string) error {
	// Ensure parent dir exists
	if err := os.MkdirAll(filepath.Dir(linkPath), 0o755); err != nil {
		return err
	}

	fi, err := os.Stat(target)
	if err != nil {
		return err
	}

	// Try symlink first (works for files and may work for dirs when allowed)
	if err := os.Symlink(target, linkPath); err == nil {
		return nil
	}

	// If target is a directory, attempt a junction via mklink /J (no elevation required)
	if fi.IsDir() {
		// mklink /J <link> <target>
		cmd := exec.Command("cmd", "/C", "mklink", "/J", linkPath, target)
		if out, err := cmd.CombinedOutput(); err == nil {
			_ = out
			return nil
		} else {
			return err
		}
	}

	// For files, no safe fallback is implemented; return error to allow retry/backoff.
	return errors.New("failed to create symlink for file target")
}
