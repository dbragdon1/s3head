package main

import (
	"dbragdon1/s3head/file"
	"dbragdon1/s3head/utils"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Usage() {
	fmt.Println("Usage: s3head [OPTION]... S3_URI")
	flag.PrintDefaults()
	fmt.Println("  S3_URI")
	fmt.Println("\tPath to s3 object (e.g. s3:/my-bucket/path/to/key)")
}

func main() {

	numLines := flag.Int("n", 10, "Number of lines to grab")

	allLines := flag.Bool("a", false, "Whether to grab all lines")

	help := flag.Bool("h", false, "Print help")

	flag.Parse()

	if *help {
		Usage()
		os.Exit(0)

	}

	positionals := flag.Args()

	var s3_uri string

	if len(positionals) < 1 {
		fmt.Println("No s3 URI provided")
		flag.Usage()
		os.Exit(1)
	} else {
		s3_uri = positionals[0]
	}

	s3_object, err := utils.ParseS3URI(s3_uri)

	if err != nil {
		fmt.Printf("Trouble parsing s3 URI: %v \n", err)
		os.Exit(1)
	}

	sess, err := session.NewSession()

	if err != nil {
		fmt.Printf("Couldn't authenticate to AWS: %v \n", err)
		os.Exit(1)
	}

	s3_svc := s3.New(sess)

	req, err := s3_svc.GetObject(&s3_object)

	if err != nil {
		fmt.Printf("Found issue when attempting GetObject: %v \n", err)
		os.Exit(1)
	}

	defer req.Body.Close()

	f := file.NewS3File(s3_uri, *numLines, *allLines, req)

	err = f.Iter(os.Stdout)

	if err != nil {
		fmt.Println(err)
	}

}
