#!/bin/bash

# rustup target add x86_64-unknown-linux-musl

cargo test
cargo build
cargo build --target x86_64-unknown-linux-musl --release

# run example to stdout
manifestgen='./target/x86_64-unknown-linux-musl/release/manifestgen'
cat project/example-var.json | $manifestgen --template project/targets.tmpl
cat project/example-var.yaml | $manifestgen --template project/targets.tmpl
echo ""
cat project/example-var.yaml | $manifestgen --template project/targets.tmpl --encode
#cat project/invalid-input | ./target/x86_64-unknown-linux-musl/release/manifestgen --template project/targets.tmpl