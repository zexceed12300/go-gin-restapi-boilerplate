package helpers

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"
)

var mimeTypes = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
	".bmp":  "image/bmp",
	".webp": "image/webp",
}

func GetMimeType(ext string) string {
	mimeType, found := mimeTypes[ext]
	if !found {
		mimeType = "application/octet-stream" // Default MIME type
	}
	return mimeType
}

func FileToBlob(fileHeader *multipart.FileHeader) ([]byte, error) {
	allowedExtensions := map[string]struct{}{
		".jpg":  {},
		".jpeg": {},
		".png":  {},
		".pdf":  {},
	}

	ext := filepath.Ext(fileHeader.Filename)
	if _, allowed := allowedExtensions[ext]; !allowed {
		return nil, errors.New("file extension not allowed")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileByte, nil
}

func BlobToBase64(b []byte, extension string) (string, error) {
	b64 := base64.StdEncoding.EncodeToString(b)

	base64url := map[string]string{
		".jpg":  "data:image/jpg;base64," + b64,
		".jpeg": "data:image/jpeg;base64," + b64,
		".png":  "data:image/png;base64," + b64,
		".pdf":  "data:application/pdf;base64," + b64,
	}

	var res string

	if _, found := base64url[extension]; found {
		res = base64url[extension]
	} else {
		return "", errors.New("file extension not found")
	}

	return res, nil
}
