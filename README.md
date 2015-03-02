# gs3pload [![GoDoc](https://godoc.org/github.com/fern4lvarez/gs3pload?status.svg)](https://godoc.org/github.com/fern4lvarez/gs3pload)

**gs3pload** is a command line tool to upload files to multiple S3 or Google Storage buckets at once.
> Disclaimer: This tool does not aim to replace `gsutil` or `s3cmd`, so if you are missing many
> features here this is probably not what you're looking for.

Prerequisites
-------------

### gsutil

~~~
sudo pip install gsutil
~~~

### swift & keystone

> For OpenStack Swift support only

~~~
sudo pip install python-swiftclient python-keystoneclient
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
},
{
    "name": "test.com",
    "type": "swift"
}
]
~~~

Current supported types:

* `s3`: AWS S3
* `gs`: Google Storage
* `swift`: OpenStack Swift

If you want to use a custom environments file, you can use the `--envs` flag:

~~~
gs3pload push --envs PATH_TO_ENVS_FILE bucket file
~~~

If for some reason you want to push only to single environment, use `--env` flag:

~~~
gs3pload push --env ENV_NAME bucket file
~~~

where `ENV_NAME` needs to be defined in default or custom config file.

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

### OpenStack Swift

> Experimental

`gs3pload` brings experimental support to OpenStack Swift. You can upload files and
directories as always, but please keep in mind:

* `--public` flag is not supported: OpenStack Swift handles permission by buckets, not files.
* `--backup` flag is not supported.
* `--recursive` flag is not required, all uploads are recursive by default in OpenStack Swift.

For each environment based on Swift you need to create a file called `<Environment Name>.boto`
inside the `.gs3pload` directory with this format, fulfilling the required credentials:

~~~
# *NOTE*: Using the 2.0 *auth api* does not mean that compute api is 2.0.  We
# will use the 1.1 *compute api*
OS_AUTH_URL=https://foo.com:5000/v2.0
OS_AUTH_VERSION=2.0

# With the addition of Keystone we have standardized on the term **tenant**
# as the entity that owns the resources.
OS_TENANT_ID=xxx
OS_TENANT_NAME="yyy"

# In addition to the owning entity (tenant), openstack stores the entity
# performing the action as the **user**.
OS_USERNAME="info@example.org"

# With Keystone you pass the keystone password.
OS_PASSWORD="secret"
~~~

##Usage

```
Usage:
  gs3pload push [--envs <file>] [--env <name>] <bucket> <name>... [-r | --recursive] [-p | --public] [-b | --backup]
  gs3pload -h | --help
  gs3pload -v | --version

Options:
  -h --help        Show help.
  --envs <file>    Use a custom environments configuration.
  -e --env <name>  Environment name.
  -p --public      Set files as public.
  -r --recursive   Do a recursive copy.
  -b --backup      Create backup of pushed files if they exist.
  -v --version     Show version.

```


## Development

This repository includes a ready-to-use Vagrant box that provides the required environment
to work on development and testing of `gs3pload`:

~~~
vagrant up
vagrant ssh
sudo su
gs3pload
~~~


##License
---------
gs3pload is MIT licensed.
