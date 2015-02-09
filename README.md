# Dockerized Go Webserver

This repository is meant as an example for a developer looking to create a minimal Docker image of a lightweight Go http fileserver.

Here are the interesting parts:

hello.go: `http.ListenAndServe(":8080", http.FileServer(http.Dir("/var/www/hello")))`

build.sh: `CGO_ENABLED=0 go build -o staticbinary -a -installsuffix cgo github.com/cfchris/dockerized-go-webserver`

Dockerfile:
```
FROM scratch

EXPOSE 8080

COPY wwwroot /var/www/hello
COPY staticbinary /bin/staticgoserver

ENTRYPOINT ["/bin/staticgoserver"]
```