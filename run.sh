#!/bin/bash
cargo test
cargo build
cargo build --release

# run example to stdout
cat project/example-var.json | ./target/release/manifestgen --template project/targets.tmpl
#cat project/example-var.yaml | ./target/release/manifestgen --template project/targets.tmpl
#cat project/invalid-input | ./target/release/manifestgen --template project/targets.tmpl