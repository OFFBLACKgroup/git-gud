package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/DataDog/zstd"
)

func CreateHash(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func CompressZstd(content string) ([]byte, error) {

	compressed, err := zstd.Compress(nil, []byte(content))
	if err != nil {
		return nil, err
	}
	return compressed, nil
}

func DecompressZstd(content string) ([]byte, error) {

	decompressed, err := zstd.Decompress(nil, []byte(content))
	if err != nil {
		fmt.Println("Error decompressing content:", err)
		return nil, err
	}
	return decompressed, nil
}

func StringToByte32(content string) [32]byte {
	byte32 := []byte(content)
	return [32]byte(byte32)
}
