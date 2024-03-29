FROM golang:1.18 as build

WORKDIR /go/src/app
COPY  go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]
