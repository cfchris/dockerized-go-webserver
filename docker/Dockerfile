FROM scratch
MAINTAINER Chris Phillips <cfchris@gmail.com>

EXPOSE 8080

COPY wwwroot /var/www/hello
COPY staticbinary /bin/staticgoserver

ENTRYPOINT ["/bin/staticgoserver"]
