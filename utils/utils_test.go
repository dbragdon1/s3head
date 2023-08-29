package utils

import (
	"testing"
)

func TestParseS3URIValid(t *testing.T) {

	uri := "s3://my-bucket/path/to/key/1"

	parsed, _ := ParseS3URI(uri)

	want_bucket := "my-bucket"

	want_key := "/path/to/key/1"

	if *parsed.Bucket != want_bucket {
		t.Fatalf("Wanted %s, got %s", want_bucket, *parsed.Bucket)
	}

	if *parsed.Key != want_key {
		t.Fatalf("Wanted %s, got %s", want_key, *parsed.Key)
	}

}

func TestParseS3URIInvalidURI(t *testing.T) {

	uri := "s3:/my-bucket/path/to/key/1"

	want_err := "Could not parse uri for bucket name: Format should be 's3://my-bucket/path/to/my/key'"

	_, err := ParseS3URI(uri)

	if err.Error() != want_err {
		t.Fatalf("Wanted %s, got %s", want_err, err.Error())
	}

}

func TestGetExtention(t *testing.T) {

	uri := "s3://my-bucket/my/path/file.csv"

	extension := GetExtention(uri)

	want_extension := "csv"

	if extension != want_extension {
		t.Fatalf("Wanted %s, got %s", want_extension, extension)
	}

}

//func TestGetExtentionMultiple(t *testing.T) {
//
//	uri := "s3://my-bucket/my/path/file.tar.gz"
//
//	extension := GetExtention(uri)
//
//	want_extension := "tar.gz"
//
//	if extension != want_extension {
//		t.Fatalf("Wanted %s, got %s", want_extension, extension)
//	}
//}

func TestGetExtentionNone(t *testing.T) {

	uri := "s3://my-bucket/my/path/file"

	extension := GetExtention(uri)

	want_extension := ""

	if extension != want_extension {
		t.Fatalf("Wanted %s, got %s", want_extension, extension)
	}
}
