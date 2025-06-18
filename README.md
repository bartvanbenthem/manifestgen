# hbstemplate - template renderer 
hbstemplate is a lightweight and fast command-line tool written in Rust that lets you generate output files from templates using the powerful Handlebars templating engine. Whether you're scaffolding config files, boilerplate code, or documentation, hbstemplate makes it easy to automate file generation with structured input data in JSON or YAML format.

# usage manual
```bash
Hbstemplate 

USAGE:
    hbstemplate [FLAGS] [OPTIONS] --template <template_file>

FLAGS:
    -e, --encode     Base64 Encoded output
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -o, --output <output_file>          Path for the generated output file
    -t, --template <template_file>      Path to the template file
    -v, --variables <variables_file>    Path to the variable input file

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
hbstemplate='./target/x86_64-unknown-linux-musl/release/hbstemplate'

# read variables from a file and write manifest to a file
$hbstemplate -v project/example-var.json -t project/targets.tmpl -o project/config.yaml

# read variables from a file and write manifest to an base64 encoded output file
$hbstemplate -v project/example-var.json -t project/targets.tmpl -o project/encoded --encode
# verify
cat project/encoded | base64 -d

# read variables from stdin and write manifest to a file
cat project/example-var.json | $hbstemplate -t project/targets.tmpl -o project/config.yaml

# read JSON variables from a file and write manifest to stdout
$hbstemplate -v project/example-var.json -t project/targets.tmpl

# read YAML variables from a file and write manifest to stdout
$hbstemplate -v project/example-var.yaml -t project/targets.tmpl

# read JSON variables from stdin and write manifest to stdout
cat project/example-var.json | $hbstemplate --template project/targets.tmpl

# read YAML variables from stdin and write manifest to stdout
cat project/example-var.yaml | $hbstemplate --template project/targets.tmpl

```
