#!/usr/bin/env bash

script::main() {
    local binDir="${1:?directory with binaries to run}"

    local ret=0
    for x in "$binDir/"*; do
        if [[ -x "$x" ]]; then
            echo "> running $x" >&2
            local status=0
            if ! "$x"; then
                status=$?
                ret=1
            else
                status=$?
            fi
            echo -e "> done: exit status $status\n" >&2
        fi
    done

    if [[ "$ret" -ne 0 ]]; then
        echo "^^^ at least one command failed ^^^" >&2
    else
        echo "NICE!" >&2
    fi

    return "$ret"
}

script::main "$@"
