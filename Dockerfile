FROM golang


COPY ./project /go/app/
WORKDIR /go/app/
RUN go mod init trppproject 
RUN go mod tidy
CMD ["go", "run", "."]