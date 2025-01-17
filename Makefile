include .env

run:
	go run cmd/web/main.go

start-nodemon:
	nodemon --exec go run cmd/web/main.go --signal SIGTERM

test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out && rm -f coverage.out

create-mock:
	mockery --all --case underscore --output ./mocks

migratecreate:
	@migrate create -ext sql -dir ${CURDIR}/db/migrations/ -seq ${name}

migrateforce:
	@migrate -path ${CURDIR}/db/migrations/ \
		-database "mysql://${MYSQL_READ_USER}:${MYSQL_READ_PASS}@tcp(${MYSQL_READ_HOST}:${MYSQL_READ_PORT})/${MYSQL_READ_DBNAME}" \
		-verbose force 1

migratedown:
	@migrate -path ${CURDIR}/db/migrations/ \
		-database "mysql://${MYSQL_READ_USER}:${MYSQL_READ_PASS}@tcp(${MYSQL_READ_HOST}:${MYSQL_READ_PORT})/${MYSQL_READ_DBNAME}" \
		-verbose down

migrateup:
	@migrate -path ${CURDIR}/db/migrations/ \
		-database "mysql://${MYSQL_READ_USER}:${MYSQL_READ_PASS}@tcp(${MYSQL_READ_HOST}:${MYSQL_READ_PORT})/${MYSQL_READ_DBNAME}" \
		-verbose up

SEED_SCRIPT = ${CURDIR}/db/seeds/seed.sh
SEEDS_SQL = ${CURDIR}/db/seeds/seeds.sql

chmod+x:
	@chmod +x $(SEED_SCRIPT)

seed: chmod+x
	@$(SEED_SCRIPT) ${MYSQL_READ_USER} ${MYSQL_READ_PASS} ${MYSQL_READ_HOST} ${MYSQL_READ_PORT} ${MYSQL_READ_DBNAME} ${"disable"}
