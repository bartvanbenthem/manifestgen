#!/bin/bash

# build manifestgen binaries
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/manifestgen
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/string-to-json
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/json-to-string

# team-a kubernetes manifest example
team="team-a"
./build/bin/manifestgen --filetype="yaml" --value="build/testing/values/$team.yaml" --template="build/testing/templates/team.yaml" --output="build/testing/output/$team.yaml"
./build/bin/manifestgen --value="build/testing/values/$team.json" --template="build/testing/templates/team.json" --output="build/testing/output/$team.json"

# team-b kubernetes manifest example
team="team-b"
./build/bin/manifestgen --filetype="yaml" --value="build/testing/values/$team.yaml" --template="build/testing/templates/team.yaml" --output="build/testing/output/$team.yaml"
./build/bin/manifestgen --value="build/testing/values/$team.json" --template="build/testing/templates/team.json" --output="build/testing/output/$team.json"

# terraform variable file example
./build/bin/manifestgen --value="build/testing/values/tf_variables.json" --template="build/testing/templates/test.tfvars.template" --output="build/testing/output/test.tfvars"

printf "\n"
# test json-to-string
./build/bin/json-to-string --escape="true" --jsonfile="./build/testing/values/tf_variables.json"

printf "\n"
# test string-to-json
./build/bin/string-to-json --string="{\"project_name\":\"dss-test\",\"namespace_name\":\"dss-test\",\"namespace_description\":\"dss-test\"}"