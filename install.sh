#!/bin/sh

touch /usr/share/docker/plugins/vzlan.sock
docker-compose up -d
