#!/bin/bash
set -ex

LOCATION=${1:-/usr/local/bin}

# if it's already installed kill the sub process so the systemd service stops quicker/reliably
kill $(ps aux | grep -m1 generate-data.sh | awk '{ print $2 }') || true
systemctl stop metrics.service || true

go build
cp metrics-server ${LOCATION}
cp startup.sh ${LOCATION}
cp generate-data.sh ${LOCATION}
mkdir -p /data
touch /data/metrics_from_special_app.txt

cp systemd-config/metrics.service /etc/systemd/system
# refresh config changes from filesystem
systemctl daemon-reload
# enable server on startup
systemctl enable metrics.service
# start server
systemctl start metrics.service
# get operational status
systemctl status metrics.service
