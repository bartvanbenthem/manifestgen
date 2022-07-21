# Manifestgen
Go module with package for generating yaml and json manifests trough Go templating.

#### Usage Manifestgen
```text
Usage of manifestgen:
  -read-from-file string
        path/file to variables file
  -template string
        path/file to template file
  -type string
        choose input type json / yaml (default "json")
  -write-to-file string
        path/file for writing generated file
```

#### Usage Serializer
```text
Usage of serializer:
  -escape string
        add \ as escape char (default "false")
  -json string
        stdin json input
  -jsonfile string
        path/file to json file
  -serialization string
        serialize / deserialize (default "serialize")
  -string string
        give json string to marshall into json object
```

#### Test & Build (Linux)
```bash
# clone repository
git clone https://github.com/bartvanbenthem/manifestgen.git
cd manifestgen

# test manifestgen module
go test -cover  ./...

# build manifestgen and serializer binaries
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/manifestgen
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/serializer

```
### ManifestGen Examples

```bash
# 01  ##########################################################
# testing manifest-printer JSON w/ pipe input | write to Stdout
cat build/testing/values/team-a.json | ./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json'

# 02  ##########################################################
# testing manifest-printer Yaml w/ pipe input  | write to Stdout
cat build/testing/values/team-a.yaml | ./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml'

# 03  ##########################################################
# testing manifest-writer JSON w/ pipe input | write to File
cat build/testing/values/team-b.json | ./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json' \
      --write-to-file='build/testing/output/team-b.json'

# 04 ##########################################################
# testing manifest-writer YAML w/ pipe input | write to File
cat build/testing/values/team-b.yaml | ./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --write-to-file='build/testing/output/team-b.yaml'

# 05  ##########################################################
# testing manifest-writer JSON w/ file reader input | write to Stdout
./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json' \
      --read-from-file='./build/testing/values/team-a.json'

# 06  ##########################################################
# testing manifest-printer JSON w/ file reader input | write to Stdout
./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/team.json' \
      --read-from-file='./build/testing/values/team-a.json'

# 07  ##########################################################
# testing manifest-printer YAML w/ file reader input | write to Stdout
./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --read-from-file='./build/testing/values/team-a.yaml'

# 08  ##########################################################
# testing manifest-writer JSON w/  pipe input | write to File
cat build/testing/values/team-a.json  | ./build/bin/manifestgen \
      --type='json'  \
      --template='build/testing/templates/team.json' \
      --write-to-file='build/testing/output/team-a.json'

# 09  ##########################################################
# testing manifest-writer YAML w/ file reader input | write to File
./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --read-from-file='./build/testing/values/team-a.yaml' \
      --write-to-file='build/testing/output/team-a.yaml'

# 10  ##########################################################
# testing manifest-writer YAML w/ file reader input | write to File
./build/bin/manifestgen \
      --type='yaml' \
      --template='build/testing/templates/team.yaml' \
      --read-from-file='./build/testing/values/team-a.yaml' \
      --write-to-file='build/testing/output/reader-to-writer.yaml'

# 11  ##########################################################
# terraform variable file example | write to File
cat 'build/testing/values/tf_variables.json' | ./build/bin/manifestgen \
      --type='json' \
      --template='build/testing/templates/test.tfvars.template' \
      --write-to-file='build/testing/output/test.tfvars'
```
### Serializer Examples

```bash

# serialization | deserialize | string-to-json
./build/bin/serializer \
    --serialization='deserialize' \
    --string="{\"project_name\":\"dss-test\",\"namespace_name\":\"dss-test\",\"namespace_description\":\"dss-test\"}" | jq .

# serialization | serialize | json-to-string | from json file input
./build/bin/serializer \
    --serialization='serialize' \
    --escape='true' \
    --jsonfile='./build/testing/values/tf_variables.json'

# serialization | serialize | json-to-string | from stdin no escape characters
STDINJSON=$(cat ./build/testing/values/tf_variables.json)
./build/bin/serializer \
    --serialization='serialize' \
    --escape='true' \
    --json="$STDINJSON"
```
