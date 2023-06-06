FROM golang:1.20-alpine AS gobuild

WORKDIR /app
COPY . .

RUN go build cmd/main.go

RUN chmod +x main

FROM scratch
WORKDIR /app
COPY --from=gobuild /app .

EXPOSE 3000

ENTRYPOINT [ "./main" ]