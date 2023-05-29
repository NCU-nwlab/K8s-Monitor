# K8s Monitor
本專案為建構基於eBPF的K8s監控平台  
Depends On:
- [Cilium](https://github.com/cilium/cilium)
- [Hubble](https://github.com/cilium/hubble)
- [Prometheus](https://github.com/prometheus)
- [BPFtrace](https://github.com/iovisor/bpftrace)
## Feature
- Node sources monitor eg. CPU, Memory, Disk I/O
- Network Flow Monitor
- Log Trace
## Install
### kubeadm
1. Deploy Kubernetes Using Docker CRI
```bash
$ export PODCIDR="192.168.0.0/16"
$ sudo kubeadm init \
    --apiserver-advertise-address=${MasterIP} \
    --pod-network-cidr=${PODCIDR} \
    --cri-socket /var/run/cri-dockerd.sock \
    --service-cidr=${SERVICE_CIDR} \
    --skip-phases=addon/kube-proxy
```
2. Enable -aggregator-routing  
Edit `/etc/kubernetes/manifests/kube-apiserver.yaml`
add `--enable-aggregator-routing=true` in `spec.containers.command`

3. Install Cilium
```bash
$ helm repo add cilium https://helm.cilium.io/
$ export API_SERVER_IP=YOUR_MASTER_IP
$ export API_SERVER_PORT=6443
$ helm install cilium cilium/cilium --version 1.13.0 \
    --namespace kube-system \
    --set kubeProxyReplacement=strict \
    --set socketLB.hostNamespaceOnly=true \ 
    --set hubble.relay.enabled=true \ 
    --set hubble.ui.enabled=true \
    --set prometheus.enabled=true
    --set k8sServiceHost=${API_SERVER_IP} \
    --set k8sServicePort=${API_SERVER_PORT}

```
4. Install Cilium Monitor
```bash
$ kubectl create namespcae cilium-monitoring
$ kubectl apply -f ./service-yaml/monitoring.yaml -n cilium-monitoring
```
5. Install Install metrics-server
```bash
$ helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
$ helm install metrics-server metrics-server/metrics-server -n kube-system
```

## Access Service
- Grafana:
  ```bash
  $ kubectl port-forward -n cilium-monitoring services/grafana 3000:3000
   ```
  Visit Grafana: [http://127.0.0.1:3000](127.0.0.1:3000)  
  
  Exit: `Ctrl + c` to stop fordward
  
- Hubble:
  ```bash
  $ kubectl port-fordward -n kube-system services/hubble-ui 8081:8081
  ```
  Visit Hubble: [http://127.0.0.1:8081](127.0.0.1:8081)

  Exit: `Ctrl + c` to stop fordward

## Remove
```bash
$ kubeadm reset --cri-socket unix:///var/run/cri-dockerd.sock
```

## Licience
本計畫使用Apache License 2.0，查閱[Licience.md](https://github.com/NCU-nwlab/K8s-Monitor/blob/main/LICENSE)了解更多細節
