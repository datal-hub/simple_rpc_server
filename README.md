# simple_rpc_server

### Initialization

This example use [reform](https://github.com/go-reform/reform) and work with database PostgreSQL.
First of all you need setup postgres and create database.
Then define connection parameters in configuration file and then initialization database by command below.
```sh
./rpc-server -vv -c rpc_server.ini adm initdb
```
### Start

For the start server run command:
```sh
./rpc-server -vv -c rpc_server.ini srv
```
