FROM golang:latest

ARG prjRoot=/go/src/github.com/whati001/acr-build-echo-golang/

COPY . $prjRoot

RUN go get -u github.com/labstack/echo/...
RUN cd $prjRoot && go build && go install

EXPOSE 2711
CMD ["acr-build-echo-golang"]
