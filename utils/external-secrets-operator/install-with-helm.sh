helm repo add external-secrets-operator https://charts.external-secrets.io

helm install external-secrets \
   external-secrets-operator/external-secrets \
    -n external-secrets \
    --create-namespace \
  # --set installCRDs=true
