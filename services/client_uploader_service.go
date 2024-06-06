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
)

const (
	projectID  = "capstone-trial-425115"
	bucketName = "capstone-trial-425115"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func NewClientUploader() *ClientUploader {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "capstone-trial-425115-dbf96ccd90f6.json")
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "productImage/",
	}
}

func (c *ClientUploader) UploadImage(file multipart.File, object string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (c *ClientUploader) DeleteImage(object string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Delete the object from the bucket.
	o := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object)
	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object.Delete: %v", err)
	}

	return nil
}
