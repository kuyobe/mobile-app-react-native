package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes)[:length], nil
}

func DownloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to download file")
	}

	fileData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, fileData, 0644)
}

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}