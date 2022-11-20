FROM golang:alpine
ENV CGO_ENABLED 0
ARG VCS_REF

COPY . /the-odd-api

WORKDIR /the-odd-api

RUN go build -ldflags "-X main.build=${VCS_REF}"

CMD [ "./the-odd-api" ]