FROM ubuntu:bionic
RUN apt-get update && apt-get install -y strace curl ca-certificates
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT [ "/entrypoint.sh" ]