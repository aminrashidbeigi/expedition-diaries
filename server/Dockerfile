FROM docker.repos.balad.ir/golang:1.18 AS BUILDER

WORKDIR /server
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN make build

FROM docker.repos.balad.ir/debian:bullseye-slim
WORKDIR /server
COPY --from=BUILDER /server/history-travels /usr/bin/history-travels
COPY --from=BUILDER /server .

EXPOSE 8080

CMD [ "history-travels"]