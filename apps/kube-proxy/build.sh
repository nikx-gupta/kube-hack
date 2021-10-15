
VERSION=$1
docker build -t nikxgupta/kube-proxy:v"$VERSION" --network=host .
docker push nikxgupta/kube-proxy:v"$VERSION"
