#!/bin/bash

LOCATION=${1:-/usr/local/bin}
${LOCATION}/generate-data.sh data/metrics_from_special_app.txt &
pid=$!
trap "/bin/kill $pid; exit 1" SIGINT
${LOCATION}/metrics-server
printf "Server failed stopping data generation"
