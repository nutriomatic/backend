package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	projectID  = "c241-ps219"
	bucketName = "nutrio-storage"
)

type ClientUploader struct {
	client      *storage.Client
	bucketName  string
	productPath string
	userPath    string
	paymentPath string
}

func NewClientUploader() *ClientUploader {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "c241-ps219-ab6446ebf2e7.json")
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &ClientUploader{
		client:      client,
		bucketName:  bucketName,
		productPath: "productImage/",
		userPath:    "profileImage/",
		paymentPath: "proofPayment/",
	}
}

func (cu *ClientUploader) ProcessImage(c echo.Context, pathPrefix string) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	uuid := uuid.New().String()
	fileName := fmt.Sprintf("%s-%s", pathPrefix, uuid)
	err = cu.uploadImage(src, pathPrefix+fileName)
	if err != nil {
		return "", err
	}

	return pathPrefix + fileName, nil
}

func (cu *ClientUploader) uploadImage(file multipart.File, object string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	wc := cu.client.Bucket(cu.bucketName).Object(object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (cu *ClientUploader) DeleteImage(pathPrefix, object string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	o := cu.client.Bucket(cu.bucketName).Object(pathPrefix + object)
	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object.Delete: %v", err)
	}

	return nil
}

func (cu *ClientUploader) ProcessImageUser(c echo.Context) (string, error) {
	return cu.ProcessImage(c, cu.userPath)
}

func (cu *ClientUploader) ProcessImageProduct(c echo.Context) (string, error) {
	return cu.ProcessImage(c, cu.productPath)
}

func (cu *ClientUploader) ProcessImageProof(c echo.Context) (string, error) {
	return cu.ProcessImage(c, cu.paymentPath)
}

func (cu *ClientUploader) DeleteImageProduct(object string) error {
	return cu.DeleteImage(cu.productPath, object)
}

func (cu *ClientUploader) DeleteImageUser(object string) error {
	return cu.DeleteImage(cu.userPath, object)
}

func (cu *ClientUploader) DeleteImageProof(object string) error {
	return cu.DeleteImage(cu.paymentPath, object)
}
