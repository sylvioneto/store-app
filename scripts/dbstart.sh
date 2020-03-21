docker run --name postgres-store \
-e POSTGRES_PASSWORD=mysecretpassword \
-e POSTGRES_DB=storedb \
-p 5432:54320 \
-d postgres 