#!/bin/bash

curl -v -s -X 'GET' \
  -H "Accept: application/json" \
  -H "Content-Type: application/json; charset=utf-8" \
  http://localhost:8080/choice | jq .
