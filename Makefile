run-app:
	docker compose up -d

stop-and-remove-app:
	docker compose down --rmi local