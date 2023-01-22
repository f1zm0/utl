package fileutil_test

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/f1zm0/utl/fileutil"
	"github.com/stretchr/testify/assert"
)

func TestSaveFileToDisk(t *testing.T) {
	testCases := map[string]struct {
		data []byte
	}{
		"generic data": {data: []byte("test data")},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			// Create temp file that will be used to write data to and then deleted
			tfName, tfDelete := createTempFile(t)
			defer tfDelete()

			// Write test data to tempfile
			saveFileErr := fileutil.SaveFileToDisk(os.TempDir(), filepath.Base(tfName), tc.data)
			assert.Nil(t, saveFileErr)

			// Read data written to file to compare with input
			// #nosec G304
			written, readFileErr := os.ReadFile(tfName)
			assert.Nil(t, readFileErr)

			assert.Equal(t, tc.data, written)
		})
	}
}

func TestReadUniqueFileLines_Errors(t *testing.T) {
	testCases := map[string]struct {
		fixture  string
		expected []string
	}{
		"non-existing file": {
			fixture:  path.Join("testdata", "not-exists.txt"),
			expected: []string{},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			got, err := fileutil.ReadUniqueFileLines(tc.fixture)
			assert.NotNil(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestReadUniqueFileLines(t *testing.T) {
	testCases := map[string]struct {
		fixture  string
		expected []string
	}{
		"duplicate and blank lines": {
			fixture: path.Join("testdata", "input_file_001.txt"),
			expected: []string{
				"a.b",
				"a.b/path/to/file.ext",
			},
		},
		"leading-trailing whitespaces": {
			fixture: path.Join("testdata", "input_file_002.txt"),
			expected: []string{
				"a.b",
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			got, err := fileutil.ReadUniqueFileLines(tc.fixture)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

// createTempFile is a helper function that creates temporary files
// Returns the tempfile name and a function closure
// that can be used to delete it when tests are completed.
func createTempFile(t *testing.T) (string, func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "tempfile")
	defer func() {
		err = tempFile.Close()
	}()
	if err != nil {
		t.Fatalf("Error while creating tempfile: %s", err)
	}

	return tempFile.Name(), func() {
		err := os.Remove(tempFile.Name())
		if err != nil {
			t.Fatalf("Error while deleting tempfile: %s", err)
		}
	}
}
