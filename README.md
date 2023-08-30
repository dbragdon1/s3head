# Background

I work with a lot of csv data, most of which is stored in s3, and wanted an easy way to get a small preview of files before downloading the full thing. Unfortunately, the aws cli does not make it easy to partially read files into a buffer, so I created this tool to make it a lot easier. It works similar to the linux `head` command, but does not have the exact same interface, so be warned. 


# Installation

```bash
go build ./
```

Then copy the `s3head` binary to your bin directory of choice.

# Usage

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

## Automatic GZIP Decompression

Note: `.tar.gz` format not implemented yet

```bash
s3head -n 1000 s3://my-bucket/path/to/my/file.csv.gz \
    | xsv headers
```

# TODO

Implement Automatic untarring if file is in `.tar` format

# Similar Projects

The following projects seem to attempt to solve a similar problem as `s3head`. Why use `s3head` over these other solutions? Perhaps you like the api better, or perhaps it **feels** faster because it's written in golang and feels more modern. 


[s3streamcat](https://github.com/samarthg/s3streamcat)
[s3curl](https://github.com/rtdp/s3curl)
