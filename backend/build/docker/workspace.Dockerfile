FROM golang:1.20

WORKDIR /code

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y bash-completion jq ca-certificates && \
    # cd /usr/local/bin && \
    # curl -s -L "https://github.com/cespare/reflex/releases/download/v0.3.1/reflex_darwin_amd64.tar.gz" | tar xvz && \
	# mv reflex_darwin_amd64/reflex /usr/local/bin/reflex && \
    # Makefile completion
    apt-get install -y bash-completion && \
    echo ". /etc/bash_completion" >> /root/.bashrc && \
    # Golint
    go install golang.org/x/lint/golint@latest

ENV GOCACHE="/go/cache"
ENV GOBIN="/go/bin"
ENV PATH="/go/bin:$PATH"