#!/bin/bash

###
#
# @usage  `source set_env.sh`
# @desc   Sets the environment for developers
#
###

#Get the script's own directory instead of where it's being called from
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Set PATH
if [[ "$PATH" != ?(*:)"$DIR/bin"?(:*) ]]; then
  export PATH=$PATH:$DIR/bin
fi

# Set GOPATH - used by GO
if [[ -z "$GOPATH" ]]; then
  export GOPATH=$DIR
fi