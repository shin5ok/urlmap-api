#!/bin/bash
gcloud services enable secretmanager.googleapis.com

echo -n $GITHUB_ARGOCD | gcloud secrets create myrepo --data-file=- --locations=us-central1 --replication-policy=user-managed

echo -n $DBUSER | gcloud secrets create DBUSER --data-file=-
echo -n $DBPASS | gcloud secrets create DBPASS --data-file=-
echo -n $DBNAME | gcloud secrets create DBNAME --data-file=-
echo -n $DBHOST | gcloud secrets create DBHOST --data-file=-
