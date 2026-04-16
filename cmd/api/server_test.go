package api

import (
	"os"
	"testing"

	log "github.com/go-admin-team/go-admin-core/logger"
	"gorm.io/gorm"
)

// newTestLogger creates a *log.Helper suitable for unit tests.
func newTestLogger() *log.Helper {
	return log.NewHelper(log.NewLogger())
}

// TestImportMenuFromConfig_FileNotFound verifies that importMenuFromConfig
// returns normally (no panic) when the menu config file does not exist.
// Validates: Requirements 4.1
func TestImportMenuFromConfig_FileNotFound(t *testing.T) {
	// Use a temp directory as working directory so the hardcoded file is absent.
	tmpDir, err := os.MkdirTemp("", "menu-test-notfound-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to chdir to temp dir: %v", err)
	}
	defer os.Chdir(origDir)

	lg := newTestLogger()
	db := &gorm.DB{} // non-nil so the nil-check is skipped

	// Should not panic; the function logs a warning and returns.
	importMenuFromConfig(db, lg)
}

// TestImportMenuFromConfig_InvalidJSON verifies that importMenuFromConfig
// handles invalid JSON content gracefully (logs error, returns without panic).
// Validates: Requirements 4.2
func TestImportMenuFromConfig_InvalidJSON(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "menu-test-badjson-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Write invalid JSON to the hardcoded filename in the temp directory.
	badJSON := []byte(`{this is not valid json!!!}`)
	if err := os.WriteFile(tmpDir+"/menu.json", badJSON, 0644); err != nil {
		t.Fatalf("failed to write invalid JSON file: %v", err)
	}

	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to chdir to temp dir: %v", err)
	}
	defer os.Chdir(origDir)

	lg := newTestLogger()
	db := &gorm.DB{} // non-nil so the nil-check is skipped

	// Should not panic; the function logs an error about JSON parsing and returns.
	importMenuFromConfig(db, lg)
}

// TestImportMenuFromConfig_NilDB verifies that importMenuFromConfig
// handles a nil database connection gracefully (logs warning, returns without panic).
// Validates: Requirements 4.4
func TestImportMenuFromConfig_NilDB(t *testing.T) {
	lg := newTestLogger()

	// Should not panic; the function logs a warning about nil db and returns.
	importMenuFromConfig(nil, lg)
}
