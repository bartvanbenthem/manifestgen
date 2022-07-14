#!/bin/bash

# build manifestgen and serializer binaries
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/manifest-writer
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/serializer

# team-a kubernetes manifest example
team="team-a"
./build/bin/manifest-writer \
      --filetype="yaml" \
      --value="build/testing/values/$team.yaml" \
      --template="build/testing/templates/team.yaml" \
      --output="build/testing/output/$team.yaml"

./build/bin/manifest-writer \
      --value="build/testing/values/$team.json" \
      --template="build/testing/templates/team.json" \
      --output="build/testing/output/$team.json"

# team-b kubernetes manifest example
team="team-b"
./build/bin/manifest-writer \
      --filetype="yaml" \
      --value="build/testing/values/$team.yaml" \
      --template="build/testing/templates/team.yaml" \
      --output="build/testing/output/$team.yaml"

./build/bin/manifest-writer \
      --value="build/testing/values/$team.json" \
      --template="build/testing/templates/team.json" \
      --output="build/testing/output/$team.json"

# terraform variable file example
./build/bin/manifest-writer \
      --value="build/testing/values/tf_variables.json" \
      --template="build/testing/templates/test.tfvars.template" \
      --output="build/testing/output/test.tfvars"


printf "\n"
# serialization | deserialize | string-to-json
./build/bin/serializer \
    --serialization='deserialize' \
    --string="{\"project_name\":\"dss-test\",\"namespace_name\":\"dss-test\",\"namespace_description\":\"dss-test\"}" | jq .

printf "\n"
# serialization | serialize | json-to-string | from json file input
./build/bin/serializer \
    --serialization='serialize' \
    --escape='true' \
    --jsonfile='./build/testing/values/tf_variables.json'

printf "\n"
# serialization | serialize | json-to-string | from stdin no escape characters
STDINJSON=$(cat ./build/testing/values/tf_variables.json)
./build/bin/serializer \
    --serialization='serialize' \
    --escape='true' \
    --json="$STDINJSON"