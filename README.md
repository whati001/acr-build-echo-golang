# Azure ACR-task golang
This repository includes a simple documentation how to enable automate container image builds on Azure. 

We use the Azure Container Registry Tasks to watch upcomming changes in your github repository. The container is running a simple [echo](https://echo.labstack.com/) server written in [golang](https://golang.org/).

# Get the source
Feel free to download the repository via git.
```bash
git clone https://github.com/whati001/acr-build-echo-golang.git
```

# Build container locally
For testing purpose, we can build the docker container on your local pc first.
```bash
# inside your repository root
docker build -t acr-golang-test .
docker run --rm -it -p 2711:2711 acr-golang-test
```

Now we can simply test if the server is listening on port `2711` via [curl](https://curl.se/).
```bash
curl http://localhost:2711/route/entpoint?param="SomeValueWeWantToPass"
```
# Create a Azure Container Registry
At first we need to create a **Azure Container Registry**.
Please reference to there page [here](https://docs.microsoft.com/en-us/azure/container-registry/).

For creating the resource group and container registry, I have pesonally used the web portal. Seems simpler to me.

# Link github repo
Please checkout [here](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-tutorial-build-task)

Now we need to create a task, which listens for commit changes in your public repository. 
If you have not installed the Azure shell on your pc, just use the web version [here](https://shell.azure.com/)

```bash
# define some global vars
ACR_NAME=whati001
GIT_PAT=<fromGithubToken>

# create a new task
az acr task create \
    --registry $ACR_NAME \
    --name taskacrgolang \
    --image acrgolang:{{.Run.ID}} \
    --context https://github.com/whati001/acr-build-echo-golang.git \
    --file Dockerfile \
    --git-access-token $GIT_PAT

# trigger by hand for testing
az acr task run --registry $ACR_NAME --name taskacrgolang
```
If both jobs finished without an error, you can see the build container image in your registry.

# Create container instance
Finally we are able to create a container instance from it.
For this, please follow [this](https://docs.microsoft.com/en-us/azure/container-instances/).

Please ensure to create an instance from your created container.
You need to enable **admin** stuff in your registry and make port `2711` public.
