FROM golang:alpine as builder

RUN apk add --update curl \
 	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ENV APP_PATH $GOPATH/src/$REPO_BASE_URL/$APP_NAME/backend

COPY . $APP_PATH
WORKDIR $APP_PATH

RUN apk add --update git && \
	dep ensure

RUN go build -o /opt/game-server ./src

FROM golang:alpine as main

COPY --from=builder /opt/game-server /opt/game-server

CMD ["sh", "-c", "/opt/game-server"]
