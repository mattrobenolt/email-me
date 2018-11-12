#!/bin/bash
set -xe

rm -rf bin/
docker build --rm -t email-me .
docker run --rm -v $PWD:/usr/src/email-me email-me
