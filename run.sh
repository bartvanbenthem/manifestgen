#!/bin/bash

# test manifestgen module
go test -cover  ./...

# build manifestgen and serializer binaries
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/manifestgen
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/serializer

printf "\n"
# 01  ##########################################################
# testing manifest-printer JSON w/ pipe input | write to Stdout
cat build/testing/values/team-a.json | ./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json'

printf "\n"
# 02  ##########################################################
# testing manifest-printer Yaml w/ pipe input  | write to Stdout
cat build/testing/values/team-a.yaml | ./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml'

printf "\n"
# 03  ##########################################################
# testing manifest-writer JSON w/ pipe input | write to File
cat build/testing/values/team-b.json | ./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json' \
      --write-to-file='build/testing/output/team-b.json'

printf "\n"
# 04 ##########################################################
# testing manifest-writer YAML w/ pipe input | write to File
cat build/testing/values/team-b.yaml | ./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --write-to-file='build/testing/output/team-b.yaml'

printf "\n"
# 05  ##########################################################
# testing manifest-writer JSON w/ file reader input | write to Stdout
./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json' \
      --read-from-file='./build/testing/values/team-a.json'

printf "\n"
# 06  ##########################################################
# testing manifest-printer JSON w/ file reader input | write to Stdout
./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json' \
      --read-from-file='./build/testing/values/team-a.json'

printf "\n"
# 07  ##########################################################
# testing manifest-printer YAML w/ file reader input | write to Stdout
./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --read-from-file='./build/testing/values/team-a.yaml'


printf "\n"
# 08  ##########################################################
# testing manifest-writer JSON w/  pipe input | write to File
cat build/testing/values/team-a.json  | ./build/bin/manifestgen \
      --type='json'  \
      --template='build/testing/templates/team.json' \
      --write-to-file='build/testing/output/team-a.json'

printf "\n"
# 09  ##########################################################
# testing manifest-writer YAML w/ file reader input | write to File
./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --read-from-file='./build/testing/values/team-a.yaml' \
      --write-to-file='build/testing/output/team-a.yaml'

printf "\n"
# 10  ##########################################################
# testing manifest-writer YAML w/ file reader input | write to File
./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --read-from-file='./build/testing/values/team-a.yaml' \
      --write-to-file='build/testing/output/reader-to-writer.yaml'

printf "\n"
# 11  ##########################################################
# terraform variable file example | write to File
cat 'build/testing/values/tf_variables.json' | ./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/test.tfvars.template' \
      --write-to-file='build/testing/output/test.tfvars'



printf "##################################################################\n"

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