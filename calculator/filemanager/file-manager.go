package filemanager

import (
	"fmt"
	"os"
)

type OsFileManager struct {
	version string
}

func New(version string) *OsFileManager {
	return &OsFileManager{
		version: version,
	}
}

func (osfm *OsFileManager) WriteToFile(filename string, data []byte, perm os.FileMode) error {
	fmt.Println("Writing object to file...")
	err := os.WriteFile(filename, data, perm)
	if err != nil {
		return fmt.Errorf("Error writing file to disk: %w", err)
	}
	return nil
}

func (osfm *OsFileManager) ReadFile(filename string) ([]byte, error) {
	fmt.Println()
	data, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading from file: %w", err)
	}
	fmt.Printf("what is the data %s\n", data)
	return data, nil
}
