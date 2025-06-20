# GITOPS WITH A GO-WEBAPP

![Image](https://github.com/user-attachments/assets/6d59bb8c-eff0-4baa-a041-346116d7a793)

# Go web app

This is a simple Static website in Golang. It uses the `net/http` package to serve HTTP requests.

## Running the server

To run the server, execute the following command:

```bash
go run main.go
```
The server will start on port 8080. Access it by this url -> http://localhost:8080 in your web browser.


# Steps for the complete Gitops-cicd

# Prereq
1. install the eksctl 
2. create the docker account
3. github
4.  install kubectl 
5. install Helm after the cluster setup


# Create the eks cluster

``` bash
eksctl create cluster --name anyname --region ur-region
```

check it by 
- kubectl config current-context

# Install the helm and create the resources by the helm 

``` bash
helm install release-name  ./Helm-charts/go-webapp-cicd-chart
```

# Install the nginx controller 

``` bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml

```

# Github action

1 . generate the accesss tokens from the Dockerhub and github and store it in the secrets in the github with the following name in the .Github/workflows file
2 . Now any push happens the github action will be triggered


# Argo cd Setup

``` bash
kubectl create namespace argocd - create the namespace for the argocd

kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml - install the argo cd 

 kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'  - it changes the type to lb so we can access the argocd ui on our browser by  "Kubectl get svc  argocd-server -n argocd -o wide"  and copy the dns shown and pass it in ur browser to access the ui 

```
In the ui page give username as admin and for the password:

->  Kubectl get secrets -n argocd 
-> kubectl get secrets argocd-initial-admin-secret -n argocd -o yaml
-> Copy the secret code and decode it by 

-> echo " ex-that code ie.like hsfbdshakqw=? " | base64 --decode 

Paste the decode code in the password and ui page will be authenticated

# Go Daddy Dns mapping

-> go to the godaddy and dns settings 
-> create the new record as cname
-> And give the Dns of lb, TO view that
``` bash
kubectl get ing  -> copy the dns 
```
-> paste the dns in the value section and give domain name as anything but make sure to update tht domain name in the Helm chart's ingress.yaml file in source code.

Now you can access the site by the custom domain name in the browser



# Clean up 
-> kubectl delete -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

-> kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml

-> helm uninstall releasename

->  eksctl delete cluster --name clustername --region ur region




