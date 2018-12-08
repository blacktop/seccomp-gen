#!/bin/bash
set -e

if [ "$1" = 'strace' ]; then
  strace -c -f -S name curl "https://github.com" 2>&1 1>/dev/null | tail -n +3 | head -n -2 | awk '{print $(NF)}'
  exit 0
fi

# curl "https://github.com"
exec "$@"