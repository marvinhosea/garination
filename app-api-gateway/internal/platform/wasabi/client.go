package wasabi

import (
	"fmt"
	"garination.com/gateway/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	wasabiUrl = "https://s3.%s.wasabisys.com"
	BUCKET    = "garination-dev-1"
)

type Client struct {
	S3       *s3.S3
	Endpoint string
	Bucket   string
}

func NewWasabiClient(c config.S3) (*Client, error) {

	s3Config := aws.Config{
		Endpoint:         aws.String(c.Endpoint),
		Region:           aws.String(c.Region),
		Credentials:      credentials.NewStaticCredentials(c.AccessId, c.AccessKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess, err := session.NewSession(&s3Config)
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)

	listBucketsInput := &s3.ListBucketsInput{}

	result, err := s3Client.ListBuckets(listBucketsInput)
	if err != nil {
		return nil, err
	}

	var bucketExists bool
	for _, b := range result.Buckets {
		if *b.Name == BUCKET {
			bucketExists = true
			break
		}
	}

	// create bucket if it doesn't exist
	if !bucketExists {
		_, err := s3Client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(BUCKET),
		})

		if err != nil {
			return nil, err
		}
	}

	wasabiClient := &Client{
		S3:       s3Client,
		Endpoint: fmt.Sprintf(wasabiUrl, c.Region),
		Bucket:   BUCKET,
	}

	return wasabiClient, nil
}
