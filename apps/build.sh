
VERSION=$1
docker build -t nikxgupta/hackube:v"$VERSION" --network=host go/

docker push nikxgupta/hackube:v"$VERSION"

sed -i -E "s/nikxgupta\/hackube:.+/nikxgupta\/hackube:v$VERSION/g" deploy-hack-kube.yaml

microk8s kubectl apply -f deploy-hack-kube.yaml
