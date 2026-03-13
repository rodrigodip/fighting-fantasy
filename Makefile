# ==================================================================================
# FFANTASY - Makefile
# ==================================================================================

.PHONY: help dev dev-build dev-down dev-logs prod prod-build prod-down prod-logs clean

# Exibe os comandos disponíveis
help:
	@echo ""
	@echo "Comandos disponíveis:"
	@echo ""
	@echo "  DEV"
	@echo "    make dev            Sobe o ambiente de desenvolvimento"
	@echo "    make dev-build      Sobe o ambiente de desenvolvimento forçando rebuild"
	@echo "    make dev-down       Derruba o ambiente de desenvolvimento"
	@echo "    make dev-logs       Exibe os logs do app em desenvolvimento"
	@echo ""
	@echo "  PROD"
	@echo "    make prod           Sobe o ambiente de produção"
	@echo "    make prod-build     Sobe o ambiente de produção forçando rebuild"
	@echo "    make prod-down      Derruba o ambiente de produção"
	@echo "    make prod-logs      Exibe os logs do app em produção"
	@echo ""
	@echo "  UTILS"
	@echo "    make clean          Remove containers, volumes e imagens órfãs"
	@echo ""

# ==================================================================================
# DEV
# ==================================================================================

dev:
	docker compose up -d

dev-build:
	docker compose up --build -d

dev-down:
	docker compose down

dev-logs:
	docker compose logs -f app

dev-db:
	docker exec -it mongo mongosh -u admin -p yourpassword --authenticationDatabase admin

# ==================================================================================
# PROD
# ==================================================================================

prod:
	docker compose -f docker-compose.prod.yml up -d

prod-build:
	docker compose -f docker-compose.prod.yml up --build -d

prod-down:
	docker compose -f docker-compose.prod.yml down

prod-logs:
	docker compose -f docker-compose.prod.yml logs -f app

# ==================================================================================
# UTILS
# ==================================================================================

clean:
	docker compose down -v --remove-orphans
	docker image prune -f
