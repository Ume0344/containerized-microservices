**How to configure Jenkins with Docker credentials**

[Tutorial](https://gcore.com/learning/building-docker-images-to-docker-hub-using-jenkins-pipelines/) 

- Install docker plugin to Jenkins.
    - Go to Manage Jenkins -> Plugins -> Install *Docker Pipeline* plugin.

- Add docker credentials to Jenkins.
    - Go to Manage Jenkins -> Credentials -> Global -> Add credentials.
    - Add usernbame and password of Dockerhub.

- Use credentials in Jenkinsfile.
