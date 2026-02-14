package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// SaveMultipartFile saves a multipart.FileHeader to the given directory and returns the saved path (relative)
func SaveMultipartFile(dir string, file multipart.File, header *multipart.FileHeader) (string, error) {
    if dir == "" {
        dir = "storage/posters"
    }
    if err := os.MkdirAll(dir, 0o755); err != nil {
        return "", err
    }
    ext := filepath.Ext(header.Filename)
    fname := fmt.Sprintf("%s%s", uuid.New().String(), ext)
    dest := filepath.Join(dir, fname)
    out, err := os.Create(dest)
    if err != nil {
        return "", err
    }
    defer out.Close()
    if _, err := io.Copy(out, file); err != nil {
        return "", err
    }
    return "/" + filepath.ToSlash(dest), nil
}
