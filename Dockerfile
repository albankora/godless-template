FROM golang:1.14-alpine

# Install base tools
RUN apk update \
    && apk upgrade \
    && apk add --no-cache wget bash git zip unzip

# Install terraform
RUN wget --quiet https://releases.hashicorp.com/terraform/0.12.24/terraform_0.12.24_linux_amd64.zip \
  && unzip terraform_0.12.24_linux_amd64.zip \
  && mv terraform /usr/bin \
  && rm terraform_0.12.24_linux_amd64.zip

# Env setup
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV GO111MODULE="on"
ENV CGO_ENABLED="0"

WORKDIR /app