#!/bin/sh

# Will delete itself on close
docker run --name postgres-db -e POSTGRES_PASSWORD=password -e POSTGRES_USER=pg -e POSTGRES_DB=pg -p 5432:5432 --rm postgres