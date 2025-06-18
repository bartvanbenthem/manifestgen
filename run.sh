#!/bin/bash

# rustup target add x86_64-unknown-linux-musl

cargo test
cargo build
cargo build --target x86_64-unknown-linux-musl --release

# run example to stdout
hbstemplate='./target/x86_64-unknown-linux-musl/release/hbstemplate'
cat project/example-var.json | $hbstemplate --template project/targets.tmpl
cat project/example-var.yaml | $hbstemplate --template project/targets.tmpl
echo ""
cat project/example-var.yaml | $hbstemplate --template project/targets.tmpl --encode
#cat project/invalid-input | ./target/x86_64-unknown-linux-musl/release/hbstemplate --template project/targets.tmpl