#!/bin/bash
cargo test
cargo build

# run example to stdout
cat project/example-var.json | ./target/debug/manifestgen --template project/targets.tmpl
#cat project/example-var.yaml | ./target/debug/manifestgen --template project/targets.tmpl
#cat project/invalid-input | ./target/debug/manifestgen --template project/targets.tmpl