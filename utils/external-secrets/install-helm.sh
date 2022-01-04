# ref: https://qiita.com/MahoTakara/items/3c509235cc18bd407f07#%E4%BA%8B%E5%89%8D%E3%81%AB%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%9E%E3%82%A4%E3%82%BA%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95
echo "Read this script"
sleep 1
# 1. Do inspect to generate yaml as helm-external-secrets.yaml
# 2. Edit the part of serviceAccount to add below,
#    - metadata.annotations.iam.gke.io/gcp-service-account
#    - metadata.name

helm --kube-context=stage-urlmap install -f helm-external-secrets.yaml external-secrets external-secrets/kubernetes-external-secrets
