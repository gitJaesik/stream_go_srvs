FROM ubuntu:18.04

# RUN cd /etc/apt && \
#   sed -i 's,archive.ubuntu.com/ubuntu,repo.net/artifactory/ubuntu-archive-remote,g' sources.list && \
#   sed -i 's,security.ubuntu.com/ubuntu,repo.net/artifactory/ubuntu-security-remote,g' sources.list

# Base
RUN apt-get update && apt-get install -y --no-install-recommends apt-utils iputils-ping

COPY ./dist/dice-api-server/dice-api-server /usr/local/bin
# data for token
# COPY ./data /usr/local/bin
RUN mkdir -p /log
WORKDIR /usr/local/bin
ENTRYPOINT ./dice-api-server

# EXPOSE 8280
# EXPOSE 8281
# EXPOSE 8282
