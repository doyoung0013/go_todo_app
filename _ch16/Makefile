.PHONY: help build build-local up down logs ps test 
.DEFAULT_GOAL := help
DOCKER_TAG := latest

build: ## Docker Hub에 배포하기 위한 프로덕션용 이미지를 빌드
	docker build -t gitwub5/gotodo:${DOCKER_TAG} \
		--target deploy ./

build-local: ## 로컬 개발 환경에서 사용할 이미지를 새로 빌드 (캐시 무시)
	docker compose build --no-cache

up: ## 도커 컴포즈로 모든 서비스를 백그라운드에서 실행
	docker compose up -d

down: ## 실행 중인 모든 도커 컴포즈 서비스를 중지하고 관련 리소스를 정리
	docker compose down

logs: ## 모든 컨테이너의 로그를 실시간으로 확인 (-f 옵션으로 실시간 추적)
	docker compose logs -f

ps: ## 현재 실행 중인 도커 컴포즈 서비스들의 상태
	docker compose ps

test: ## Go 언어 테스트를 실행:
      ## -race: 동시성 문제 감지
      ## -shuffle=on: 테스트 순서를 랜덤화하여 의존성 문제 체크
      ## ./...: 모든 하위 패키지 포함
	go test -race -shuffle=on ./...

help: ## 사용 가능한 모든 make 명령어와 각각의 설명
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'