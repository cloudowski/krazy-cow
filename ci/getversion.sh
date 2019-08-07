#!/bin/sh

VER=""

TAG="$(git tag -l --points-at HEAD)"

if [ "${TAG:-}" ];then
    VER="$TAG"
elif [ -f .version ];then
    VER="$(cat .version)"
else
    VER="latest"
fi

echo "$VER"

