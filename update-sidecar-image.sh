#!/bin/bash
version=$1
projectId=$PROJECT_ID

if [[ -z $projectId ]];then
  echo "Set env variable PROJECT_ID"
  exit 1
fi

docker build -t sidecar-example-sidecar:latest ./app
docker tag sidecar-example-sidecar:latest "eu.gcr.io/$projectId/sidecar-container:$version"
docker push "eu.gcr.io/$projectId/sidecar-container:$version"