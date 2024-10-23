DATABASE_URL := "postgres://root:root@127.0.0.1:5432/pos?sslmode=disable"

up:
	migrate -path ./database/migrations -database $(DATABASE_URL) up

down:
	@read -p "Enter the number of migrations to roll back (default is 1): " NUM; \
	if [ -z "$$NUM" ]; then NUM=1; fi; \
	migrate -path ./database/migrations -database $(DATABASE_URL) down $$NUM

# Add a new target that accepts an ID as an argument
force:
	@read -p "Enter the migration ID: " ID; \
	migrate -path ./database/migrations -database $(DATABASE_URL) force $$ID

version:
	migrate -path ./database/migrations -database $(DATABASE_URL) version

drop:
	migrate -path ./database/migrations -database $(DATABASE_URL) drop -f
