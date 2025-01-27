package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
)

func upload(ctx context.Context, minioClient minio.Client, bucketName string) error {
	objectName := "testdata"
	filePath := "./"
	contentType := "application/octet-stream"

	// Upload the test file
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}
	fmt.Println("[UPLOAD SUCCESSFUL]")
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return nil
}
