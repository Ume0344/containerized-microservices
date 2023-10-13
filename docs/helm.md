# Helm
- A package manager for kubernetes 
- Manages and package all of the manifest files (Helm Chart) needed for the applications for public and private distribution.
- Forexample, Elastic stack deployment for logging has the same setup configuration. So, if we have helm chart(bundle of yaml files) in helm hub for it, we donot have to create manifest files again.
- Helm is a Templating Engine: A blueprint to use same configuration files for similar applications. Forexample, we have three applications but their docker image tags are different. We can make the image attribute inside the template config file as dynamic value ({{.Values.container.image}}). The dynamic values then can be replaced by the values define in values.yaml file.

## Heml Chart Structure

mychart/
    Chart.yaml - Meta info about the chart (name etc)
    values.yaml - values for the template config file
    charts/ - chart dependencies will be stored here if this chart has any other dependencies.
    templates/ - the actual template files are stored here.
```
helm install <chartname> to deploy the chart.
```
## Value injection into template files

We can chnage the values defined in values.yaml file by two ways;
1- by creating a new values.yaml(version=1.1) file (values1-yaml (version=2.1)) and mention the values we want to override.
```
helm install --values=values1.yaml <chartname> 
```
2- by using set flag;
```
helm install --set version=2.1
```

## Helm Installation
- Donwload the [suitable release](https://github.com/helm/helm/releases)
- Unpack it (`tar -zxvf helm-v3.13.1-linux-amd64.tar.gz`)
- `mv linux-amd64/helm /usr/local/bin/helm`
