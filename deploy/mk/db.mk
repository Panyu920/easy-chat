db_url=mysql://root:easy-chat@tcp(localhost:13306)/easy_chat?multiStatements=true
migrate_path=db/migration

.PHONY: create_db new_migration sqlc 
create_db:
	docker exec -it mysql mysql -u root -peasy-chat -e "CREATE DATABASE easy_chat;"

new_migration:
	migrate create -ext sql -dir $(migrate_path) -seq $(name)

migrate_up:
	migrate -path $(migrate_path) -database "$(db_url)" -verbose up

migrate_down:
	migrate -path $(migrate_path) -database "$(db_url)" -verbose down

sqlc:
	sqlc generate