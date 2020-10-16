#!/bin/bash

find . -name '*.md' -type f | while read fname; do
  if ! iconv -f UTF-8 "${fname}" -o /dev/null; then
    echo "${fname}: Incorrect UTF-8 encoding"
    break
  fi
done
