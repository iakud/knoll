#!/bin/sh

source ../../var.sh

# alias antlr4='java -Xmx500M -cp "$LOCAL_LIB/antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
# antlr4 -Dlanguage=Go -no-visitor -o parser -package parser *.g4