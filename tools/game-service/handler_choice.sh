#!/bin/bash

curl -v -s -X 'GET' http://localhost:8080/choice | jq .
