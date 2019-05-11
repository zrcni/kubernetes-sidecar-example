#!/bin/bash
version=$1
projectId=$PROJECT_ID

if [[ -z $projectId ]];then
  echo "set env variable PROJECT_ID"
  exit 1
fi

if [[ -z $version ]];then
  echo "provide container tag as the first arg"
  exit 1
fi

docker build -t sidecar-example-app:latest ./app
docker tag sidecar-example-app:latest "eu.gcr.io/$projectId/app-container:$version"
docker push "eu.gcr.io/$projectId/app-container:$version"
