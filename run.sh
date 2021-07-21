go build .
./manifestgen --value="testing/values/team.yaml" --template="testing/templates/team.yaml" --output="testing/output/team.yaml" --filetype="yaml"
./manifestgen --value="testing/values/team.json" --template="testing/templates/team.json" --output="testing/output/team.json" --filetype="json"