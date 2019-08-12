#!/bin/bash

set -e

CREATE_PREVIEW="n"
IMGNAME="cloudowski/test1"
ENVPREFIX="kcow"
BASEDIR=deploy/envs

while getopts "pt:v:i:" OPTION;do
    case $OPTION in
        h) print_usage; exit 0; ;;
        p) CREATE_PREVIEW="y";;
        t) TARGET_ENV="$OPTARG";;
        v) IMGTAG="$OPTARG";;
        i) IMGNAME="$OPTARG";;
    esac
done


in_ci() {
    [ "${JOB_URL:-x}x" != "xx" ]
    return $?
}

get_preview_name() {
    local suffix
    if in_ci;then
        if [ "$GIT_BRANCH" ];then
            suffix="$GIT_BRANCH"
        else
            suffix=`git rev-parse --abbrev-ref HEAD`
        fi
    else
            suffix=`git rev-parse --abbrev-ref HEAD`
    fi

    # TODO: sanitize name - e.g. replace "/" with "-"
    echo "${ENVPREFIX}-${suffix}"
}




[ "${IMGTAG:-}" ] || IMGTAG=`ci/getversion.sh`

if [ "${CREATE_PREVIEW}" = "y" ];then
    echo "Creating preview environment from preview-template"
    TARGET_ENV=preview-template
elif ! [ "$TARGET_ENV" ];then
    echo "Usage: $0 [ -t ENVIRONMENT_NAME | -p ] [ -v VERSION ] [ -i IMAGE_NAME ]"
    exit 2
fi

cd $BASEDIR/$TARGET_ENV

if ! in_ci;then
    # save original kustomization.yaml
    cp kustomization.yaml kustomization.yaml.orig
fi


kustomize edit set image cloudowski/krazy-cow=$IMGNAME:$IMGTAG

if [ "${CREATE_PREVIEW}" = "y" ];then
    NS="`get_preview_name`"
    echo "# Setting new namespace for preview: $NS" >&2
    kustomize edit set namespace ${NS}

    cat << EOF > namespace.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: $NS
EOF
    kustomize edit add resource namespace.yaml
fi


kustomize build|kubectl apply -f-

if ! in_ci;then
    # save original kustomization.yaml
    cp kustomization.yaml.orig kustomization.yaml
fi
