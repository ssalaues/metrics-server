#!/bin/bash
set -ex

LOCATION=${1:-/usr/local/bin}

go build
cp metrics-server ${LOCATION}
cp startup.sh ${LOCATION}
cp generate-data.sh ${LOCATION}
mkdir -p /data

cp systemd-config/metrics.service /etc/systemd/system
# refresh config changes from filesystem
systemctl daemon-reload
# enable server on startup
systemctl enable metrics.service
# start server
systemctl start metrics.service
# get operational status
systemctl status