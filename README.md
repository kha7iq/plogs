<h2 align="center">
  <p align="center"><img width=30% src="./.github/img/logo.png"></p>
</h2>

<p align="center">
  <a href="#usage">Usage</a> •
  <a href="#installation">Installation</a> •
  <a href="#flags">Flags</a> •
  <a href="#contributing">Contributing</a> •
</p>


# Pod Logs (plogs)
`plogs` is a kubernetes plugin that facilitates retrieving logs from Kubernetes pods with various filtering and highlighting options. It allows users to specify the namespace, container name, labels, and mark specific words for highlighting within the logs stream.


## Usage 

<img alt="plogs" src="./.github/img/plogs.gif" width="1000" />

## Installation

   **You can install the "plogs" plugin using** `krew`

```bash
kubectl krew install plogs
```

#### Alternative Installation Methods

<details>
    <summary>Linux Packages</summary>

- Bash Install Script

  By default, plogs is going to be installed at `/usr/bin/`. `sudo` privileges are required for this operation.

  If you would like to provide a custom install path, you can do so as an input to the script. 
  For example, you can run `./install.sh $HOME/bin` to install plogs in the specified directory.

```bash
curl -sL https://bit.ly/installplogs | sudo sh
```
OR
```bash
curl -s https://raw.githubusercontent.com/kha7iq/plogs/main/install.sh | sudo sh
```
- DEB & RPM
```bash
# DEB
export PLOGS_VERSION="0.1.1"
wget -q https://github.com/kha7iq/plogs/releases/download/v${PLOGS_VERSION}/plogs_amd64.deb
sudo dpkg -i plogs_amd64.deb
# RPM
sudo rpm -i plogs_amd64.rpm
```

</details>

<details>
    <summary>Manual</summary>

```bash
# Chose desired version
export PLOGS_VERSION="0.1.1"
wget -q https://github.com/kha7iq/plogs/releases/download/v${PLOGS_VERSION}/plogs_linux_amd64.tar.gz && \
tar -xf plogs_linux_amd64.tar.gz && \
chmod +x plogs && \
sudo mv plogs /usr/bin/kubectl-plogs
```
</details>

#### Checkout [release](https://github.com/kha7iq/plogs/releases) page for all supported OS binaries.

### Flags

- `--mark, -m`
  - Usage: `--mark <word>`
  - Alias: `-m`
  - Marks logs containing the specified word or sentence.

- `--namespace, -n`
  - Usage: `--namespace <namespace>`
  - Alias: `-n`
  - Specifies the Kubernetes namespace. If not specified, uses the namespace set in the context. If context namespace is also empty, defaults to "default".

- `--container, -c`
  - Usage: `--container <containerName>`
  - Alias: `-c`
  - Specifies the container name within the pod. If not specified and the pod has multiple containers, defaults to the first container.

- `--label, -l`
  - Usage: `--label <labels>`
  - Alias: `-l`
  - Specifies labels to match for pods.

- `--pod, -p`
  - Usage: `--pod <podName>`
  - Alias: `-p`
  - Specifies the pod name.

- `--follow, -f`
  - Usage: `--follow`
  - Alias: `-f`
  - Specifies to follow logs continuously.

- `--tail, -t`
  - Usage: `--tail <lines>`
  - Alias: `-t`
  - Specifies the number of lines to show from the end of logs.

#### Alias
Consider adding this alias to your `.bashrc` or `.zshrc` for convenience.
```bash
echo "alias kpl='kubectl plogs'" >> .bashcr

echo "alias kplf='kubectl plogs -f'" >> .bashrc

``` 

#### Examples

1. Retrieve logs from all pods matching label

```bash
kpl -n dev -l app=myapp -m "fail" -p my-pod -t 10
```
2. Retrieve logs for a specific pod:

```bash
kplf -n dev -p my-pod -t 10
```
### Prerequisites
- Kubernetes cluster access or kubeconfig set up

## Issues
If you encounter any problems or have suggestions for improvements, please [open an issue](https://github.com/kha7iq/plogs/issues) on GitHub.

## Contributing

Contributions, issues and feature requests are welcome!<br/>Feel free to check
[issues page](https://github.com/kha7iq/plogs/issues). You can also take a look
at the [contributing guide](https://github.com/kha7iq/plogs/blob/master/CONTRIBUTING.md).