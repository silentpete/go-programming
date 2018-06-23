#!/bin/sh

## Run this file while writing code in the subdirectories. If test files are found in the
#  directory of the changed file, the tests will be run.

## Example from: http://tysonmaly.com/programming/go/test-driven-developement-with-golang-on-osx/
# inotifywait -mr . -e close_write | while read path action file; do
#   if [[ "$file" =~ .*go$ ]]; then
#     go test
# #   go test | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
#   fi
# done
## Found that the above doesn't fully work, need to run in the directory, so pass the $path.

inotifywait --monitor --recursive . --event close_write | while read path action file; do
  if [[ "$file" =~ .*go$ ]]; then
    output=$(go test $path)
    if [[ "$?" -eq "0" ]]; then
      echo -e "\e[32mPASS\e[0m"
    elif [[ "$?" -eq "1" ]]; then
      echo -e "\e[31mFAIL\e[0m"
    else
      echo 'should never see this, something is wacky'
      echo -e "status code: $?"
    fi
    echo -e "edited file: $file"
    echo -e "$output"
  fi
done
