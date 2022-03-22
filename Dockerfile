FROM golang:alpine AS build
# add gcc tools
RUN apk add build-base
# run build process in /app directory
WORKDIR /app
# copy dependencies and get them
COPY ./go.mod ./
COPY ./go.sum ./
# copy go src file(s)
COPY ./main.go ./main.go
COPY ./counter ./counter

# build the binary
RUN GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -o echo ./main.go

FROM alpine
# copy the binary from build stage
COPY --from=build /app/echo /bin/echo
# use non root
USER 1000:1000
# start server
CMD ["./bin/echo"]