# gs3pload [![GoDoc](https://godoc.org/github.com/fern4lvarez/gs3pload?status.svg)](https://godoc.org/github.com/fern4lvarez/gs3pload)

**gs3pload** is a command line tool to upload files to multiple S3 or Google Storage buckets at once.
> Disclaimer: This tool do not aim to replace `gsutil` or `s3cmd`, so if you are missing many
> features here this is probably not what you're looking for.

Prerequisites
-------------

### gsutil

~~~
$ sudo pip install gsutil
~~~

Install
-------

~~~
$ go get github.com/fern4lvarez/gs3pload
~~~

or [download the binary file](http://gobuild.io/github.com/fern4lvarez/gs3pload) for your architecture.


Configure
---------

Create a `~/.gs3pload/envs.json` file with desired environments:

~~~
[
{
    "name": "live.com",
    "type": "s3"
},
{
    "name": "dev.com",
    "type": "gs"
}
]
~~~

### Amazon S3 Environments

For each environment based on S3 you need to create a file called `<Environment Name>.boto`
inside the `.gs3pload` directory with this format, fulfilling the required credentials (see https://github.com/fern4lvarez/gs3pload/issues/1):

~~~
[Credentials]
aws_access_key_id = <PLACE YOUR KEY ID HERE>
aws_secret_access_key = <PLACE YOUR ACCESS KEY HERE>

[Boto]
https_validate_certificates = False
~~~

### Google Storage Environments

For each environment based on GS, run this command and follow all steps:

~~~
BOTO_CONFIG=~/.gs3pload/<Environment Name>.boto gsutil config
~~~


##Usage

```
Usage:
  gs3pload push <bucket> <name>... [-r | --recursive] [-p | --public] [-b | --backup]
  gs3pload -h | --help
  gs3pload -v | --version

Options:
  -h --help        Show help.
  -p --public      Set files as public.
  -r --recursive   Do a recursive copy.
  -b --backup      Create backup of pushed files if they exist.
  -v --version     Show version.

```


## Development

This repository includes a ready-to-use Vagrant box that provides the required environment
to work on development and testing of `gs3pload`:

~~~
$ vagrant up
$ vagrant ssh
$ sudo su
$ gs3pload
~~~


##License
---------
gs3pload is MIT licensed.
