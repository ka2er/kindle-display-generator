# kindle-display-generator
 Simple GO script to fetch some info and output a png suitable for kindle display

## Prerequisite

- an old Kindle
- a server to run a cron job (this program)
- a kindle cron to fetch image that this job output

## Dependency

https://github.com/fogleman/gg
go install github.com/fogleman/gg@latest

## Setup

Just copy config.json.template to config.json, edit your settings and run !

````
cp config.json.template config.json
... edit config.json ...
go install github.com/fogleman/gg@latest
go run main.go
````