# Manifestgen

* Generate yaml and json manifests trough Go templating 
* Variable json and yaml input to default templated output
* Convert quickly between yaml and json file types

#### Usage
```bash
Usage of manifestgen:
  -filetype string
        yaml / json (default "yaml")
  -output string
        path/file to output file (default "./output.yaml")
  -template string
        path/file to template file (default "./template.yaml")
  -value string
        path/file to values file (default "./value.yaml")
```

#### Examples 
* Linux bash
```bash
team="team-a"
./build/manifestgen --value="build/testing/values/$team.yaml" --template="build/testing/templates/team.yaml" --output="build/testing/output/$team.yaml" --filetype="yaml"
./build/manifestgen --value="build/testing/values/$team.json" --template="build/testing/templates/team.json" --output="build/testing/output/$team.json" --filetype="json"
```

* Windows PowerShell
```powershell
$team="team-a"
.\build\manifestgen.exe --value="build\testing\values\$team.yaml" --template="build\testing\templates\team.yaml" --output="build\testing\output\$team.yaml" --filetype="yaml"
.\build\manifestgen.exe --value="build\testing\values\$team.json" --template="build\testing\templates\team.json" --output="build\testing\output\$team.json" --filetype="json"
```