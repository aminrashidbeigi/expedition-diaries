FROM golang:1.18 AS BUILDER

WORKDIR /server
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN make build

FROM debian:bullseye-slim
WORKDIR /server
COPY --from=BUILDER /server/expedition-diaries /usr/bin/expedition-diaries
COPY --from=BUILDER /server .

EXPOSE 8080

CMD [ "expedition-diaries", "-config=./local-config.json"]