#!/bin/bash
cargo build

#cargo run -- -v project/example-var.json -t project/targets.tmpl -o project/config.yaml
#cat project/example-var.json | ./target/debug/manifestgen -t project/targets.tmpl -o project/config.yaml
#./target/debug/manifestgen -v project/example-var.json -t project/targets.tmpl

cat project/example-var.json | ./target/debug/manifestgen --template project/targets.tmpl