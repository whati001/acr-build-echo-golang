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
