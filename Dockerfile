FROM alpine:3.4

ADD *.go /poc-pg-mirror/

RUN apk add --update bash \
  && apk --update add git \
  && apk --update add go \
  && export GOPATH=/gopath \
  && REPO_PATH="github.com/utilitywarehouse/poc-pg-mirror" \
  && mkdir -p $GOPATH/src/${REPO_PATH} \
  && mv poc-pg-mirror/* $GOPATH/src/${REPO_PATH} \
  && rm -rf poc-pg-mirror \
  && cd $GOPATH/src/${REPO_PATH} \
  && go get -t ./... \
  && go build \
  && mv poc-pg-mirror /poc-pg-mirror \
  && apk del go git \
  && rm -rf $GOPATH /var/cache/apk/*

CMD [ "/poc-pg-mirror" ]
