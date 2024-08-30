package objects

import (
	"fmt"
	"git-gud/internal/repository"
	"git-gud/internal/utils"
	"path/filepath"
)

func CreateBlobObject(content string, hashString string) error {
	blobHeader := fmt.Sprintf("blob %d\x00", len(content))
	blobContent := blobHeader + content
	isFileCreated := repository.CheckIfFileExists(filepath.Join(".gg", "objects", hashString[:2], hashString[2:]))
	if !isFileCreated {
		compressedContent, err := utils.CompressZstd(blobContent)
		if err != nil {
			return err
		}
		repository.CreateDirectory(filepath.Join(".gg", "objects", hashString[:2]))
		err = repository.WriteFile(filepath.Join(".gg", "objects", hashString[:2], hashString[2:]), compressedContent)
		if err != nil {
			return err
		}
	}
	return nil
}
