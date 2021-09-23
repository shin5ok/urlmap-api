#!/bin/bash

PROJECT=$(gcloud config get-value project)
ZONE=us-central1-c
REGION=us-central1
NAME=$1;shift
gcloud beta container --project $PROJECT clusters create $NAME \
--zone "${ZONE}" \
--no-enable-basic-auth \
--cluster-version "1.20.9-gke.1001" \
--release-channel "regular" \
--machine-type "e2-medium" \
--image-type "COS_CONTAINERD" \
--disk-type "pd-standard" \
--disk-size "100" \
--metadata disable-legacy-endpoints=true \
--scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
--max-pods-per-node "110" \
--num-nodes "3" \
--logging=SYSTEM,WORKLOAD \
--monitoring=SYSTEM \
--enable-ip-alias \
--network "projects/${PROJECT}/global/networks/default" \
--subnetwork "projects/${PROJECT}/regions/${REGION}/subnetworks/default" \
--no-enable-intra-node-visibility \
--default-max-pods-per-node "110" \
--no-enable-master-authorized-networks \
--addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver \
--enable-autoupgrade \
--enable-autorepair \
--max-surge-upgrade 1 \
--max-unavailable-upgrade 0 \
--workload-pool "${PROJECT}.svc.id.goog" \
--enable-shielded-nodes \
--node-locations "${ZONE}" \
$@