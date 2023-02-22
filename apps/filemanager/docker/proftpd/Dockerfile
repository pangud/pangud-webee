FROM debian:stable-slim
RUN apt-get update \
    && apt-get install -y \
    proftpd \
    proftpd-mod-mysql \
    proftpd-mod-crypto \
    && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/usr/sbin/proftpd", "--nodaemon"]