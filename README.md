gs3pload
========

[Documentation online](http://godoc.org/github.com/fern4lvarez/gs3pload)

**gs3pload** is a command line too too upload files to different S3 or Google Storage buckets at once.

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
    "name": "live",
    "type": "s3",
    "domain": "live.com"
},
{
    "name": "dev",
    "type": "gs",
    "domain": "dev.com"
}
]
~~~


### S3 environments

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
  gs3pload push <bucket> <name>... [--public]
  gs3pload -h | --help
  gs3pload --version

Options:
  -h --help     Show help.
  --public      Set files as public.
  --version     Show version.
```

##License
---------
gs3pload is MIT licensed.