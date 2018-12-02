FROM ubuntu:bionic

RUN apt-get update && apt-get install -y strace curl

CMD ["strace","-ff","curl","github.com"]