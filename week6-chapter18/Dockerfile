# 배 포 용 컨 테 이 너 에 포 함 시 킬 바 이 너 리 를 생 성 하 는 컨 테 이 너
FROM golang:1.23.1 as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
#binary build
RUN go build -trimpath -ldflags "-w -s" -o app

# 배 포 용 컨 테 이 너
FROM ubuntu:latest as deploy

RUN apt-get update && apt-get install -y libgcc-s1

COPY --from=deploy-builder /app/app .
CMD ["./app"]

# 로 컬 개 발 환 경 에 서 사 용 하 는 자 동 새 로 고 침 환 경
FROM golang:1.23.1 as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]