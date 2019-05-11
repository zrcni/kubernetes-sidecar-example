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

docker build -t sidecar-example-sidecar:latest ./sidecar
docker tag sidecar-example-sidecar:latest "eu.gcr.io/$projectId/sidecar-container:$version"
docker push "eu.gcr.io/$projectId/sidecar-container:$version"

echo "sidecar-container tag updated to $version. Update it in the deployment"
