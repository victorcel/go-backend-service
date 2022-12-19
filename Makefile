ifneq (,$(findstring xterm,${TERM}))
	BLACK        := $(shell tput -Txterm setaf 0)
	RED          := $(shell tput -Txterm setaf 1)
	GREEN        := $(shell tput -Txterm setaf 2)
	YELLOW       := $(shell tput -Txterm setaf 3)
	LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
	PURPLE       := $(shell tput -Txterm setaf 5)
	BLUE         := $(shell tput -Txterm setaf 6)
	WHITE        := $(shell tput -Txterm setaf 7)
	RESET := $(shell tput -Txterm sgr0)
else
	BLACK        := ""
	RED          := ""
	GREEN        := ""
	YELLOW       := ""
	LIGHTPURPLE  := ""
	PURPLE       := ""
	BLUE         := ""
	WHITE        := ""
	RESET        := ""
endif

help:
	@echo "${PURPLE}"
	@echo ------------------------------------------------------------
	@echo .: Amaris Consulting:. - HELP - SHOW ALL COMMANDS
	@echo ------------------------------------------------------------
	@echo ""${GREEN}""

	@echo make build -- Checked
	@echo make run   -- Start
	@echo make test  -- testing coverage
	@echo make up    -- docker up
	@echo make down  -- docker down
#go tool commands
build:
	@cp .env.example .env
	@go build ./...

run:
	@go run main.go

## tests
test:
	@go test ./... --cover
## docker compose
up:
	docker-compose up -d --build
down:
	docker-compose down --remove-orphans