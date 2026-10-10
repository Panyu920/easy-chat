db_url=mysql://root:easy-chat@tcp(localhost:13306)/easy_chat?multiStatements=true
sail_url=mysql://root:easy-chat@tcp(localhost:13306)/sail?multiStatements=true
migrate_path=db/migration
sail_path=db/migration/sail

.PHONY: create_db new_migration sqlc 
create_db:
	docker exec -it mysql mysql -u root -peasy-chat -e "CREATE DATABASE easy_chat;"
create_db_sail:
	docker exec -it mysql mysql -u root -peasy-chat -e "CREATE DATABASE sail;"	

new_migration:
	migrate create -ext sql -dir $(migrate_path) -seq $(name)
new_migration_sail:
	migrate create -ext sql -dir $(sail_path) -seq $(name)

migrate_up:
	migrate -path $(migrate_path) -database "$(db_url)" -verbose up

migrate_down:
	migrate -path $(migrate_path) -database "$(db_url)" -verbose down

migrate_down_1:
	migrate -path $(migrate_path) -database "$(db_url)" -verbose down 1

migrate_up_sail:
	migrate -path $(sail_path) -database "$(sail_url)" -verbose up	
migrate_down_sail:
	migrate -path $(sail_path) -database "$(sail_url)" -verbose down
sqlc:
	sqlc generate