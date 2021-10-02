# kube-hack

1. env_from_pod_manifest
    - Access Dynamic Pod manifest values which are available only after pod creation in your container 
2. downward_api_vol
    - Access Dynamic Pod manifest values from files


- Find Secrets and Cert in Below location inside each Pod
  ```bash
  $ ls -al /var/run/secrets/kubernetes.io/serviceaccount/
  ```
