#!/bin/bash
cargo build
# run example to stdout
cat project/example-var.json | ./target/debug/manifestgen --template project/targets.tmpl