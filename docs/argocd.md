<center><h2><b>ArgoCD</b></h2></center>

Documentation for leaning ArgoCD. 

*Significance*

If we CD with Jenkins, It has the following probelms;
- We need to install and setup tools like kubectl.
- Need to configure credentials of kubernetes in Jenkins (A security issue)
- Jenkins has no way to confirm if deployment gets successful in kubernetes. It can only be tested by adding further tests to Jenkins piepline (Major bottleneck and ArgoCD cater this problem)

**ArgoCD**
- A pull-based CD tool to automate deployment of applications on kubernetes cluster.

**ArgoCD Workflow**
- Run ArgoCD server on k8s cluster.
- Update the k8s manifest files through CI pipeline.
- ArgoCd will detect any changes in the k8s yaml file.
- And, update the deployment on cluster.

- Its always better to have separate repos for application and kubernetes configuration files. It will increase the security.

**Benefits of ArgoCD**
- k8s configuration defined as code in git repo. Config files are not applied manually from a Personal latops. Everyone has same interface (git commit) to update the k8s files.
- ArgoCD watches changes in cluster and git repo and considers git repo as Single Source of Truth. If there is a mismatch (if someone try to deploy k8s chnages directly through cluster, kubectl apply) in git repo and cluster, argocd detects it and overwrite the manual changes and again sync itself to git repo state. (if replica = 1 in git, but I manually apply k.yaml having replica = 2, argocd detects it and syenc its state to git repo whichc is replica = 1).
-  Easy rollback if soemthing went wrong or breaks in k8s cluster (by reverting the git commit).
- Cluster Disaster Discovery - If a cluster in one region gets failed, we can replicate it to other region through argocd using same git repo.
- Easy cluster access management through pull requests and merging in k8s files repo.
- No cluster credentials outside of k8s cluster
- It uses kubernetes components i.e etcd, k8s controllers.
- Git -> ArgoCD -> Kubernetes.

**How to configure ArgoCD into k8s cluster**
- argocd is itself a custom resource (Application) in k8s, so it can be configured through k8s yaml files (Just like p4 programs were being configured in my thesis).

**Questions**
- How can we have multiple clusters (deployment, stagin, production clusters) and configured with argocd such that if only deployemnt stage gets sucessful, the code deployed to staging and then to production. 
    - One way could be have separate branches of k8s yaml files repo each for different stages.
    - other way is to use kustomize.
