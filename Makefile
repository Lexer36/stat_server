migrate:
	migrate -path ./migrations -database 'postgres://postgres:1@0.0.0.0:5432/stat_srv?sslmode=disable' up
