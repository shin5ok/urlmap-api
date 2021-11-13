kubectl create serviceaccount --namespace urlmap urlmap-api
gcloud iam service-accounts create urlmap-api
gcloud projects add-iam-policy-binding $PROJECT     --member "serviceAccount:urlmap-api@$PROJECT.iam.gserviceaccount.com"     --role "roles/secretmanager.secretAccessor"
gcloud iam service-accounts add-iam-policy-binding     --role roles/iam.workloadIdentityUser     --member "serviceAccount:$PROJECT.svc.id.goog[urlmap/urlmap-api]"     urlmap-api@$PROJECT.iam.gserviceaccount.com
kubectl annotate serviceaccount \
    --namespace urlmap urlmap-api \
    iam.gke.io/gcp-service-account=urlmap-api@$PROJECT.iam.gserviceaccount.com
