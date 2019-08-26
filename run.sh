#!/bin/sh

EXERCISE=$1
OUTPUT_PATH=sample go run "$1/main.go" && cat sample && rm sample
