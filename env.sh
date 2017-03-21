SCRIPT=`python -c "import os,sys; print(os.path.realpath(os.path.expanduser(sys.argv[1])))" "${BASH_SOURCE:-$0}"`
export DIR=$(dirname $SCRIPT)

export GOPATH="$DIR"
export PATH="$DIR/bin:$PATH"
