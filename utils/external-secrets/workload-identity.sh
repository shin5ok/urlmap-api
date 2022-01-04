kubectl create serviceaccount --namespace default external-secrets
gcloud iam service-accounts create external-secrets
gcloud projects add-iam-policy-binding $PROJECT     --member "serviceAccount:external-secrets@$PROJECT.iam.gserviceaccount.com"     --role "roles/secretmanager.secretAccessor"
gcloud iam service-accounts add-iam-policy-binding     --role roles/iam.workloadIdentityUser     --member "serviceAccount:$PROJECT.svc.id.goog[default/external-secrets]"     external-secrets@$PROJECT.iam.gserviceaccount.com
kubectl annotate serviceaccount \
    --namespace default external-secrets \
    --overwrite \
    iam.gke.io/gcp-service-account=external-secrets@$PROJECT.iam.gserviceaccount.com
