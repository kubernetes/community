#/bin/bash

CONTENT_PATH="${PWD}/png"
OUTPUT_FILENAME="kubernetes-visio-stencil"

curl -s -N https://raw.githubusercontent.com/hoveytechllc/visio-stencil-creator/master/scripts/build-and-run.sh | bash -s -- --content-path=${CONTENT_PATH} --output-filename=${OUTPUT_FILENAME}

if [ -f ${CONTENT_PATH}/${OUTPUT_FILENAME}.vssx ]; then
    mv ${CONTENT_PATH}/${OUTPUT_FILENAME} ./visio/${OUTPUT_FILENAME}.vssx
else
    echo "ERROR -> Visio Stencil was not in expected path: "${CONTENT_PATH}/${OUTPUT_FILENAME}
    exit 1
fi
