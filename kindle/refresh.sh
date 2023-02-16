#!/bin/sh

curl http://<server>/out.png -o /mnt/us/out.png
eips -c
eips -g /mnt/us/out.png
