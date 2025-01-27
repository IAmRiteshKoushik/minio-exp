package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// Base configs
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := "test-user"
	secretAccessKey := "test-key"
	useSSL := true
	bucketName := "sample-bucket"
	// location := "dev"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println("[MINIO INITIALIZATION FAILED]")
		log.Fatalln(err)
	}
	fmt.Println("[MINIO INITIALIZATION SUCCESS]")
	log.Printf("%+v\n", minioClient)

	// Bucket Creation (if not exists)
	// err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	// if err != nil {
	// 	// Check for existance
	// 	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	// 	if errBucketExists == nil && exists {
	// 		log.Printf("We already own %s\n", bucketName)
	// 	} else {
	// 		log.Fatalln(err)
	// 	}
	// } else {
	// 	log.Printf("Successfully created %s\n", bucketName)
	// }

	// fileUploader.go
	err = upload(ctx, *minioClient, bucketName)
	if err != nil {
		fmt.Printf("\n[UPLOAD FAILED]\n")
		log.Fatalln(err)
	}
}
