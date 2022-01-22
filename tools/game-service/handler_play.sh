#!/bin/bash

curl -v -s -X 'POST' \
  -d '{"player": 1}' \
  -H "Accept: application/json" \
  -H "Content-Type: application/json; charset=utf-8" \
  http://localhost:8080/play | jq .
