package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ungerik/go-dry"

)

type S3 struct {
	bucket     string
	downloader *s3manager.Downloader
	uploader    *s3manager.Uploader
	client     *s3.S3
}

// New returns S3 struct
func NewClient(accessKey, secretKey, bucket, region string) *S3 {
	cred := credentials.NewStaticCredentials(accessKey, secretKey, "")
	configs := session.New(&aws.Config{
		Credentials: cred,
		Region:      aws.String(region),
	})

	return &S3{
		bucket: bucket,
		// チューニングのためバッファサイズを増やす
		downloader: s3manager.NewDownloader(configs, func(d *s3manager.Downloader) {
			d.PartSize = 1024 * 1024 * 1024 * 1
			d.Concurrency = 20
		}),
		uploader: s3manager.NewUploader(configs),
		client: s3.New(configs),
	}
}

// Exists returns file exists
func (s S3) Exists(key string) bool {
	res, err := s.client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(key),
	})

	if err != nil {
		fmt.Println("Failed to create file", err)
		return false
	}

	if dry.IsZero(len(res.Contents)) {
		fmt.Println("path=%s", key)
		return false
	}
	return true
}

func main() {
	s3 := NewClient("AKIAJGFVPMAPJEEMUUKA", "OknwHzAIf45m7rnD7JfhVUIg/jRNdF+0SWKyvNxm", "saigon-lap-thumbnails", "ap-northeast-1")

	/* check exists file */
	filepath := "b123456/z97.jpg"
	if !s3.Exists(filepath) {
		fmt.Println("file not exists ", filepath)

		file, err := os.Open("./z97.jpg")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()

		result, err := s3.uploader.Upload(&s3manager.UploadInput{
			Body: file,
			Bucket: aws.String("saigon-lap-thumbnails"),
			Key: aws.String("b123456/z97.jpg"),
		})

		if err != nil {
			fmt.Println("Failed to upload", err)
			return
		}

		fmt.Printf("Successful uploaded. %+v ", result)

		return
	}

	fmt.Println("file already exists ", filepath)
}