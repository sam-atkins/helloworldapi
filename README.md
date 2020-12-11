# Hello World API
- [Hello World API](#hello-world-api)
  - [Kubernetes](#kubernetes)
    - [Create secret to pull image from Github](#create-secret-to-pull-image-from-github)
    - [Local dev using Skaffold](#local-dev-using-skaffold)
    - [Run locally using Kubectrl](#run-locally-using-kubectrl)
  - [Local Docker commands](#local-docker-commands)
  - [Unit Tests](#unit-tests)

## Kubernetes

### Create secret to pull image from Github

1. Create new Github Personal Access Token with read:packages scope at https://github.com/settings/tokens/new
2. Base-64 encode <your-github-username>:<TOKEN>

    ```bash
    $ echo -n sam-atkins:<TOKEN> | base64
    <AUTH>
    ```

3. Manually create the secret:

    ```bash
    $ echo '{"auths":{"docker.pkg.github.com":{"auth":"<AUTH>"}}}' | kubectl create secret generic dockerconfigjson-github-com --type=kubernetes.io/dockerconfigjson --from-file=.dockerconfigjson=/dev/stdin
    ```

3. Now, you can reference the above secret from your pod's spec definition via imagePullSecrets field:

    ```yaml
    spec:
      containers:
      - name: your-container-name
        image: docker.pkg.github.com/<ORG>/<REPO>/<PKG>:<TAG>
      imagePullSecrets:
      - name: dockerconfigjson-github-com
    ```

### Local dev using Skaffold

Recommended as it watches for file changes locally and allows quicker develop.

Install [Skaffold](https://skaffold.dev).

Once installed, run this command:

```bash
$ skaffold dev --port-forward

Port forwarding service/helloworld-api-load-balancer in namespace devsam, remote port 8080 -> address 127.0.0.1 port 8080

# In a separate terminal, test the endpoint (below example uses HTTPie)
$ http localhost:8080

HTTP/1.1 200 OK
Content-Length: 39
Content-Type: application/json
Date: Fri, 11 Dec 2020 12:12:45 GMT

{
    "Message": "Hello, World",
    "Status": 200
}
```

### Run locally using Kubectrl

Instructions for running on macOS using minikube

```bash
# start minikube
$ minkube start

# switch to using a namespace
$ kubectl config set-context --current --namespace <NAMESPACE>

# apply Kubernetes config
$ kubectl apply -f k8s

# Start minikube tunnel service
$ minikube service helloworld-api-load-balancer -n <NAMESPACE>
# outputs something like:

|-----------|------------------------------|-------------|-------------------------|
| NAMESPACE |             NAME             | TARGET PORT |           URL           |
|-----------|------------------------------|-------------|-------------------------|
| devsam    | helloworld-api-load-balancer |        8080 | http://172.17.0.2:32464 |
|-----------|------------------------------|-------------|-------------------------|
🏃  Starting tunnel for service helloworld-api-load-balancer.
|-----------|------------------------------|-------------|------------------------|
| NAMESPACE |             NAME             | TARGET PORT |          URL           |
|-----------|------------------------------|-------------|------------------------|
| devsam    | helloworld-api-load-balancer |             | http://127.0.0.1:62282 |
|-----------|------------------------------|-------------|------------------------|
🎉  Opening service devsam/helloworld-api-load-balancer in default browser...
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.

# To also test on the command line
$ http http://127.0.0.1:62282
```

## Local Docker commands

```bash
# build
$ docker build -t helloworldapi .

# run using local image
$ docker run -p 8080:8080 -it helloworldapi
```

## Unit Tests

```bash
go test -v ./...
```
