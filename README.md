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

## Build

If you need to cross compile to a linux target system

````
GOOS=linux GOARCH=amd64 go build -o main-linux main.go
````

## Deploy

You could file *inventory* file with IP of the server and the Kindle.
Just check that you are able to ssh to them.
Then you juste have to do a : 

`````
deploy.sh
`````

# Credits

https://rentafounder.com/convert-image-to-grayscale-golang/
