#!/bin/sh
docker build . -t reptile-generator:latest && \
docker run --name reptile --network=public --restart=unless-stopped --detach \
--label 'traefik.enable=True' \
--label 'traefik.http.routers.reptile.entrypoints=localsecure' \
--label 'traefik.http.routers.reptile.rule=Host(`reptile.jph2.tech`)' \
--label 'traefik.http.services.reptile.loadbalancer.server.port=80' \
reptile-generator:latest
