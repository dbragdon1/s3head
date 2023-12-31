# Background

I work with a lot of csv data, most of which is stored in s3, and wanted an easy way to get a small preview of files before downloading the full thing. Unfortunately, the aws cli does not make it easy to partially read files into a buffer, so I created this tool to make it a lot easier. It works similar to the linux `head` command, but does not have the exact same interface, so be warned. 

# Installation

```bash
go build ./
```

Then copy the `s3head` binary to your bin directory of choice.

# Usage

`s3head` is indifferent to the type of file you pass in. It simply iterates over the lines in the file. 

## AWS Authentication

You must obviously be authenticated to AWS to use this command. `s3head` defers AWS authentication to the steps taken by the `NewSession` object provided by `aws-sdk-go/aws/session`.

## Default behavior

```bash
# Prints the first five lines in the file

s3head s3://my-bucket/path/to/my/key
```

## Specify number of lines

```bash
# Prints the first 10 lines of the file

s3head -n 10 s3://my-bucket/path/to/my/key
```

## Grab the entire file
```
s3head -a s3://my-bucket/path/to/my/key
```

## Pipe the output

```bash
# pipes output to xsv
# https://github.com/BurntSushi/xsv

s3head -n 1000 s3://my-bucket/path/to/my/csv/file \
    | xsv select "firstname,lastname" \
    | xsv sample 10
```

```bash
# pipes output to jq
# https://github.com/jqlang/jq

s3head -a s3://my-bucket/path/to/my/json/file \
    | jq .my_key
```

## Save to a file

```bash
s3head -a s3://my-bucket/path/to/my/csv/file > myfile.csv
```

## Automatic GZIP Decompression

```bash
s3head -n 1000 s3://my-bucket/path/to/my/file.csv.gz \
    | xsv headers
```

### Why not use the the AWS CLI `s3api get-object` command instead?

1. For some reason, attempting to pipe the stream from `aws s3-api get-object` consistently results in a Broken Pipe error, which doesnt look very clean
2. Working with gzipped data is a lot more concise with `s3head`:

```bash
aws s3api get-object --bucket my-bucket --key path/to/my/key.gz /dev/stdout \
    | gunzip -c \
    | head -n 2
```

versus

```bash
s3head -n 2 s3://my-bucket/path/to/my/key.gz
```

# Similar Projects

The following projects seem to attempt to solve a similar problem as `s3head`. Why use `s3head` over these other solutions? Perhaps you like the api better, or perhaps it **feels** faster because it's written in golang and feels more "modern". 


[s3streamcat](https://github.com/samarthg/s3streamcat)

[s3curl](https://github.com/rtdp/s3curl)
