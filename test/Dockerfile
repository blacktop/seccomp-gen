FROM ubuntu:bionic
RUN apt-get update && apt-get install -y strace curl ca-certificates
COPY entrypoint.sh /entrypoint.sh
# ENTRYPOINT [ "/entrypoint.sh" ]
ENTRYPOINT [ "strace", "-c", "-f", "-S", "name" ]
CMD [ "curl", "https://github.com" ]