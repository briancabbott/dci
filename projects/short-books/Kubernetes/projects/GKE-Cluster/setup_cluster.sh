gcloud config set account briancabbott@gmail.com
gcloud config set project kubernetes-409717
gcloud config set compute/zone us-west1-a

gcloud container clusters create kuar-cluster

gcloud auth application-default login
