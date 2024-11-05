.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: # # 배 포 용 도 커 이 미 지 빌 드
	docker build -t junstory/go_todo_app/section57:${DOCKER_TAG} --target deploy ./

build-local: # \ 로 컬 환 경 용 도 커 이 미 지 빌 드
	docker compose build --no-cache

up: # # 자 동 새 로 고 침 을 사 용 한 도 커 구 성 기 실 행
	docker compose up -d
down: # # 도 커 구 성 기 종 료
	docker compose down

logs: # # 도 커 구 성 기 로 그 출 력
	docker compose logs -f
ps: # 컨 테 이 너 상 태 확 인
	docker compose ps

test: # # 테 스 트 실 행
	go test -race -shuffle=on ./...

help: # 옵 션 보 기
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n",
$$1, $$2}'