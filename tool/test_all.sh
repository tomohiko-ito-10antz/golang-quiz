#!/bin/sh

for Q in $(ls | grep quiz-); do
    echo "${Q} : $(sh tool/test.sh ${Q} 2> /dev/null)"
done
