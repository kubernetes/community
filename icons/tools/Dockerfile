FROM ubuntu:18.04

RUN apt-get update && apt-get install -y librsvg2-bin

WORKDIR /convertor

ENTRYPOINT ["rsvg-convert"]
CMD ["-d", "300", "-p", "300"]
