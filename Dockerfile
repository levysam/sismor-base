FROM clearlinux/golang

RUN swupd bundle-add c-extras-gcc10

RUN go install github.com/go-jet/jet/v2/cmd/jet@latest

RUN go install github.com/cosmtrek/air@latest

WORKDIR /src/sismor-base

# RUN go get

RUN air init

CMD ["air"]