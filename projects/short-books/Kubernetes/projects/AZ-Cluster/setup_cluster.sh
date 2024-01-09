
az group create --name=kuar --location=westus

az aks create --resource-group=kuar --name=kuar-cluster
az aks get-credentials --resource-group=kuar --name=kuar-cluster

az aks install-cli 