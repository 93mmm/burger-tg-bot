# Установка зависимостей
deps: .install-deps

# Генерация
gen: .install-deps .buflint .bufgen .mockgen

# Линт проекта
lint-all: .lint-changes .lint-full .buflint

# Запуск docker
up: .docker-up


# Установка зависимостей
.install-deps:
	go install tool



# Линт протофайлов
.buflint:
	@buf lint

# Генерация протофайлов
.bufgen:
	@buf generate

.mockgen:
	@go generate -run minimock ./...



# Проверяем не сделали ли мы обратно-несовместимые изменения контракта по сравнению с develop веткой для разработки
.lint-changes:
	golangci-lint run \
		--new-from-rev=origin/main \
		--config=.golangci.yml \
		./...


# Линт всех файлов
.lint-full:
	golangci-lint run \
		--config=.golangci.yml \
		./...


# Докер
.docker-up:
	@cd dev && docker compose up
