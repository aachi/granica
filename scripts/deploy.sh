#!/bin/sh
# Build
./scripts/docker_build.sh
./scripts/helm_delete_release.sh
./scripts/envvar_update.sh
./scripts/config_secrets_update.sh
./scripts/helm_install_chart.sh