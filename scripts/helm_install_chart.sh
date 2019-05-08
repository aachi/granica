#!/bin/sh
# Install
helm install --name mikrowezel_granica -f ./deployments/helm/values.yaml ./deployments/helm
