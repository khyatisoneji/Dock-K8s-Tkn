# containers-assignment
This is part of the assignments assigned under the training period to understand containers and how to deploy and run containers in Kubernetes

**How to create secrets:**

- `kubectl create secret generic mysql-root-pass --from-literal=password=rootpass`
- `kubectl create secret generic mysql-user-pass --from-literal=username=sameAsInGoServer --from-literal=password=sameAsInGoServer`
- `kubectl create secret generic mysql-db --from-literal=sameAsInGoServer`
