#!/bin/bash

# Obtain credentials
aws ecr get-login-password --region us-west-1 | docker login --username AWS --password-stdin 213288384225.dkr.ecr.us-west-1.amazonaws.com/hello-world

docker build -t hello-world .

docker tag hello-world:latest 213288384225.dkr.ecr.us-west-1.amazonaws.com/hello-world:latest

docker push 213288384225.dkr.ecr.us-west-1.amazonaws.com/hello-world:latest