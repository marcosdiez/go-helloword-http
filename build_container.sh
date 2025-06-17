#!/bin/bash
set -x -e
./build.sh
docker build . -t marcosdiez/helloworld-http