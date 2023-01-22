package fileutil

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/f1zm0/utl/randutil"
)

// fsExists checks if a directory or a file exists at the provided path;
// returns a boolean value that indicates if the file or the directory exist and an error.
func fsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// SaveFileToDisk saves the file to disk, and creates the directory if it doesn't exist.
// The file is saved at the provided path with 0600 permission bits.
// It returns an error if any error occurs.
func SaveFileToDisk(outDirPath, outFileName string, fileContent []byte) error {
	dirExists, err := fsExists(outDirPath)
	if err != nil {
		return err
	}

	if !dirExists {
		err := os.MkdirAll(outDirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Write content to file
	outFilePath := path.Join(outDirPath, outFileName)
	err = os.WriteFile(outFilePath, fileContent, 0o600)
	if err != nil {
		return err
	}

	return nil
}

// ReadUniqueFileLines reads the content of the provided file,
// and returns a slice of strings that contains the unique lines,
// or an error if any error occurs.
func ReadUniqueFileLines(fname string) ([]string, error) {
	fReader, err := os.Open(path.Clean(fname))
	if err != nil {
		return []string{}, err
	}

	defer func() {
		err = fReader.Close()
	}()

	// Read from file
	linesMap := make(map[string]bool)
	sc := bufio.NewScanner(fReader)
	for sc.Scan() {
		// Trim whitespaces before and after line
		line := strings.Trim(sc.Text(), " ")

		// Skip blank lines
		if line == "" {
			continue
		}

		// Add line to map
		linesMap[line] = true
	}

	// Make slice from map
	lines := make([]string, len(linesMap))
	i := 0
	for k := range linesMap {
		lines[i] = k
		i++
	}

	return lines, nil
}

// GetFileNameFromURL returns a string generated from the provided URL, that can be used as a file name.
// returns a string that contains the file name and an error if any error occurs.
func GetFileNameFromURL(resourceURL string) string {
	var normBaseURL string

	bu, err := url.Parse(resourceURL)
	if err != nil {
		normBaseURL = ""
	} else {
		if bu.Path != "" && bu.Path != "/" {
			normBaseURL = fmt.Sprintf("%s-%s-%s", bu.Scheme, bu.Host, strings.Replace(bu.Path, "/", "-", -1))
		} else {
			normBaseURL = bu.Host
		}
	}

	// Check if inline script
	if !strings.HasSuffix(strings.ToLower(resourceURL), ".js") {
		return fmt.Sprintf("inline-%s-%s.js", normBaseURL, randutil.GetRandomStringN(16))
	}

	return fmt.Sprintf("%s-%s.js", normBaseURL, randutil.GetRandomStringN(16))
}
