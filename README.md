# PDF compressor in Golang

### Build Docker compose
``` shell
docker-compose up -d --build
```

### Stop Docker Compose

``` shell
docker-compose down      
```

### Deloy to Docker hub
```shell
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t thebinij/pdf-compressor:latest --push .

```

### DockerHUB Link

[Docker Hub Repository](https://hub.docker.com/r/thebinij/pdf-compressor)
