FROM debian:latest
ADD ./.env /.env
ADD ./sysapi /sysapi
EXPOSE 8000
CMD ["/sysapi"]
