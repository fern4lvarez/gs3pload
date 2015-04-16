#!/usr/bin/env bash
set -eu

apt-get update

DEBIAN_FRONTEND=noninteractive apt-get install -y -q \
	python-dev python-pip git

pip install gsutil python-swiftclient python-keystoneclient

GOVERSION=1.4.1
FILENAME="godeb-amd64.tar.gz"
GODEB="https://godeb.s3.amazonaws.com/${FILENAME}"
wget ${GODEB}
tar -C ${HOME} -xzf ${FILENAME}
${HOME}/godeb install ${GOVERSION}



export GOPATH=${HOME}
export PATH=${PATH}:${GOPATH}/bin
echo "export GOPATH=${HOME}" >> ${HOME}/.bashrc
echo "export PATH=$PATH:${GOPATH}/bin" >> ${HOME}/.bashrc
mkdir -p ${GOPATH}/src/github.com/fern4lvarez/gs3pload
cd ${GOPATH}/src/github.com/fern4lvarez/gs3pload
cp -r /vagrant/* .

go get ./...
gs3pload -h