package index

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"git-gud/internal/repository"
	"io"
	"os"
	"path/filepath"
)

type IndexEntry struct {
	Hash [32]byte
	Path string
}

var INDEXFILE_PATH = filepath.Join(".gg", "index")

func IndexFile(indexEntry IndexEntry) error {
	isIndexed, fileAtEntry, err := scanForEntry(indexEntry)
	defer fileAtEntry.Close()
	if err != nil {
		return err
	}

	if !isIndexed {
		return appendEntry(indexEntry)
	} else {
		return updateHashIfChanged(fileAtEntry, indexEntry.Hash)
	}
}

func scanForEntry(indexEntry IndexEntry) (isIndexed bool, hashStart *os.File, err error) {
	file, err := repository.OpenFile(INDEXFILE_PATH, os.O_RDWR, 0644)
	if err != nil {
		return false, nil, err
	}

	hashBuf := make([]byte, 32)  // Buffer to hold the 32-byte hash
	lengthBuf := make([]byte, 2) // Buffer to hold the 2-byte path length

	for {
		nextByte := make([]byte, 1)
		if _, err := file.Read(nextByte); err != nil {
			if err == io.EOF {
				return false, nil, nil
			}
			return false, nil, fmt.Errorf("error reading next byte: %w", err)
		}
		file.Seek(-1, io.SeekCurrent)

		// Read the first 34 bytes: 32 for hash + 2 for path length
		if _, err := file.Read(hashBuf); err != nil {
			return false, nil, err
		}
		if _, err := file.Read(lengthBuf); err != nil {
			return false, nil, err
		}

		// Extract the path length
		pathLength := binary.BigEndian.Uint16(lengthBuf)

		// Read the path based on the path length
		pathBytes := make([]byte, pathLength)
		if _, err := file.Read(pathBytes); err != nil {
			return false, nil, err
		}

		// Convert bytes to string
		path := string(pathBytes)

		// Compare the path with the target pathname
		if path == indexEntry.Path {
			entryLength := int64(34 + pathLength)
			file.Seek(-entryLength, io.SeekCurrent)
			return true, file, nil
		}
	}
}

func updateHashIfChanged(hashStart *os.File, newHash [32]byte) error {
	hashBuf := make([]byte, 32)
	_, err := hashStart.Read(hashBuf)
	if err != nil {
		return err
	}
	originalHash := [32]byte{}
	copy(originalHash[:], hashBuf)

	if newHash != originalHash {
		_, err := hashStart.Seek(-32, io.SeekCurrent)
		if err != nil {
			return fmt.Errorf("error seeking to update hash: %w", err)
		}
		_, err = hashStart.Write(newHash[:])
		if err != nil {
			return fmt.Errorf("error writing new hash: %w", err)
		}
	}
	return nil
}

func entryToBytes(indexEntry IndexEntry) ([]byte, error) {
	// Create buffer
	buf := new(bytes.Buffer)

	// Write hash to buffer
	_, err := buf.Write(indexEntry.Hash[:])
	if err != nil {
		return nil, err
	}

	// Write path length - 2 bytes -  to buffer
	pathLength := uint16(len(indexEntry.Path))
	err = binary.Write(buf, binary.BigEndian, pathLength)
	if err != nil {
		return nil, err
	}

	// Write path as bytes to buffer
	_, err = buf.Write([]byte(indexEntry.Path))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func appendEntry(indexEntry IndexEntry) error {
	bytes, err := entryToBytes(indexEntry)
	if err != nil {
		return err
	}
	err = repository.AppendToFile(INDEXFILE_PATH, bytes)
	if err != nil {
		return err
	}
	return nil
}
