#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export RESOLUCIONES_CRUD_V2_PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/resoluciones_crud_v2/db/username --output text --query Parameter.Value)"
  export RESOLUCIONES_CRUD_V2_PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/resoluciones_crud_v2/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"