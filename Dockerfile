#FROM scratch

#ADD bin/helloworld-linux64-static /helloworld

#ENTRYPOINT ["/helloworld"]

FROM php:5.6-apache

COPY src/ /var/www/html/