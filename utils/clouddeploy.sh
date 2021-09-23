PROJECT=$(gcloud config get-value project)
# Create each stages and a pipeline
gcloud beta deploy apply --file clouddeploy.yaml --region=us-central1 --project=$PROJECT
# Show the list of deploy targets
gcloud beta deploy targets list --region=us-central1 --project=$PROJECT