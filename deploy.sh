#!/bin/sh

ansible-playbook -i inventory kindle/kindle.yml
ansible-playbook -i inventory server/server.yml
