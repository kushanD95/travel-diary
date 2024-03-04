FROM golang:1.22.0-bullseye as build
ARG GOLANG_VERSION

ENV GONOPROXY="gitub.com"
ENV GOPRIVATE="github.com"
ENV GOLANG_VERSION=${GOLANG_VERSION}

RUN apt-get clean
RUN apt-get update && apt-get install
RUN apt install -y make

WORKDIR /
ADD . /app

WORKDIR /app

RUN make build

FROM gcr.io/distroless/base-debian11
COPY --from=build /app/build /

ARG PGHOST
ARG PGPORT
ARG PGUSRNAME
ARG PGDB
ARG PGPWD

ENV PGHOST=${PGHOST}
ENV PGPORT=${PGPORT}
ENV PGDB=${PGDB}
ENV PGUSRNAME=${PGUSRNAME}
ENV PGPWD=${PGPWD}



ENTRYPOINT ["/diary-srv"]