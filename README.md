# kube-hack

1. env_from_pod_manifest
    - Access Dynamic Pod manifest values which are available only after pod creation in your container 
2. downward_api_vol
    - Access Dynamic Pod manifest values from files
3. with_kube_proxy_container
    - Run a side car container for kube proxy to access API server

4. Find Token and Cert in Below location inside each Pod
  ```bash
  $ ls -al /var/run/secrets/kubernetes.io/serviceaccount/
  ```


Kubernetes API Client Libraries
1. Golang client—https://github.com/kubernetes/client-go
2. Python—https://github.com/kubernetes-incubator/client-python
3. Java client by Fabric8—https://github.com/fabric8io/kubernetes-client
4. Java client by Amdatu—https://bitbucket.org/amdatulabs/amdatu-kubernetes
5. Node.js client by tenxcloud—https://github.com/tenxcloud/node-kubernetes-client
6. Node.js client by GoDaddy—https://github.com/godaddy/kubernetes-client
7. PHP—https://github.com/devstub/kubernetes-api-php-client
8. Another PHP client—https://github.com/maclof/kubernetes-client
9. Ruby—https://github.com/Ch00k/kubr
10. Another Ruby client—https://github.com/abonas/kubeclient
11. Clojure—https://github.com/yanatan16/clj-kubernetes-api
12. Scala—https://github.com/doriordan/skuber
13. Perl—https://metacpan.org/pod/Net::Kubernetes