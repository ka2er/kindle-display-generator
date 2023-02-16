#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o main-linux main.go
ansible-playbook -i inventory kindle/kindle.yml
ansible-playbook -i inventory server/server.yml
