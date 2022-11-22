wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
wget https://go.dev/dl/go1.17.11.linux-amd64.tar.gz
wget https://go.dev/dl/go1.16.15.linux-amd64.tar.gz
wget https://go.dev/dl/go1.15.linux-amd64.tar.gz
wget https://go.dev/dl/go1.14.15.linux-amd64.tar.gz
wget https://go.dev/dl/go1.13.11.linux-amd64.tar.gz
wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz

mkdir /usr/local/lib/go1.13.11 /usr/local/lib/go1.14.15 /usr/local/lib/go1.15 /usr/local/lib/go1.16.15 /usr/local/lib/go1.17.11 /usr/local/lib/go1.18.3 /usr/local/lib/go1.19.2

tar zxf go1.13.11.linux-amd64.tar.gz -C /usr/local/lib/go1.13.11/
tar zxf go1.14.15.linux-amd64.tar.gz -C /usr/local/lib/go1.14.15/
tar zxf go1.15.linux-amd64.tar.gz -C /usr/local/lib/go1.15/
tar zxf go1.16.15.linux-amd64.tar.gz -C /usr/local/lib/go1.16.15/
tar zxf go1.17.11.linux-amd64.tar.gz -C /usr/local/lib/go1.17.11/
tar zxf go1.18.3.linux-amd64.tar.gz -C /usr/local/lib/go1.18.3/
tar zxf go1.19.2.linux-amd64.tar.gz -C /usr/local/lib/go1.19.2/

cd /usr/local/bin
vim goenv.sh

#!/bin/bash 
version=$1

if [[ ${version} == "" ]]; then
        version="1.18.3"
fi

GOROOTTMP=/usr/local/lib/go${version}

if [[ ! -d ${GOROOTTMP} ]]; then
        echo "go ${version} not exist, hoose another version !"
        echo "available go version: 1.13.11 | 1.14.15 | 1.15 | 1.16.15 | 1.17.11 | 1.18.3"
else
        if [[ -L /usr/local/lib/go ]]; then
                rm -rf /usr/local/lib/go
        fi
        ln -s ${GOROOTTMP}/go /usr/local/lib
fi

chmod u+x goenv.sh 

vim ~/.bash_profile

export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/lib/go/bin:$PATH
export GO111MODULE=on
export GOPATH=/appcom/gopath
export GOBIN=${GOPATH}/bin
export GOPROXY=https://goproxy.cn
export GOROOT=/usr/local/lib/go

source ~/.bash_profile

goenv.sh xxxx