#!/bin/sh

set -eux

if [ $1 = "" ]; then
    echo "positional argument <quiz-dir> is required"
    exit 1
fi

QUIZ=$1

: > "${QUIZ}/out"

go build -o "${QUIZ}/main" "${QUIZ}/main.go"
if [ $? -ne 0 ]; then
    echo "1" >> "${QUIZ}/out"
else 
    set +e
    "./${QUIZ}/main" >> "${QUIZ}/out"
    if [ $? -ne 0 ]; then
        echo "2" >> "${QUIZ}/out"
    fi
    set -eux
    sed -i.backup -e 's|X|3|g' "${QUIZ}/out"
    sed -i.backup -e 's|Y|4|g' "${QUIZ}/out"
fi

show_diff () {
    cat "${QUIZ}/ans"
    cat "${QUIZ}/out"
}

if [ "$(grep 1 < ${QUIZ}/ans || true)" != "$(grep 1 < ${QUIZ}/out || true)" ]; then echo "NG" && show_diff && exit 1; fi
if [ "$(grep 2 < ${QUIZ}/ans || true)" != "$(grep 2 < ${QUIZ}/out || true)" ]; then echo "NG" && show_diff && exit 1; fi
if [ "$(grep 3 < ${QUIZ}/ans || true)" != "$(grep 3 < ${QUIZ}/out || true)" ]; then echo "NG" && show_diff && exit 1; fi
if [ "$(grep 4 < ${QUIZ}/ans || true)" != "$(grep 4 < ${QUIZ}/out || true)" ]; then echo "NG" && show_diff && exit 1; fi

echo "OK"
