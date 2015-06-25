#!/bin/bash
set -x

rm -rf bin/
docker build --rm -t email-me .
docker run --rm -v $PWD:/go/src/app -it email-me
