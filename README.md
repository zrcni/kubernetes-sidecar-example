## kubernetes-sidecar-example

[http://sidecar.samulir.site](http://sidecar.samulir.site)

The docker-compose is only for local development.

### app
Web server that serves a page with images. The images are served from a shared volume which is mounted at /images.

### sidecar
Downloads images and saves them to the shared volume mounted at /images. Images are updated once a minute.
