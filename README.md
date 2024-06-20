# jwt-with-rbac
Implementation of simple JWT Token With Role Based Access Control

how to start??

### copy .env.example file
```bash
$ cp .env.example .env
```

### fill the .env file using auto generated pem-keys
```bash
$ make
```

## COPY YOUR KEY TO .env

### don't forget to clean up your mess using
```bash
$ make clean
```

### run the server
```bash
$ make start
```