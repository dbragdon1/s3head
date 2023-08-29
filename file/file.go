package file

import (
	"bufio"
	"compress/gzip"
	"dbragdon1/s3head/utils"
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
)

type S3File struct {
	S3URI    string
	NumLines int
	AllLines bool
	Obj      *s3.GetObjectOutput
}

func NewS3File(s3_uri string, numLines int, allLines bool, object *s3.GetObjectOutput) S3File {

	return S3File{S3URI: s3_uri, NumLines: numLines, AllLines: allLines, Obj: object}

}

func iter(scanner *bufio.Scanner, numLines int, allLines bool) {
	if allLines {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		for curr_line := 0; curr_line < numLines; curr_line++ {

			if scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	}

}
func (f *S3File) StandardIter() *bufio.Scanner {

	return bufio.NewScanner(f.Obj.Body)

}

func (f *S3File) GzIter() (*bufio.Scanner, error) {
	uncompressed, err := gzip.NewReader(f.Obj.Body)

	if err != nil {
		return &bufio.Scanner{}, err
	}

	buf := bufio.NewScanner(uncompressed)

	return buf, nil
}

func (f *S3File) Iter() error {

	ext := utils.GetExtention(f.S3URI)

	var err error

	var buf *bufio.Scanner

	switch ext {

	case "gz":
		{
			buf, err = f.GzIter()

		}
	default:
		buf = f.StandardIter()
	}

	iter(buf, f.NumLines, f.AllLines)

	return err

}
