#!/bin/sh

source ../variables.env

MANIFESTS="../manifests/"

FILES=`ls $MANIFESTS`

for FILE in $FILES; do
  envsubst < ${MANIFESTS}${FILE} | kubectl apply -f -
done
