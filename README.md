# kubectl unlimited ♾️
A kubectl plugin for displaying information about running containers with no `limits` set in a Kubernetes cluster.

Why would you need this? Because these pesky unlimited containers can affect other workloads in the cluster and even cause node they are running on to become unresponsive.

Have in mind that CPU limits are a somewhat controversial topic, see [CPU limits on Kubernetes are an antipattern](https://home.robusta.dev/blog/stop-using-cpu-limits) and [Kubernetes: Make your services faster by removing CPU limits](https://news.ycombinator.com/item?id=24351566).

## Installation

Just download the binary from the [Releases](https://github.com/nilic/kubectl-unlimited/releases) page and place it in your `PATH`.

## Usage

```
Usage:
  kubectl unlimited [flags]
  kubectl unlimited [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  cpu         Display information about running containers with no CPU limits set
  help        Help about any command
  memory      Display information about running containers with no memory limits set
  version     Print kubectl-unlimited version

Flags:
      --context string      name of the kubeconfig context to use
  -h, --help                help for kubectl-unlimited
      --kubeconfig string   path to the kubeconfig file
  -l, --labels string       labels to filter pods with
  -n, --namespace string    only analyze containers in this namespace (by default all containers from all namespaces are shown)
  -o, --output string       output format, one of: [table json yaml] (default "table")
```

## Examples

```
# get containers with either CPU or memory limits unset
$ kubectl unlimited

# same, but in JSON
$ kubectl unlimited -o json

# get containers with only CPU limits unset
$ kubectl unlimited cpu

# get containers with only memory limits unset
$ kubectl unlimited memory
```

## Filtering output

By default, all unlimited containers from all namespaces are shown. This can be filtered on a namespace or pod label basis using appropriate flags (see [Usage](#usage)). If used, pod labels should be defined in the `key1=value1,key2=value2,..` format, e.g.

```
kubectl unlimited -l app=foo,owner=bar
```