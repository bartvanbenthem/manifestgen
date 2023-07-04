# manifestgen
CLI tool for generating configuration manifests. Both JSON and YAML are valid input types that can be converted from template to manifest.

# usage manual
```bash
Template Renderer 

USAGE:
    manifestgen [OPTIONS] --template <template_file>

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -o, --output <output_file>          Path to the output file
    -t, --template <template_file>      Path to the template file
    -v, --variables <variables_file>    Path to the JSON file

```

## examples
```bash
# read variables from a file and write manifest to a file
./target/debug/manifestgen -v project/example-var.json -t project/targets.tmpl -o project/config.yaml

# read variables from stdin and write manifest to a file
cat project/example-var.json | ./target/debug/manifestgen -t project/targets.tmpl -o project/config.yaml

# read JSON variables from a file and write manifest to stdout
./target/debug/manifestgen -v project/example-var.json -t project/targets.tmpl

# read YAML variables from a file and write manifest to stdout
./target/debug/manifestgen -v project/example-var.yaml -t project/targets.tmpl

# read JSON variables from stdin and write manifest to stdout
cat project/example-var.json | ./target/debug/manifestgen --template project/targets.tmpl

# read YAML variables from stdin and write manifest to stdout
cat project/example-var.yaml | ./target/debug/manifestgen --template project/targets.tmpl

```
