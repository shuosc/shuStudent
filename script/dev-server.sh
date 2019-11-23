#!/usr/bin/env bash

export DB_ADDRESS="postgres://localhost:5432/postgres?sslmode=disable"
export PORT="8001"
gin -p 8000 run main.go