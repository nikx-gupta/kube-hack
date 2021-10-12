1. Enable RBAC on respective kubernetes environment
    - On microk8s
      ```bash
      microk8s enable rbac
      ```

2. Create Two namespaces namely "order" and "notify"
    ```bash
    $ kubcetl create ns order
    $ kubectl create ns notify
    ```

3. Two default service account in respective namespaces are enabled with name "default"
4. The default service account in new namespace is not allowed to access services for that namespace
5. Try to "list services" in order namespace from a pod inside order namespace
    1. Create a pod inside "order" namespace
        ```bash
        $ kubectl run pod-orderns -n order --image=luksa/kubectl-proxy
        $ kubectl run pod-notifyns -n notify --image=luksa/kubectl-proxy
        ```
    2. ssh into pod and try to "list services"
        ```bash
        $ kubectl exec -it pod pod-orderns -- sh
        $ curl localhost:8001/api/v1/namespaces/order/services
        
        $ kubectl exec -it pod pod-notify -- sh
        $ curl localhost:8001/api/v1/namespaces/notify/services
        ```
    3. Receive no acces error
6. Create a Role to allow access to "list services" in order namespace
    ```bash
    $ kubectl create role orderns-role -n order --verb=get,list --resources=services
    ```
7. To allow a specific subject to access the resource now we have to create a Role Binding for that subject
8. Create RoleBinding to subject i.e. in our case "default service account" inside order namespace
    ```bash
    $ kubcetl create rolebinding orderns-role-bind --role=orderns-role --serviceaccount=order:default -n order 
    ```
9. Try executing curl command in Step 5.2 again. the calls made from inside pod running under order namespace shall be fine
10. Let's allow the "default serviceaccount" in notify namespace to list services in order namespace. To do this
    we will edit the above rolebinding and add notify namespace as subject
    ```bash
    $ KUBE_EDITOR=nano kubcetl edit rolebinding orderns-role-bind 
    ```
    - Add below in the yaml under "Subjects"
    ```text
    subjects:
    - kind: ServiceAccount
      name: default
      namespace: notify
    ```
11. The pod running under notify namespace should be able to list the services inside order namespace
