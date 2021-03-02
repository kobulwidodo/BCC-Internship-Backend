#!/bin/bash

cd /var/app/current/;

export HOME=/root;
export GOCACHE=/root/builds;

go run main.go;
