## kubernetes-sidecar-example

Replace <PROJECT-ID> in deployment.yaml with Google Cloud project id

The docker-compose is only for local development.

```console
docker build -t sidecar-example-app:latest ./app
docker build -t sidecar-example-sidecar:latest ./sidecar

docker tag sidecar-example-app:latest eu.gcr.io/<project-id>/app-container:<version>
docker tag sidecar-example-sidecar:latest eu.gcr.io/<project-id>/sidecar-container:<version>

docker push eu.gcr.io/<project-id>/app-container:<version>
docker push eu.gcr.io/<project-id>/sidecar-container:<version>

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### app
Container with a web server that serves a page with images in it. The images are served from a shared volume which is mounted at /images.

### sidecar
Container with a program that downloads images and saves them to the shared volume mounted at /images. Images are updated every X seconds.
