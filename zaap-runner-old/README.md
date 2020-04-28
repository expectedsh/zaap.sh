# zaap-runner

### Requirements

- Docker
- Traefik

### Install docker

```bash
curl https://get.docker.com | sh
```

### Install traefik

Using docker swarm :
```bash
docker network create -d overlay traefik

docker service create \
  --name zaap-etcd \
  --network traefik \
  quay.io/coreos/etcd:v2.3.8 \
  -name etcd0 \
  -advertise-client-urls http://zaap-etcd:2379,http://zaap-etcd:4001 \
  -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
  -initial-advertise-peer-urls http://localhost:2380 \
  -listen-peer-urls http://localhost:2380 \
  -initial-cluster-token etcd-cluster-1 \
  -initial-cluster etcd0=http://localhost:2380 \
  -initial-cluster-state new

docker service create \
  --name zaap-traefik \
  -p 80:80 \
  -p 443:443 \
  -p 8080:8080 \
  --mount type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock \
  --constraint=node.role==manager \
  --network traefik \
  traefik:v2.2 \
  --providers.docker=true \
  --providers.docker.swarmMode=true \
  --providers.docker.exposedByDefault=false \
  --entrypoints.http.address=:80 \
  --entrypoints.https.address=:443 \
  --api.insecure=true \
  --providers.docker.network=traefik \
  --certificatesResolvers.simple.acme.email=youremail@address.com \
  --certificatesResolvers.simple.acme.httpChallenge.entryPoint=http \
  --providers.etcd.endpoints=zaap-etcd:2379
```

### Install runner

```bash
# Used to configure forward our dns entries to your server
# Ips are separated by a comma
export EXTERNAL_IPS=127.0.0.1

# Used for authentication things
export RUNNER_TOKEN=my-token
```
