# Manifestgen
Go Module with package for generating yaml and json manifests trough Go templating.

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

#### Install and run (Linux)
```bash
git clone https://github.com/bartvanbenthem/manifestgen.git
cd manifestgen
sudo cp build/manifestgen /usr/bin
manifestgen --help

# Example
team="team-a"
manifestgen --value="build/testing/values/$team.yaml" --template="build/testing/templates/team.yaml" --output="build/testing/output/$team.yaml" --filetype="yaml"
manifestgen --value="build/testing/values/$team.json" --template="build/testing/templates/team.json" --output="build/testing/output/$team.json" --filetype="json"
```

#### Install and run (Windows)
```powershell
git clone https://github.com/bartvanbenthem/manifestgen.git
cd manifestgen

# Example
$team="team-a"
.\build\manifestgen.exe --value="build\testing\values\$team.yaml" --template="build\testing\templates\team.yaml" --output="build\testing\output\$team.yaml" --filetype="yaml"
.\build\manifestgen.exe --value="build\testing\values\$team.json" --template="build\testing\templates\team.json" --output="build\testing\output\$team.json" --filetype="json"
```
