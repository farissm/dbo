############################
# STEP 1 build executable binary
############################
FROM golang:1.21.10-alpine3.18 AS builder
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o dbo .

############################
# STEP 2 build a small image
############################
FROM alpine
RUN mkdir /server
RUN apk update && apk add --no-cache tzdata && \
apk add p7zip
ENV TZ="Asia/Jakarta"
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/.env /server/.env
COPY --from=builder /app/dbo /server
WORKDIR /server 
EXPOSE 3013
CMD ["/server/dbo"]
