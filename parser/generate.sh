#!/bin/bash

source ~/.venv/bin/activate

export ANTLR4_TOOLS_ANTLR_VERSION=4.13.2

#antlr4 -Dlanguage=Go -no-visitor -package parsing -o ../parsing *.g4
antlr4 -Dlanguage=Go  -package parsing -o ../parsing *.g4
