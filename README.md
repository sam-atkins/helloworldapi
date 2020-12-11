# Hello World API

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
    echo '{"auths":{"docker.pkg.github.com":{"auth":"<AUTH>"}}}' | kubectl create secret generic dockerconfigjson-github-com --type=kubernetes.io/dockerconfigjson --from-file=.dockerconfigjson=/dev/stdin
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

### Run locally

Instructions for running on macOS using minikube

```bash
# start minikube
minkube start

# switch to using a namespace
kubectl config set-context --current --namespace <NAMESPACE>

# apply Kubernetes config
kubectl apply -f k8s

# Start minikube tunnel service
minikube service helloworld-api-load-balancer -n <NAMESPACE>
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
http http://127.0.0.1:62282
```

## Local Docker commands

```bash
# build
docker build -t helloworldapi .

# run using local image
docker run -p 8080:8080 -it helloworldapi
```
