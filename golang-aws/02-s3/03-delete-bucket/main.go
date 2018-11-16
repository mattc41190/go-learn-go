package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

)

// This program allows a user to delete an arbitrary bucket

func main() {

	// Allow user to pass values to program.
	profile := flag.String("profile", "default", "The AWS profile you wish to use")
	region := flag.String("region", "us-east-1", "The AWS regions you wish to use")
	bucket := flag.String("bucket", "matt-personal-files", "The name of S3 bucket you wish to act upon")
	flag.Parse()

	// Set an environment variable based on the profile passed to the command line
	os.Setenv("AWS_PROFILE", *profile)

	// Create an AWS session config
	conf := aws.Config{
		Region:      aws.String(*region),
		Credentials: credentials.NewSharedCredentials("", *profile),
	}

	sess, err := session.NewSession(&conf)

	if err != nil {
		fmt.Println(err)
		panic("Could not start AWS session. Exiting.")
	}

	store := s3.New(sess)

	resp, err := store.DeleteBucket(&s3.DeleteBucketInput{Bucket: bucket})

	if err != nil {
		fmt.Println(err)
		panic("Unable to delete bucket: " + *bucket)
	}

	fmt.Println("Deleted: ", resp.String())

}
