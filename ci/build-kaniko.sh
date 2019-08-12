
TARGET=${1:-cloudowski/test1:latest}
DOCKERFILE=${2:-Dockerfile.test}
KANIKO_OPTS="--insecure --skip-tls-verify"

if [ "${JOB_URL}" ];then
    echo "Jenkins environment detected" >&2
    /kaniko/executor -f `pwd`/$DOCKERFILE -c `pwd` --destination $TARGET $KANIKO_OPTS
else
    echo "Standalone environment detected - using docker runtime" >&2
    docker run --rm -ti -v `pwd`:/context -v $HOME/.docker/config.json.plain:/kaniko/.docker/config.json gcr.io/kaniko-project/executor:debug-v0.10.0 \
        -f /context/$DOCKERFILE -c /context --destination $TARGET $KANIKO_OPTS
fi
