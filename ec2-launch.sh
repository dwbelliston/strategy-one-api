#!/bin/bash

# https://jonathanmh.com/deploying-go-apps-systemd-10-minutes-without-docker/

yum update -y
yum install -y git
yum install -y golang

# GOROOT is the location where Go package is installed on your system
export GOROOT=/usr/lib/golang
# GOPATH is the location of your work directory
export GOPATH=$HOME/projects
# PATH in order to access go binary system wide
export PATH=$PATH:$GOROOT/bin

go get github.com/dwbelliston/strategy-one-api
cd $GOPATH/src/github.com/dwbelliston/strategy-one-api
go build

touch /lib/systemd/system/api.service

tee /lib/systemd/system/api.service >/dev/null <<EOF
  [Unit]
  Description=strategy-one-api

  [Service]
  Type=simple
  Restart=always
  RestartSec=5s
  ExecStart=$GOPATH/src/github.com/dwbelliston/strategy-one-api/strategy-one-api

  [Install]
  WantedBy=multi-user.target
EOF

service api start
