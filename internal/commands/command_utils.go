package commands

import (
	"git-gud/internal/index"
	"git-gud/internal/objects"
	"git-gud/internal/repository"
	"git-gud/internal/utils"
)

func handleFileAdd(filePath string) error {
	content, err := repository.ReadFile(filePath)
	if err != nil {
		return err
	}
	hashString := utils.CreateHash(content)
	indexEntry := index.IndexEntry{
		Hash: utils.StringToByte32(hashString),
		Path: filePath,
	}
	err = index.IndexFile(indexEntry)
	if err != nil {
		return err
	}
	err = objects.CreateBlobObject(content, hashString)
	if err != nil {
		return err
	}
	return nil
}
