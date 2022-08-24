# jsonfmt
[![.github/workflows/ci.yaml](https://github.com/x0y14/jsonfmt/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/x0y14/jsonfmt/actions/workflows/ci.yaml)

simple json formatter

## require
- go 1.19

## build

clone this repo
```shell
$ git clone https://github.com/x0y14/jsonfmt.git
$ cd jsonfmt
```
go build (To place the executable file in the bin directory)
```shell
$ mkdir -p bin
$ go build -o ./bin/jsonfmt ./cmd/jsonfmt/main.go 
```

## how to use
```shell
$ jsonfmt <flags> <args>
```

### example using ./samples/sample1.json

Indentation is based on four spaces, Print only.
```shell
$ ./bin/jsonfmt -p -i 4 ./samples/sample1.json
```

Indentation is based on four spaces, Save the formatted data as a name.
```shell
$ ./bin/jsonfmt -o ./samples/sample1_formatted.json -i 4 ./samples/sample1.json
```

Indentation is based on four spaces, Overwrite the original file with formatted data
```shell
$ ./bin/jsonfmt -w ./samples/sample1.json 
```

### before & after using ./samples/sample3.json
before
```json
{"name":  "john", "age": 18, "sex": "male", "come-from": "usa", "weight": 61.2, "height": 175.0, "parent": {},"partner": null, "married": false, "children": []}
```
after (./bin/jsonfmt ./samples/sample3.json)
```json
{
  "name": "john", 
  "age": 18, 
  "sex": "male", 
  "come-from": "usa", 
  "weight": 61.2, 
  "height": 175, 
  "parent": {}, 
  "partner": null, 
  "married": false, 
  "children": []
}
```

## flags
### mode
If multiple settings are made, priority is given from the top to the bottom.
- output / o  
boolean,
- overwrite / w  
boolean
- print / p   
  boolean, default mode
### config
- ident / i  
int, default is 2
### help
- help