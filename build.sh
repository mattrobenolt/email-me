#!/bin/bash
set -xe

rm -rf bin/
docker build --rm -t email-me .
docker run --rm -v $PWD:/go/src/app email-me
