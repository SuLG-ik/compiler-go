#!/usr/bin/env sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
JAR_PATH="$ROOT_DIR/tools/antlr-4.13.1-complete.jar"
GRAMMAR_PATH="$ROOT_DIR/antlr/KotlinFunction.g4"
OUT_DIR="$ROOT_DIR/antlrgen"

rm -rf "$OUT_DIR"
java -jar "$JAR_PATH" -Dlanguage=Go -package antlrgen -no-listener -no-visitor -o "$OUT_DIR" "$GRAMMAR_PATH"