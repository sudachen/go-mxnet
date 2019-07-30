FROM  sudachen/mxnet:latest
LABEL maintainer="Alexey Sudachen <alexey@sudachen.name>"
ENV GOLANG_VERSION 1.12.7
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

USER root

RUN set -ex \
 && apt-get update --fix-missing \
 && apt-get install -qy --no-install-recommends \
    gcc \
    g++ \
    m4 \
	libc6-dev \
    pkg-config \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && curl -L https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz | tar zx -C /usr/local \
 && go version \
 && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
 && chmod -R 777 "$GOPATH"

USER $USER
WORKDIR $GOPATH
