# Redeploy only the app service (faster for code changes)
ra:
	docker-compose stop app
	docker-compose rm -f app
	docker-compose build app
	docker-compose up -d app
