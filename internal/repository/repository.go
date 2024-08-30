package repository

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateFile(path string) {

}

func OpenFile(filePath string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(filePath, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadFile(filePath string) (string, error) {
	file, err := OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteFile(filePath string, content interface{}) error {
	fileRef, err := OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New("error creating files")
	}
	defer fileRef.Close()

	var writeErr error
	switch v := content.(type) {
	case string:
		_, writeErr = fileRef.WriteString(v)
	case []byte:
		_, writeErr = fileRef.Write(v)
	default:
		return errors.New("unsupported content type")
	}
	if writeErr != nil {
		return errors.New("error writing to file")
	}
	return nil
}

func AppendToFile(filePath string, content interface{}) error {
	fileRef, err := OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("error opening file")
	}
	defer fileRef.Close()

	var writeErr error
	switch v := content.(type) {
	case string:
		_, writeErr = fileRef.WriteString(v)
	case []byte:
		_, writeErr = fileRef.Write(v)
	default:
		return errors.New("unsupported content type")
	}
	if writeErr != nil {
		return errors.New("error writing to file")
	}
	return nil
}

func DeleteFile() {

}

func CreateDirectory(directory string) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		fmt.Println("Failed to create directory")
		return
	}
}

func CheckIfFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

func isVisible(name string) bool {
	return !strings.HasPrefix(name, ".")
}

type fileWithPath struct {
	fs.DirEntry
	Path string
}

func ReadDirectoryFiles() []fileWithPath {
	var visibleFiles []fileWithPath

	// creating a slice of visible files
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip hidden directories and their contents
		if d.IsDir() && !isVisible(d.Name()) {
			// current (root) directory not included in skipping
			if path != "." {
				return filepath.SkipDir
			}
		}

		// append visible files (main.go excluded for local version)
		if isVisible(d.Name()) && !d.IsDir() && d.Name() != "main.go" {
			visibleFiles = append(visibleFiles, fileWithPath{d, path})
		}

		return nil
	})
	if err != nil {
		log.Fatalf("impossible to walk directories: %s", err)
	}
	return visibleFiles
}
