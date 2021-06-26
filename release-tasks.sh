env

ls -la
ls -la bin
make build
ls -la
ls -la build
#docker run -v migrations:/migrations --network host migrate/migrate -path=/migrations/ -database ${DATABASE_URL} up