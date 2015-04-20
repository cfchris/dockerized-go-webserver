# Dockerized Go Webserver

This repository is meant as an example for a developer looking to create a minimal Docker image of a lightweight Go http fileserver.

Here's the TLDR:

server.go: (Serve files located at "/var/www/hello".)
```
http.ListenAndServe(":8080", http.FileServer(http.Dir("/var/www/hello")))
```

build.sh: (Compile a non dynamically linked binary.)
```
CGO_ENABLED=0 go build -o staticbinary -a -installsuffix cgo github.com/cfchris/dockerized-go-webserver
```

Dockerfile: (Have Docker copy the static files into the container at "/var/www/hello".)
```
FROM scratch

EXPOSE 8080

COPY wwwroot /var/www/hello
COPY staticbinary /bin/staticgoserver

ENTRYPOINT ["/bin/staticgoserver"]
```