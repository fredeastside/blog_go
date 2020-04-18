.PHONY: build

build:
	@cd frontend && npm run-script build
	@echo "[✔️] Frontend build complete!"

certbot-test:
	@chmod +x ./docker/frontend/register_ssl.sh
	@sudo ./docker/frontend/register_ssl.sh \
								--domains "$(DOMAINS)" \
								--email $(EMAIL) \
								--data-path ./docker/certbot \
								--staging 1

certbot-prod:
	@chmod +x ./docker/frontend/register_ssl.sh
	@sudo ./docker/frontend/register_ssl.sh \
								--domains "$(DOMAINS)" \
								--email $(EMAIL) \
								--data-path ./docker/certbot \
								--staging 0

run-db:
	@docker-compose \
					-f docker-compose.yml \
					-f docker-compose.dev.yml \
					up -d --build

deploy-prod:
	@docker-compose \
					-f docker-compose.yml \
					-f docker-compose.prod.yml \
					up -d --build --force-recreate