package common

import (
	"context"
	"fmt"
	"garination.com/gateway/internal/platform/wasabi"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"mime/multipart"
)

type WasabiRepo interface {
	UploadImage(ctx context.Context, folderName string, file *multipart.FileHeader) (string, error)
	UploadDocument(ctx context.Context, folderName string, file *multipart.FileHeader) (string, error)
	DeleteFile(ctx context.Context, folderName string, fileName string) error
}

type wasabiRepoImpl struct {
	wasabiClient *wasabi.Client
}

func extensionFromImageMime(mime string) string {
	switch mime {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/bmp":
		return ".bmp"
	case "image/tiff":
		return ".tiff"
	default:
		return ""
	}

	return ""
}

func extensionFromDocumentMime(mime string) string {
	switch mime {
	case "application/pdf":
		return ".pdf"
	case "application/msword":
		return ".doc"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return ".docx"
	case "application/vnd.ms-powerpoint":
		return ".ppt"
	case "application/vnd.openxmlformats-officedocument.presentationml.presentation":
		return ".pptx"
	}
	return ""
}

func (w *wasabiRepoImpl) UploadDocument(ctx context.Context, folder string, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}

	extension := extensionFromDocumentMime(file.Header.Get("Content-Type"))
	if extension == "" {
		return "", fmt.Errorf("invalid file type: %s", file.Header.Get("Content-Type"))
	}

	fileName := fmt.Sprintf("%s/documents/%s%s", folder, uuid.NewString(), extension)

	input := &s3.PutObjectInput{
		Bucket: aws.String(w.wasabiClient.Bucket),
		Key:    &fileName,
		Body:   f,
	}

	_, err = w.wasabiClient.S3.PutObject(input)
	if err != nil {
		return "", err
	}

	fileName = fmt.Sprintf("%s/%s", w.wasabiClient.Endpoint, fileName)

	return fileName, nil
}

func (w *wasabiRepoImpl) UploadImage(ctx context.Context, folder string, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}

	extension := extensionFromImageMime(file.Header.Get("Content-Type"))
	if extension == "" {
		return "", fmt.Errorf("invalid file type: %s", file.Header.Get("Content-Type"))
	}

	fileName := fmt.Sprintf("%s/pictures/%s%s", folder, uuid.NewString(), extension)

	input := &s3.PutObjectInput{
		Bucket: aws.String(w.wasabiClient.Bucket),
		Key:    &fileName,
		Body:   f,
	}

	_, err = w.wasabiClient.S3.PutObject(input)
	if err != nil {
		return "", err
	}

	fileName = fmt.Sprintf("%s/%s", w.wasabiClient.Endpoint, fileName)

	return fileName, nil
}

func (w *wasabiRepoImpl) DeleteFile(ctx context.Context, bucketName string, fileName string) error {
	_, err := w.wasabiClient.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(w.wasabiClient.Bucket),
		Key:    aws.String(fileName),
	})

	if err != nil {
		return err
	}

	return nil
}

func NewWasabiRepo(wasabiClient *wasabi.Client) WasabiRepo {
	return &wasabiRepoImpl{wasabiClient: wasabiClient}
}
