GOOS=windows GOARCH=amd64 go build cmd/manifestgen.go
GOOS=linux GOARCH=amd64 go build cmd/manifestgen.go

mv manifestgen build/
mv manifestgen.exe build/

team="team-a"
./build/manifestgen --value="build/testing/values/$team.yaml" --template="build/testing/templates/team.yaml" --output="build/testing/output/$team.yaml" --filetype="yaml"
./build/manifestgen --value="build/testing/values/$team.json" --template="build/testing/templates/team.json" --output="build/testing/output/$team.json" --filetype="json"

team="team-b"
./build/manifestgen --value="build/testing/values/$team.yaml" --template="build/testing/templates/team.yaml" --output="build/testing/output/$team.yaml" --filetype="yaml"
