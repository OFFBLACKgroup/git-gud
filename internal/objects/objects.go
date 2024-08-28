package objects

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"git-gud/internal/repository"
	"path/filepath"

	"github.com/DataDog/zstd"
)

func createHash(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func compressZstd(content string) ([]byte, error) {

	compressed, err := zstd.Compress(nil, []byte(content))
	if err != nil {
		return nil, err
	}

	fmt.Println("Compressed content:", compressed)
	return compressed, nil
}

func decompressZstd(content string) {

	decompressed, err := zstd.Decompress(nil, []byte(content))
	if err != nil {
		fmt.Println("Error decompressing content:", err)
		return
	}
	fmt.Println("Decompressed content:", string(decompressed))
}

func CreateBlobObject(content string) {
	blobHeader := fmt.Sprintf("blob %d\x00", len(content))
	blobContent := blobHeader + content
	hashString := createHash(blobContent)
	compressedContent, err := compressZstd(blobContent)
	if err != nil {
		fmt.Println("Error compressing content:", err)
		return
	}
	repository.CreateDirectory(filepath.Join(".gg", "objects", hashString[:2]))
	repository.WriteFile(filepath.Join(".gg", "objects", hashString[:2], hashString[2:]), compressedContent)
}
