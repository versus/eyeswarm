#!/bin/bash

sudo apt update && apt upgrade -y && curl -fsSL get.docker.com -o get-docker.sh && chmod 0755 get-docker.sh && sudo get-docker.sh && sudo usermod -aG docker vagrant

