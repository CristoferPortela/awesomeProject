FROM golang:1.19

WORKDIR /go/src/awe
COPY . .

RUN go mod download

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

CMD ["./wait-for-it.sh", "0.0.0.0:5432", "--timeout=2", "--", "go", "run", "."]

#ENTRYPOINT "tail"
#CMD ["tail", "-f", "/dev/null"]