FROM alpine:edge

MAINTAINER Chuanjian Wang <chuanjian@staff.sina.com.cn>

RUN apk add --update wget && \
	cd /tmp && \
	wget https://github.com/prometheus/prometheus/releases/download/v2.0.0-beta.2/prometheus-2.0.0-beta.2.linux-amd64.tar.gz && \
	tar zxvf prometheus-2.0.0-beta.2.linux-amd64.tar.gz && \
	cp prometheus-2.0.0-beta.2.linux-amd64/prometheus /bin/prometheus && \
	rm -rf /tmp/*

COPY bundles/spond /bin/spond
COPY bundles/sponctl /bin/sponctl

EXPOSE 8080


