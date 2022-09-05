# OS_Metrics for Test Task
The repository is for Junior DevOps Task.

[![Os_Metrics_Server](https://github.com/MojitoTea/Os_Metrics_Server/actions/workflows/cicd.yml/badge.svg)](https://github.com/MojitoTea/Os_Metrics_Server/actions/workflows/cicd.yml)


## Description
A simple application for obtaining data from the system.
Metrics are output in json format and displayed on the web server

#### Requirements for manual deploy
There are packages below that should be installed on the (local) host where you'll be running deploy:
 * golang >= 1.15
 * terraform >= 1.0.0
 * ansible >= 2.9
 * aws cli 
 
## Supported OS

||loadavg|uptime|cpu|memory|network|disk i/o|
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
|Linux|yes|yes|yes|yes|yes|yes|
|Darwin|yes|yes|*1|yes|yes|no|
|FreeBSD|yes|yes|no|yes|yes|no|
|NetBSD|yes|yes|no|no|yes|no|
|OpenBSD|yes|yes|no|no|no|no|
|Windows|no|yes|no|yes|no|no|

### CI/CD 
Workflows:
* [cicd.yml](/.github/workflows/cicd.yml)               - the main workflow (Linter, Build and push docker image), 

Jobs:
  * Golang Linter with Review Dog
  * Build and push docker image 
  
  ### Installation instructions to build the project by **manual**

*1. Get source code for install project:*  
```
git clone https://github.com/MojitoTea/Os_Metrics_Server
```
*2. Go to the project folder

*3. Deploy the infrastructure. For all deploy by terraform.
```
terraform init
terraform apply
```
*4.  Go to ansible directory and install APP by Ansible:*
```
ansible-playbook web-server-provision.yaml
```

*5 It’s important that you secure your app with HTTPS. To accomplish this, you’ll deploy nginx-proxy via Docker Compose, along with its Let’s Encrypt add-on. This secures Docker containers proxied using nginx-proxy, and takes care of securing your app through HTTPS by automatically handling TLS certificate creation and renewal.
Here you’re defining two containers: one for nginx-proxy and one for its Let’s Encrypt add-on (letsencrypt-nginx-proxy-companion). For the proxy, you specify the image jwilder/nginx-proxy, expose and map HTTP and HTTPS ports, and finally define volumes that will be accessible to the container for persisting Nginx-related data.

In the second block, you name the image for the Let’s Encrypt add-on configuration. Then, you configure access to Docker’s socket by defining a volume and then the existing volumes from the proxy container to inherit. Both containers have the restart property set to always, which instructs Docker to always keep them up (in the case of a crash or a system reboot).
```
docker-compose -f nginx-proxy-compose.yaml up -d
```

*6 Remember to replace your_domain both times with your domain name. Save and close the file.

This Docker Compose configuration contains one container (go-web-app), which will be your Go web app. It builds the app using the Dockerfile you’ve created in the previous step, and takes the current directory, which contains the source code, as the context for building. Furthermore, it sets two environment variables: VIRTUAL_HOST and LETSENCRYPT_HOST. nginx-proxy uses VIRTUAL_HOST to know from which domain to accept the requests. LETSENCRYPT_HOST specifies the domain name for generating TLS certificates, and must be the same as VIRTUAL_HOST, unless you specify a wildcard domain.

```
docker-compose -f go-app-compose.yaml up -d
```

*7 You can build and run the container yourself using commands
```
docker built
docker run
```
