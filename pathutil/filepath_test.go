package pathutil_test

import (
	"testing"

	"github.com/f1zm0/utl/pathutil"
	"github.com/stretchr/testify/assert"
)

func TestParseFilepath(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name     string
		args     args
		expected *pathutil.Filepath
	}{
		{
			name: "TestParseFilepath",
			args: args{
				f: "test.txt",
			},
			expected: &pathutil.Filepath{
				IsDir:     false,
				Dirpath:   ".",
				Filepath:  "test",
				Extension: "txt",
			},
		},
		{
			name: "Test ParseFilepath with no extension",
			args: args{
				f: "test",
			},
			expected: &pathutil.Filepath{
				IsDir:     false,
				Dirpath:   ".",
				Filepath:  "test",
				Extension: "",
			},
		},
		{
			name: "Test ParseFilepath with absolute path",
			args: args{
				f: "/home/user/test.txt",
			},
			expected: &pathutil.Filepath{
				IsDir:     false,
				Dirpath:   "/home/user",
				Filepath:  "test",
				Extension: "txt",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := pathutil.ParseFilepath(tt.args.f)
			assert.Equal(t, tt.expected.IsDir, fp.IsDir)
			assert.Equal(t, tt.expected.Dirpath, fp.Dirpath)
			assert.Equal(t, tt.expected.Filepath, fp.Filepath)
			assert.Equal(t, tt.expected.Extension, fp.Extension)
		})
	}
}
