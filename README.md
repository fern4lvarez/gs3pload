gs3pload
========

[Documentation online](http://godoc.org/github.com/fern4lvarez/gs3pload)

**gs3pload** is a command line tool to upload files to multiple S3 or Google Storage buckets at once.

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
inside the `.gs3pload` directory with this format, fulfilling the required credentials:

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
-------
```
Usage:
  gs3pload push <bucket> <name>... [-p | --public]
  gs3pload -h | --help
  gs3pload -v | --version

Options:
  -h --help        Show help.
  -p --public      Set files as public.
  -v --version     Show version.
```

##License
---------
gs3pload is MIT licensed.
