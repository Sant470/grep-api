package s3

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sant470/grep-api/config"
	"github.com/sant470/grep-api/types/consts"
)

// Client ...
type Client struct {
	s3Client *s3.S3
	session  *session.Session
	lgr      *log.Logger
}

// NewClient ...
func NewClient(lgr *log.Logger) (*Client, error) {
	cred := config.GetAppConfig("config.yaml", ".").AWS
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(consts.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cred.AccessKeyID, cred.SecretAccessKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting aws session %v", err)
	}
	s3Client := s3.New(s)
	return &Client{s3Client: s3Client, session: s, lgr: lgr}, nil
}

// // GetS3ObjectBuffer copy s3 file in memory ...
func (cli *Client) GetS3ObjectBuffer(bucket, path string) ([]byte, error) {
	result, err := cli.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting s3 object :%v", err)
	}
	defer result.Body.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, result.Body); err != nil {
		return nil, fmt.Errorf("error copying file into buffer :%v", err)
	}
	return buf.Bytes(), nil
}
