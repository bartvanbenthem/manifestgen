# Manifestgen
Go module with package for generating yaml and json manifests trough Go templating.

#### Usage Manifestgen
```bash
Usage of manifestgen:
  -filetype string
        json / yaml (default "json")
  -output string
        path/file to output file (default "./output.json")
  -template string
        path/file to template file (default "")
  -value string
        path/file to values file (default "")
```

#### Usage Serializer
```bash
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

#### Install and run (Linux)
```bash
# clone repository
git clone https://github.com/bartvanbenthem/manifestgen.git
cd manifestgen

# build manifestgen and serializer binaries
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/manifestgen
GOOS=linux GOARCH=amd64 go build -o build/bin ./cmd/serializer

# team-a kubernetes manifest example
team="team-a"
./build/bin/manifestgen \
      --filetype="yaml" \
      --value="build/testing/values/$team.yaml" \
      --template="build/testing/templates/team.yaml" \
      --output="build/testing/output/$team.yaml"

./build/bin/manifestgen \
      --value="build/testing/values/$team.json" \
      --template="build/testing/templates/team.json" \
      --output="build/testing/output/$team.json"

# team-b kubernetes manifest example
team="team-b"
./build/bin/manifestgen \
      --filetype="yaml" \
      --value="build/testing/values/$team.yaml" \
      --template="build/testing/templates/team.yaml" \
      --output="build/testing/output/$team.yaml"

./build/bin/manifestgen \
      --value="build/testing/values/$team.json" \
      --template="build/testing/templates/team.json" \
      --output="build/testing/output/$team.json"

# terraform variable file example
./build/bin/manifestgen \
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
```
