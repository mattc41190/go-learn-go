package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
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
	key := flag.String("key", "mattcale-data" , "The key / name of the object you intend to delete from s3")
	flag.Parse()

	// Set an Enviornment Variable
	os.Setenv("AWS_PROFILE", *profile)

	// Create AWS session config
	conf := &aws.Config{
		Region:      region,
		Credentials: credentials.NewSharedCredentials("", *profile),
	}

	sess := session.New(conf)

	store := s3.New(sess)
	
	_, err := store.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(*bucket), Key: aws.String(*key)})


	if err != nil{
		panic(err)
	}

}
