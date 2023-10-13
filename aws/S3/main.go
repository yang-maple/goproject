package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	//output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	//	Bucket: aws.String("whalescheduler"),
	//})
	output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	//for _, object := range output.Contents {
	//	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	//}
	for _, bucket := range output.Buckets {
		fmt.Println(bucket.Name, bucket.CreationDate)
	}
}
