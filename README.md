# kubectl unlimited ♾️
A [kubectl plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) for displaying information about running containers with no `limits` set in a Kubernetes cluster.

Why would you need this? Because these pesky unlimited containers can affect other workloads in the cluster and even cause node they are running on to become unresponsive.

Have in mind that CPU limits are a somewhat controversial topic, see [CPU limits on Kubernetes are an antipattern](https://home.robusta.dev/blog/stop-using-cpu-limits) and [Kubernetes: Make your services faster by removing CPU limits](https://news.ycombinator.com/item?id=24351566).

## Installation

### krew

Using [krew](https://krew.sigs.k8s.io/) kubectl plugin manager:

```
kubectl krew install unlimited
```

### Manual

Just download the binary for your OS and architecture from the [Releases](https://github.com/nilic/kubectl-unlimited/releases) page and place it in your `PATH`.

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
  -A, --all-namespaces      analyze containers in all namespaces
      --context string      name of the kubeconfig context to use
  -h, --help                help for kubectl-unlimited
      --kubeconfig string   path to the kubeconfig file
  -l, --labels string       labels to filter pods with
  -n, --namespace string    only analyze containers in this namespace
  -o, --output string       output format, one of: [table json yaml name] (default "table")
```

## Examples

```
# get containers with either CPU or memory limits unset in the current namespace
kubectl unlimited

# get all containers with either CPU or memory limits unset
kubectl unlimited -A

# same, but in JSON
kubectl unlimited -A -o json

# get all containers with only CPU limits unset
kubectl unlimited cpu -A

# get all containers with only memory limits unset
kubectl unlimited memory -A

# get containers with only memory limits unset in namespace kube-system
kubectl unlimited memory -n kube-system
```

## Filtering output

Containers can be filtered on a pod label basis. If used, pod labels should be defined in the `key1=value1,key2=value2,..` format, e.g.

```
kubectl unlimited -l app=foo,owner=bar
```

## Output format

Default output format is `table` which prints information in a `kubectl`-like table:

```
$ kubectl unlimited memory -n kube-system
NAMESPACE     POD                               CONTAINER        CPU REQ   CPU LIM   MEM REQ   MEM LIM
kube-system   metrics-server-668d979685-tk6ws   metrics-server   100m      0m        70Mi      0Mi
kube-system   svclb-traefik-21193086-wwtg6      lb-tcp-443       0m        0m        0Mi       0Mi
kube-system   svclb-traefik-21193086-wwtg6      lb-tcp-80        0m        0m        0Mi       0Mi
kube-system   traefik-5b77545fd4-sqpzz          traefik          0m        0m        0Mi       0Mi
```

Output is sorted first by namespace, then by pod name and then by container name.

Alternatively, output format can be set to `json`, `yaml` and `name` using the `-o` or `--output` flags.

```
$ kubectl unlimited memory -n kube-system -o yaml
- limits:
    cpu: "0"
    memory: "0"
  name: metrics-server
  namespace: kube-system
  pod: metrics-server-668d979685-tk6ws
  requests:
    cpu: 100m
    memory: 70Mi
- limits:
    cpu: "0"
    memory: "0"
  name: lb-tcp-443
  namespace: kube-system
  pod: svclb-traefik-21193086-wwtg6
  requests:
    cpu: "0"
    memory: "0"
- limits:
    cpu: "0"
    memory: "0"
  name: lb-tcp-80
  namespace: kube-system
  pod: svclb-traefik-21193086-wwtg6
  requests:
    cpu: "0"
    memory: "0"
- limits:
    cpu: "0"
    memory: "0"
  name: traefik
  namespace: kube-system
  pod: traefik-5b77545fd4-sqpzz
  requests:
    cpu: "0"
    memory: "0"
```
