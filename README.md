## kubernetes-sidecar-example

[http://sidecar.samulir.site](http://sidecar.samulir.site)

Replace \<PROJECT-ID\> with Google Cloud project id in deployment.yaml

The docker-compose is only for local development.

```console
./update-app-image <tag>
./update-sidecar-image <tag>

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### app
Container with a web server that serves a page with images in it. The images are served from a shared volume which is mounted at /images.

### sidecar
Container with a program that downloads images and saves them to the shared volume mounted at /images. Images are updated every X seconds.
