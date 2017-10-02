#!/bin/bash

# Colorizing Go test output:
# This is meant to be used in place of `go test`.  Provided this script is in 
# your PATH, calling `color-go-test` will call through to `go test` and then 
# colorize and reformat the output.

RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
COLOR_RESET=$(tput sgr0)
BOLD=$(tput bold)

previous_line_fail=false
verbose_output=false
verbose_flag_prefix="-v "
pass_count=0
fail_count=0
errors=()

echo_last_line() {
  local time_string=$1
  local color=$GREEN
  if [ $verbose_output = false ]; then
    echo -e "\n"
    if [ ${#errors[@]} -gt 0 ]; then
      for error in "${errors[@]}"; do
        echo -e "$error"
      done
    fi
  fi
  if [ $fail_count -gt 0 ]; then
    color=$RED
  fi
  local num_tests=$((pass_count + fail_count))
  echo -e "\n${color}${BOLD}$num_tests tests, $fail_count failure, run time ($time_string)${COLOR_RESET}"
}

colorize_output() {
  while read line; do 
    if echo $line | grep --quiet '^FAIL$'; then
      continue

    elif echo $line | grep --quiet '^PASS$'; then
      continue

    elif echo $line | grep --quiet '^=== RUN'; then
      continue

    elif echo $line | grep --quiet '^exit status 1$'; then
      continue

    elif echo $line | grep --quiet 'FAIL'; then
      if echo $line | grep --quiet "\-\-\- FAIL:"; then
        fail_count=$((fail_count + 1))
        error_message="${RED}$(echo $line | sed 's/--- FAIL:/✗/')${COLOR_RESET}"

        if [ $verbose_output = true ]; then
          echo $error_message
        else
          errors+=("$error_message")
          printf "${RED}.${COLOR_RESET}"
        fi
        previous_line_fail=true
      else
        local test_run_time=$(echo $line | grep -o '[0-9]*\.[0-9]*s$')
        echo_last_line $test_run_time
        previous_line_fail=false
      fi

    elif [ $previous_line_fail = true ]; then
      error_message="  ${YELLOW}➝ $line${COLOR_RESET}"
      if [ $verbose_output = true ]; then
        echo -e "$error_message"
      else
        errors+=("$error_message")
      fi
      previous_line_fail=false

    elif echo $line | grep --quiet 'PASS'; then
      if echo $line | grep --quiet "\-\-\- PASS:"; then
        if [ $verbose_output = true ]; then
          echo "${GREEN}$(echo $line | sed 's/--- PASS:/✔/')${COLOR_RESET}"
        else
          printf "${GREEN}.${COLOR_RESET}"
        fi
        pass_count=$((pass_count + 1))
      else
        local test_run_time=$(echo $line | grep -o '[0-9]*\.[0-9]*s$')
        echo_last_line $test_run_time
      fi

      previous_line_fail=false

    elif echo $line | grep --quiet '^ok '; then
      local test_run_time=$(echo $line | grep -o '[0-9]*\.[0-9]*s$')
      echo_last_line $test_run_time
      previous_line_fail=false

    else
      echo $line
      previous_line_fail=false
    fi
  done
}

for flag in $@; do
  if [ "$flag" = "-v" ]; then
    verbose_output=true
    verbose_flag_prefix=""
  fi
done

go test ${verbose_flag_prefix}$@ | colorize_output
