#!/bin/sh

kubectl create -f virtimagefile.yaml
kubectl create -f virtimagerepo.yaml
kubectl create -f virtnode.yaml
kubectl create -f virtmachine.yaml
echo "Sleeping 10 secs for TPRs to become activate"
sleep 10
