gcloud iam service-accounts create gmp-sa \
&&
gcloud iam service-accounts add-iam-policy-binding \
  --role roles/iam.workloadIdentityUser \
  --member "serviceAccount:$PROJECT.svc.id.goog[gmp-sa/default]" \
  gmp-sa@$PROJECT.iam.gserviceaccount.com \
&&
kubectl annotate serviceaccount \
  --namespace urlmap \
  default \
  iam.gke.io/gcp-service-account=gmp-sa@$PROJECT.iam.gserviceaccount.com

gcloud projects add-iam-policy-binding $PROJECT \
  --member=serviceAccount:gmp-sa@$PROJECT.iam.gserviceaccount.com \
  --role=roles/monitoring.viewer
