#!/bin/bash

# rustup target add x86_64-unknown-linux-musl

cargo test
#cargo build
#cargo build --release
cargo build --target x86_64-unknown-linux-musl --release

# run example to stdout
cat project/example-var.json | ./target/x86_64-unknown-linux-musl/release/manifestgen --template project/targets.tmpl
cat project/example-var.yaml | ./target/x86_64-unknown-linux-musl/release/manifestgen --template project/targets.tmpl
#cat project/invalid-input | ./target/x86_64-unknown-linux-musl/release/manifestgen --template project/targets.tmpl