#!/bin/bash

docker-compose -f ./docker-compose/docker-compose-all.yml pull \
  && docker-compose -f ./docker-compose/docker-compose-all.yml up -d

