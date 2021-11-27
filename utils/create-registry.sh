gcloud beta artifacts repositories create myrepo --location=us-central1 --repository-format=docker
gcloud artifacts repositories add-iam-policy-binding myrepo \
    --location=us-central1 \
    --member=serviceAccount:$(gcloud projects describe $PROJECT --format=json | jq .projectNumber -r)-compute@developer.gserviceaccount.com \
    --role="roles/artifactregistry.reader"
