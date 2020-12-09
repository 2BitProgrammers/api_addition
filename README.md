# 2bitprogrammers/api_addition

This is a simple API example which sums two values (a and b).  It is meant to be used for instructional purposes only.

The API listens on port:  1234

**<u>API Enpoints</u>**:
* **/status** (GET) - this states whether the app is up and healthy
* **/add** (POST) - this returns the sum of the provided variables. This requires a json body to be included with the request.
  * <u>Example JSON Body:</u>   _{ "a": 11, "b": 22 }_


## Run as Standalone GoLang App
This will run the application with the go application.  It assumes that you have installed your golang environment correctly.

```bash
$ cd src
$ go run main.go

2bitprogrammers/api_addition v2007.21a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2020/12/08 21:02:05 Starting App on Port 1234

CTRL+C
```


## Run within Docker 
This will run the components on your local system without using minikube or kubernetes.

### Building the Docker Image
```bash
$ docker build . -t 2bitprogrammers/api_addition

Sending build context to Docker daemon  16.38kB
Step 1/11 : FROM golang:alpine AS builder
 ---> b3bc898ad092
Step 2/11 : ENV GO111MODULE=on     CGO_ENABLED=0     GOOS=linux     GOARCH=amd64
 ---> Using cache
 ---> 8462443c0070
Step 3/11 : WORKDIR /build
 ---> Using cache
 ---> 99600623930c
Step 4/11 : COPY $PWD/src/go.mod .
 ---> Using cache
 ---> e88906c713f8
Step 5/11 : COPY $PWD/src/main.go .
 ---> Using cache
 ---> d63b1c87b2ae
Step 6/11 : RUN go mod download
 ---> Using cache
 ---> 437d19f670ae
Step 7/11 : RUN go build -o api_addition .
 ---> Using cache
 ---> 4f23ea20130b
Step 8/11 : FROM scratch
 --->
Step 9/11 : WORKDIR /
 ---> Using cache
 ---> a66c59ea194a
Step 10/11 : COPY --from=builder /build/api_addition .
 ---> Using cache
 ---> 7c1bad4f3dc7
Step 11/11 : ENTRYPOINT [ "/api_addition" ]
 ---> Running in efc8216a4811
Removing intermediate container efc8216a4811
 ---> 6e118aa42002
Successfully built 6e118aa42002
Successfully tagged 2bitprogrammers/api_addition:latest
SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is recommended to double check and reset permissions for sensitive files and directories.
```

### Image Status
```bash
$ docker images

REPOSITORY                              TAG                                              IMAGE ID            CREATED              SIZE
2bitprogrammers/api_addition            latest                                           6e118aa42002        About a minute ago   6.68MB
```

### Running the Container
```bash
$ docker run --rm --name "api_addition" -p 1234:1234 2bitprogrammers/api_addition 

2bitprogrammers/api_addition v2007.21a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2020/12/08 21:02:05 Starting App on Port 1234

CTRL+C
```

### Check the Container Status (docker)
```bash
$ docker ps

CONTAINER ID    IMAGE                           COMMAND             CREATED              STATUS              PORTS                    NAMES
6d7546f90be1    2bitprogrammers/api_addition    "/api_addition"     About a minute ago   Up About a minute   0.0.0.0:1234->1234/tcp   api_addition
```

### Check API Status (health check):
```bash
$ curl http://localhost:1234/status

{"date":"2020-07-10T03:17:02.9034438-07:00","statusCode":200,"statusText":"OK","data":"{ \"healthy\": true}","errors":"","request":{"uri":"/status","method":"GET","payload":""}}
```

### Watch Container Logs
```bash
$ docker logs -f 2bitprogrammers/api_addition

2bitprogrammers/api_addition v2007.21a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2020/07/10 11:07:22 Starting App on Port 1234
2020/07/10 11:08:22 POST /add 200 { "a": 11, "b": 22 }

CTRL+C
```

### Stopping the Container
```bash
$ docker stop api_addition
```

## Using the API
For the below examples, we will assume the following:
* Server:  locahost (127.0.0.1)
* Bind Port: 1234
* Method: POST
* URI: /add 
* Body Data:   _{ "a": 11, "b": 22 }_

For Linux:
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{ "a": 11, "b": 22 }' http://127.0.0.1:1234/add

{"date":"2020-07-10T03:17:02.9034438-07:00","statusCode":200,"statusText":"OK","data":"{ \"value\": 33 }","errors":"","request":{"uri":"/add","method":"POST","payload":"{\"a\": 11, \"b\": 22 }"}}
```

For Windows (powershell):
```powershell
C:\ curl -X POST -H "Content-Type: application/json" -d "{ """a""": 11, """b""": 22 }" http://127.0.0.1:1234/add

{"date":"2020-07-10T03:17:02.9034438-07:00","statusCode":200,"statusText":"OK","data":"{ \"value\": 33 }","errors":"","request":{"uri":"/add","method":"POST","payload":"{\"a\": 11, \"b\": 22 }"}}
```
