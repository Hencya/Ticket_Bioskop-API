# Builder
#FROM golang:1.17-alpine AS builder
#RUN mkdir /app
#ADD . /app
#WORKDIR /app
#RUN go mod tidy
#RUN go build -o main
#
## Runner
#FROM alpine:3.14
#WORKDIR /app
#COPY --from=builder /app/main .
#COPY --from=builder /app/.env.docker .
#EXPOSE 8000
#CMD ["./main"]

#==================================================================
# builder
FROM golang:1.17-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main

# Runner
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8000
CMD ["./main"]
