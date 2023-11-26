package util

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

const maxFileSizeMB = 10

func Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	cld.Config.URL.Secure = true
	return cld
}

func UploadToCloudinary(fileHeader *multipart.FileHeader) (imageUrl string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileSizeMB := float64(fileHeader.Size) / (1024 * 1024)
	if fileSizeMB > maxFileSizeMB {
		return "", fmt.Errorf("file size exceeds the maximum allowed size of %d MB", maxFileSizeMB)
	}

	uid := uuid.New()

	cld := Credentials()

	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       "Reproduction-Health/" + uid.String(),
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return
	}

	imageUrl = resp.SecureURL

	return imageUrl, nil
}
