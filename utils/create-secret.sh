gcloud services enable secretmanager.googleapis.com
echo -n $GITHUB_ARGOCD | gcloud secrets create myrepo --data-file=- --locations=us-central1 --replication-policy=user-managed
