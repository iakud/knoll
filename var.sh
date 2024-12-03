PWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

shopt -s expand_aliases

export PATH="`brew --prefix openjdk@11`/bin:$PATH"
#export CLASSPATH=".:$PWD/local/lib/antlr-4.13.2-complete.jar:$CLASSPATH"
#alias antlr4='java -Xmx500M -cp "$CLASSPATH" org.antlr.v4.Tool'