
FROM postgres:alpine AS u-shylianok-task05-pg

COPY ./postgres_init.sql /docker-entrypoint-initdb.d/postgres_init.sql
EXPOSE 5432