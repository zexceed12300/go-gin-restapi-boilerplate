package handlers

import (
	"go-gin-restapi-boilerplate/errorhandler"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(file *multipart.FileHeader) (filename string, err error) {
	blob, err := helpers.FileToBlob(file)

	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()

	ext := filepath.Ext(file.Filename)

	filename = uuid + ext

	model := models.File{
		Name: filename,
		Body: blob,
	}

	if err := initializers.DB.Create(&model).Error; err != nil {
		return "", err
	}

	return filename, nil
}

func UploadFiles(files []*multipart.FileHeader) ([]string, error) {
	filenames := make([]string, 0, len(files))

	filesDB := []models.File{}

	for _, f := range files {
		blob, err := helpers.FileToBlob(f)

		if err != nil {
			return nil, err
		}

		uuid := uuid.NewString()

		ext := filepath.Ext(f.Filename)

		filename := uuid + ext

		filenames = append(filenames, filename)

		filesDB = append(filesDB, models.File{Name: filename, Body: blob})
	}

	if err := initializers.DB.Create(filesDB).Error; err != nil {
		return nil, err
	}

	return filenames, nil
}

func GetFile(filename string) (base64 string, err error) {
	file := models.File{}

	if err := initializers.DB.First(&file, models.File{Name: filename}).Error; err != nil {
		return "", err
	}

	ext := filepath.Ext(filename)

	b64, err := helpers.BlobToBase64(file.Body, ext)

	if err != nil {
		return "", err
	}

	return b64, nil
}

func UpdateFile(filename string, newFile *multipart.FileHeader) error {
	oldfile := models.File{}

	if err := initializers.DB.First(&oldfile, models.File{Name: filename}).Error; err != nil {
		return err
	}

	blob, err := helpers.FileToBlob(newFile)

	if err != nil {
		return err
	}

	if err := initializers.DB.Model(&oldfile).Updates(models.File{Body: blob}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteFile(filename string) error {
	file := models.File{}

	if err := initializers.DB.First(&file, models.File{Name: filename}).Error; err != nil {
		return err
	}

	if err := initializers.DB.Delete(&file).Error; err != nil {
		return err
	}

	return nil
}

func DeleteFiles(filenames []string) error {
	files := []models.File{}

	if err := initializers.DB.Find(&files, "name IN ?", filenames).Error; err != nil {
		return err
	}

	if err := initializers.DB.Delete(&files).Error; err != nil {
		return err
	}

	return nil
}

func FileGet(c *gin.Context) {
	uri := models.FileFilenameParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Error getting file",
		})
		return
	}

	file := models.File{}

	if err := initializers.DB.First(&file, models.File{Name: uri.Filename}).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Error getting file",
		})
	}

	mime := helpers.GetMimeType(filepath.Ext(file.Name))

	c.Header("Content-Type", mime)

	c.Writer.Write(file.Body)
}
