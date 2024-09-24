package AmazonS3FileUpload

import (
	"fmt"
	"github.com/yusufkorucu/amazonS3Api"
	"log"
)

func main() {

	amazonS3ApiConfig := amazonS3Api.S3ApiConfig{
		CdnUrl:          "",
		BucketName:      "",
		SecretAccessKey: "",
		AccessKeyId:     "",
	}

	amazonS3Api.SetS3ApiConfig(amazonS3ApiConfig)

	s3Client := amazonS3Api.NewAmazonS3ApiClient()
	if s3Client == nil {
		log.Fatalf("Failed to create S3 client")
	}

	uploadPath := "your/upload/path"
	fileName := "example.txt"
	fileContent := []byte("Hello, S3!")

	downloadUrl, err := s3Client.UploadFile(uploadPath, fileName, fileContent)
	if downloadUrl == nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	fmt.Printf("File uploaded successfully! Download URL: %s\n", downloadUrl)
}
