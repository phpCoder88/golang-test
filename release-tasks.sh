go env
#docker run -v migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${DATABASE_URL} up