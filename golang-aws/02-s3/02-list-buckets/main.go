package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// This program will create an s3 bucket `--bucket` in the desired region `--region` 

func main() {

	// Allow user to pass AWS configuration via basic AWS flags
	var profile = flag.String("profile", "default", "Name of AWS profile to use")
	var region = flag.String("region", "us-east-1", "Name of AWS region to use")
	
	// Create pointers to the configuration data the application was passed
	flag.Parse()

	// Set an environment variable based on the profile passed to the command line
	os.Setenv("AWS_PROFILE", *profile)

	// Create an AWS Config object. Taking mostly default values, but including a specific AWS_PROFILE
	conf := aws.Config{
		Region:      aws.String(*region),
		Credentials: credentials.NewSharedCredentials("", *profile),
	}

	// Create an AWS session
	sess, err := session.NewSession(&conf)

	// Check to see if there was an error opening the session
	if err != nil {
		fmt.Println(err)
	}

	// Create an s3 client and call it `store`
	store := s3.New(sess)

	fmt.Println("The API Version for the S3 instance you created is:", store.APIVersion)

	// Create a bucket
	resp, err := store.ListBuckets(nil)

	// Check to see if there was an error creating the bucket
	if err != nil {
		fmt.Println(err)
	}

	for _, bucket := range resp.Buckets {
		fmt.Println(bucket.String())
	}

}
