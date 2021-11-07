#!/bin/bash
# Exit with an error code if the commmited .md files are not UTF-8 compliant.

find . -name '*.md' -type f | while read fname; do
  if ! iconv -f UTF-8 "${fname}" -o /dev/null; then
    echo "${fname}: Incorrect UTF-8 encoding"
    break
  fi
done
