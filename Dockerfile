FROM golang:1.23.1
WORKDIR /test
COPY . /test
RUN go build /test
# EXPOSE 3000
ENTRYPOINT [ "./go-url-shortener" ]
