#!/bin/sh
docker build -t wtchat:latest . && \
docker run --rm --name chat --network=public -p ':3122:3122/udp' \
--label 'traefik.enable=true' \
--label 'traefik.http.routers.chat.entrypoints=localsecure' \
--label 'traefik.http.routers.chat.rule=Host(`chat.jph2.tech`)' \
wtchat:latest
