#!/bin/bash
set -e

if [ "$1" = 'strace' ]; then
  strace -ff curl "https://github.com"
fi

curl "https://github.com"
# exec "$@"