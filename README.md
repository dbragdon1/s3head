# Background

I work with a lot of csv data, most of which are stored in s3, and wanted an easy way to get a small preview of files before downloading the full thing. Unfortunately, the aws cli does not make it easy to partially read files into a buffer, so I created this tool to make it a lot easier. It works similar to the linux `head` command, but does not have the exact same interface, so be warned. 


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

s3head -n 1000 s3://my-bucket/path/to/my/key \
    | xsv select "firstname,lastname" \
    | xsv sample 10
```
