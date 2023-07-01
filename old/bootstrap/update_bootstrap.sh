#!/bin/bash
set -e

# assume not a symlinks
# http://stackoverflow.com/questions/59895/can-a-bash-script-tell-what-directory-its-stored-in
#http://stackoverflow.com/questions/4774054/reliable-way-for-a-bash-script-to-get-the-full-path-to-itself
pushd `dirname ${BASH_SOURCE[0]}`
BOOTSTRAP_DIR=`pwd`
popd

PROJECT_DIR=`dirname $BOOTSTRAP_DIR`

cd $PROJECT_DIR

$BOOTSTRAP_DIR/buildutil.par build //buildutil:buildutil.par

mv build/buildutil/buildutil.par $BOOTSTRAP_DIR
