#!/bin/bash

curl -v -s -X 'GET' \
  -H "accept: application/json" \
  -H "Content-Type: application/json" \
  http://localhost:8080/choices | jq .
