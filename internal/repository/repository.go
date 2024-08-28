package repository

import (
	"errors"
	"fmt"
	"os"
)

func CreateFile() {

}

func OpenFile(filePath string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(filePath, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadFile() {

}

func WriteFile(filePath string, content string) error {
	fileRef, err := OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New("error creating files")
	}
	defer fileRef.Close()

	_, err = fileRef.WriteString(content)
	if err != nil {
		return errors.New("error writing to file")
	}
	return nil
}

func AppendToFile() {

}

func DeleteFile() {

}

func CreateDirectory(directory string) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		fmt.Println("Failed to create directory")
		return
	}
}
