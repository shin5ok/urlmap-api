NAMESPACE=urlmap
SANAME=external-secret-o

kubectl create serviceaccount --namespace $NAMESPACE $SANAME
gcloud iam service-accounts create $SANAME
gcloud projects add-iam-policy-binding $PROJECT     --member "serviceAccount:$SANAME@$PROJECT.iam.gserviceaccount.com"     --role "roles/secretmanager.secretAccessor"
gcloud iam service-accounts add-iam-policy-binding     --role roles/iam.workloadIdentityUser     --member "serviceAccount:$PROJECT.svc.id.goog[$NAMESPACE/$SANAME]"     $SANAME@$PROJECT.iam.gserviceaccount.com
kubectl annotate serviceaccount \
    --namespace $NAMESPACE $SANAME \
    --overwrite \
    iam.gke.io/gcp-service-account=$SANAME@$PROJECT.iam.gserviceaccount.com
