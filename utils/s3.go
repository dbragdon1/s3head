package utils

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3URIHelp() string {
	return "Format should be 's3://my-bucket/path/to/my/key'"
}
func ParseS3URI(s3_uri string) (s3.GetObjectInput, error) {

	bucket_pattern := regexp.MustCompile(`^s3://([^/]+)/`)

	key_pattern := regexp.MustCompile(`^s3://[^/]+(/.+)$`)

	bucket_matches := bucket_pattern.FindStringSubmatch(s3_uri)

	key_matches := key_pattern.FindStringSubmatch(s3_uri)

	var s3Bucket string

	var s3Key string

	if len(bucket_matches) >= 2 {
		s3Bucket = bucket_matches[1]
	} else {
		return s3.GetObjectInput{}, errors.New(fmt.Sprintf("Could not parse uri for bucket name: %s", getS3URIHelp()))
	}

	if len(key_matches) >= 2 {
		s3Key = key_matches[1]
	} else {
		return s3.GetObjectInput{}, errors.New(fmt.Sprintf("Could not parse uri for key: %s", getS3URIHelp()))
	}

	return s3.GetObjectInput{Bucket: &s3Bucket, Key: &s3Key}, nil

}

func GetExtention(s3_uri string) string {

	//re := regexp.MustCompile(`\.(.*?)$`)

	re := regexp.MustCompile(`\.([^.]*)$`)

	match := re.FindStringSubmatch(s3_uri)

	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}
