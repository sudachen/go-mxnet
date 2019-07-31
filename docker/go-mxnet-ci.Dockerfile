FROM sudachen/go1127:latest
LABEL maintainer="Alexey Sudachen <alexey@sudachen.name>"

USER root
RUN set -ex \
&& apt-get update --fix-missing \
 && apt-get install -qy --no-install-recommends \
    software-properties-common \
    mercurial \
    xvfb \
    parallel \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && curl -L --retry 3 --output /usr/bin/jq https://circle-downloads.s3.amazonaws.com/circleci-images/cache/linux-amd64/jq-latest \
 && chmod +x /usr/bin/jq \
 && jq --version \
 && export DOCKER_VERSION=$(curl --silent --fail --retry 3 https://download.docker.com/linux/static/stable/x86_64/ | grep -o -e 'docker-[.0-9]*\.tgz' | sort -r | head -n 1) \
 && DOCKER_URL="https://download.docker.com/linux/static/stable/x86_64/${DOCKER_VERSION}" \
 && echo Docker URL: $DOCKER_URL \
 && curl --silent --show-error --location --fail --retry 3 "${DOCKER_URL}" | tar -xz -C /tmp \
 && mv /tmp/docker/* /usr/bin \
 && rm -rf /tmp/docker  \
 && which docker \
 && (docker version || true) \
 && COMPOSE_URL="https://circle-downloads.s3.amazonaws.com/circleci-images/cache/linux-amd64/docker-compose-latest" \
 && curl --silent --show-error --location --fail --retry 3 --output /usr/bin/docker-compose $COMPOSE_URL \
 && chmod +x /usr/bin/docker-compose \
 && docker-compose version \
 && DOCKERIZE_URL="https://circle-downloads.s3.amazonaws.com/circleci-images/cache/linux-amd64/dockerize-latest.tar.gz" \
 && curl --silent --show-error --location --fail --retry 3 $DOCKERIZE_URL | tar -C /usr/local/bin -xzv  \
 && dockerize --version \
 && curl -L https://github.com/sudachen/mxnet/releases/download/1.5.0-mkldnn-static/libmxnet_cpu.7z -o /tmp/mxnet.7z \
 && 7z x /tmp/mxnet.7z -o/ \
 && rm /tmp/mxnet.7z \
 && chmod +r -R /opt/mxnet \
 && chmod +x $(find /opt/mxnet -type d) \
 && chmod +x $(find /opt/mxnet/lib -type f) \
 && ln -sf libmxnet_cpu.so /opt/mxnet/lib/libmxnet.so

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | INSTALL_DIRECTORY=/usr/local/bin sh
RUN curl -sSL https://github.com/gotestyourself/gotestsum/releases/download/v0.3.4/gotestsum_0.3.4_linux_amd64.tar.gz | \
  tar -xz -C /usr/local/bin gotestsum

RUN curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > /usr/local/bin/cc-test-reporter && \
    chmod +x /usr/local/bin/cc-test-reporter

USER $USER
CMD ["/bin/sh"]
