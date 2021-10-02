# kube-hack

1. env_from_pod_manifest
    - Access Dynamic Pod manifest values which are available only after pod creation in your container 
2. downward_api_vol
    - Access Dynamic Pod manifest values from files
3. with_kube_proxy_container
    - Run a side car container for kube proxy to access API server

- Find Token and Cert in Below location inside each Pod
  ```bash
  $ ls -al /var/run/secrets/kubernetes.io/serviceaccount/
  ```
