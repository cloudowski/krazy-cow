#!/bin/sh

set -e

VER=""

tag="$(git tag -l --points-at HEAD &> /dev/null|tail -n1)"

if [ "${tag:-}" ];then
    VER="$tag"
elif [ "${GIT_COMMIT:-}" ];then
    VER="$(echo $GIT_COMMIT|cut -c1-7)"
elif [ -f .ci_version ];then
    VER="$(cat .ci_version)"
elif git rev-parse HEAD|cut -c1-7 &> /dev/null;then
    VER="$(git rev-parse HEAD|cut -c1-7)"
fi


echo "$VER"

