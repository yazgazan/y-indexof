
TESTDIR="test"
BASEDIR="init_base"

rm -rvf "$TESTDIR"
mkdir -p "$TESTDIR"
cd "$TESTDIR"
y-indexof init
cd -
diff "$TESTDIR" "$BASEDIR"


