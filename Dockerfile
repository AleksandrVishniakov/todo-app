FROM golang:1.21-alpine as build

WORKDIR /usr/local/go/src/todo-app/app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY app ./
RUN go build -o ./bin/app cmd/main.go


FROM scratch
WORKDIR /bin
COPY configs ./configs/
COPY .env ./
COPY --from=build /usr/local/go/src/todo-app/app/bin/app /bin/app
CMD ["app"]