#!/bin/bash

set -e
set -u
set -o pipefail
set -x

rm -rf /dist
bash run_build.sh
ret=$?
if [ $ret -ne 0 ]; then
  echo $ret
  exit 0
fi

# docker registry

docker build --no-cache --pull -f CI/Dockerfile -t gitjaesik/dice-api-server:latest .
# docker tag gitjaesik/dice-api-server:latest gitjaesik/dice-api-server:${BUILD_NAME}
# docker push gitjaesik/dice-api-server:latest
# docker push gitjaesik/dice-api-server:${BUILD_NAME}

# docker images
# docker rmi gitjaesik/dice-api-server:latest
# docker rmi gitjaesik/dice-api-server:${BUILD_NAME}
# docker images
