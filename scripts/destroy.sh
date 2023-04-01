#!/bin/bash
echo "Destroying dev server"

docker-compose down
rm tmp -rf