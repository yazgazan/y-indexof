
TESTDIR="test"
BASEDIR="init_base"

rm -rvf "$TESTDIR"
mkdir -p "$TESTDIR"

# Testing init ...
cd "$TESTDIR"
y-indexof init -f ../init.tar -l
cd -
diff --exclude=.git -r "$TESTDIR" "$BASEDIR"

# Testing start ...
cd "$TESTDIR"
y-indexof start
cd -

