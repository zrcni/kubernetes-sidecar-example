## kubernetes-sidecar-example

The docker-compose is only for local development.

```console
docker-compose build
docker tag kubernetes-sidecar-example_app:latest eu.gcr.io/<project-id>/app-container:<version>
docker tag kubernetes-sidecar-example_sidecar:latest eu.gcr.io/<project-id>/sidecar-container:<version>

docker push eu.gcr.io/<project-id>/app-container:<version>
docker push eu.gcr.io/<project-id>/sidecar-container:<version>

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### app
Container with a web server that serves a page with images in it. The images are served from a shared volume which is mounted at /images.

### sidecar container
Container with a program that downloads images and saves them to the shared volume mounted at /images. Images are updated every X seconds.
