#!/bin/sh
# Delete
kubectl delete secret granica-envvars
# Create
kubectl create secret generic granica-envvars --from-file=configs/config.txt
