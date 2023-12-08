#!/bin/bash
##                                              ##
# Use this script for generating live dummy data #
##                                              ##
set -e
FILE=${1:-/data/metrics_from_special_app.txt}
INTERVAL=1
GREEN="\e[32m"
REDBOLD="\e[1;31m"
ENDCOLOR="\e[0m"

printf "${GREEN}This will write dummy data to ${FILE} every ${INTERVAL}s${ENDCOLOR}\n"
printf "${REDBOLD}press ctl+c to stop${ENDCOLOR}\n"

storage=100
while true ; do
    cat <<- EOF > $FILE
	CPU=$((1 + $RANDOM % 1000))m
	MEMORY=$((1 + $RANDOM % 3000))Mi
	LATENCY=$((1 + $RANDOM % 10))ms
	STORAGE=${storage}Gi
	EOF
    ((storage++))
    if [ $storage == 199 ]; then
        storage=100
    fi
    sleep ${INTERVAL}
done
