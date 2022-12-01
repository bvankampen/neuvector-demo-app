#!/bin/bash

docker build -t $APP_IMAGE .
docker push $APP_IMAGE