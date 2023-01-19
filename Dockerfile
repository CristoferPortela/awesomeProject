FROM golang:1.19

WORKDIR /go/src/awe
ADD . /go/src/awe
WORKDIR /go/src/awe


RUN go mod download
RUN go mod tidy
RUN #go build -buildmode=plugin

RUN for file in ./addon/*; \
    do \
      if [ -f "$file/main.go" ]; \
        then \
          go build -buildmode=plugin -o $file/main.so $file/main.go; \
      else \
        echo $file; \
      fi \
    done

RUN go build

RUN go run .

ENTRYPOINT ["tail"]
CMD ["-f","/dev/null"]