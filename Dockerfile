FROM golang:1.22.5 AS builder

WORKDIR  /app

COPY go.mod .

RUN go mod download

COPY . . 

RUN go build -o main.exe .


FROM gcr.io/distroless/base

COPY --from=builder /app/main.exe .

COPY --from=builder /app/landing_page ./landing_page

EXPOSE 8080
CMD [ "./main.exe" ]



