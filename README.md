# manifestgen
CLI tool for generating configuration manifests, effective for integration with configuration management pipelines to reduce the ammount of handwritten specifications. Both JSON and YAML are valid input types that can be converted from template to manifest.

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

## build

```bash
rustup target add x86_64-unknown-linux-musl

cargo test

cargo build
cargo build --target x86_64-unknown-linux-musl --release
```

## examples
```bash
manifestgen='./target/x86_64-unknown-linux-musl/release/manifestgen'

# read variables from a file and write manifest to a file
$manifestgen -v project/example-var.json -t project/targets.tmpl -o project/config.yaml

# read variables from a file and write manifest to an base64 encoded output file
$manifestgen -v project/example-var.json -t project/targets.tmpl -o project/encoded --encode
# verify
cat project/encoded | base64 -d

# read variables from stdin and write manifest to a file
cat project/example-var.json | $manifestgen -t project/targets.tmpl -o project/config.yaml

# read JSON variables from a file and write manifest to stdout
$manifestgen -v project/example-var.json -t project/targets.tmpl

# read YAML variables from a file and write manifest to stdout
$manifestgen -v project/example-var.yaml -t project/targets.tmpl

# read JSON variables from stdin and write manifest to stdout
cat project/example-var.json | $manifestgen --template project/targets.tmpl

# read YAML variables from stdin and write manifest to stdout
cat project/example-var.yaml | $manifestgen --template project/targets.tmpl

```
