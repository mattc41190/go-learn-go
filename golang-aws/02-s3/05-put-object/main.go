package main

import (
	"fmt"
	"strings"
	"math/rand"
	"strconv"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"flag"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// Allow user to pass flags
	profile := flag.String("profile", "default", "the AWS profile you wish to use for interacting with AWS")
	region := flag.String("region", "us-east-1", "The AWS region you wish to use to initialize the session")
	bucket := flag.String("bucket", "mattcale-personal-files", "The name of the S3 bucket you wish to act upon")
	key := flag.String("key", "mattcale-data" + strconv.Itoa(rand.Intn(100000)) , "The key / name of the object you intend to put into s3")
	flag.Parse()

	// Set an Enviornment Variable
	os.Setenv("AWS_PROFILE", *profile)

	// Create AWS session config
	conf := &aws.Config{
		Region:      region,
		Credentials: credentials.NewSharedCredentials("", *profile),
	}

	sess := session.New(conf)

	uploader := s3manager.NewUploader(sess)

	obj := &s3manager.UploadInput{
		Bucket: bucket,
		Key:    key,
		Body:   strings.NewReader("Hi Matt!"),
	}

	resp, err := uploader.Upload(obj)

	if err != nil {
		fmt.Println(err)
		panic("Failed to write object to bucket: " + *bucket)
	}

	fmt.Println(resp)



}