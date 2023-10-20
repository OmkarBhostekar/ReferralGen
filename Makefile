postgres:
	docker run --name postgres16 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root referralgen
dropdb:
	docker exec -it postgres16 dropdb referralgen
migrateup:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/referralgen?sslmode=disable" -verbose up
migratedown:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/referralgen?sslmode=disable" -verbose down

# migrateup1:
# 	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up 1 
# migratedown1:
# 	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 1

# sqlc:
# 	sqlc generate
# test:
# 	go test -v -cover ./...
# mock:
# 	mockgen -destination db/mock/store.go example.com/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test mock

# to run sqlc on windows - docker run --rm -v "D:\Development\Go\Referral Generater Webapp:/src" -w /src sqlc/sqlc generate