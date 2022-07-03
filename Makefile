.PHONY: up
up:
	@echo "Running docker-compose up..."
	docker-compose up -d
	@echo "Done!"

.PHONY: down
down:
	@echo "Running docker-compose down..."
	docker-compose down
	@echo "Done!"