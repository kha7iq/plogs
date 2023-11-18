<h2 align="center">
  <p align="center"><img width=30% src="./.github/img/logo.png"></p>
</h2>

# Pod Logs (plogs)
`plogs` is a kubernetes plugin that facilitates retrieving logs from Kubernetes pods with various filtering and highlighting options. It allows users to specify the namespace, container name, labels, and mark specific words for highlighting within the logs.

## Usage

### Flags

- `-m, --mark [string]`: Mark the given word or sentence in logs.
- `-n, --namespace [string]`: Specify the Kubernetes namespace.
- `-c, --containerName [string]`: Specify the container name within the pod.
- `-l, --label [string]`: Specify labels to match.
- `-p, --pod [string]`: Specify the pod name.
- `-f, --follow`: Specify to follow logs.
- `-t, --tail [int]`: Specify the number of lines from the end of the logs to show.

### Examples

#### Get logs with highlighting
```bash
plogs --namespace=default --label=app=myapp --mark=error --pod=my-pod --follow --tail=100
```
Get logs without specifying a container name

```bash
plogs --namespace=default --pod=my-pod --follow --tail=100
```

## Installation


### Prerequisites

- Kubernetes cluster access or kubeconfig set up

