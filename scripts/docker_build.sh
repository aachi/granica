#!/bin/shs
# Build
./build.sh

# Docker build
docker login
docker build -t adrianpksw/granica .
#docker push adrianpksw/granica
