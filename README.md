# cli
Command Line Interface for waitfor

## Installation

### Binary

You can download the latest binaries from [here](https://github.com/go-waitfor/cli/releases).

### Source

#### Go < 1.17

```shell
go get -u github.com/go-waitfor/cli/waitfor@latest
```

#### Go >= 1.17
```shell
go install github.com/go-waitfor/cli/waitfor@latest
```

### SSH
```shell
curl https://raw.githubusercontent.com/go-waitfor/waitfor/master/install.sh | sudo sh
```

## Quick start
```bash
    waitfor -r postgres://locahost:5432/mydb?user=user&password=test -r http://myservice:8080 npm start
```

## Options
```bash
NAME:
   waitfor - Tests and waits on the availability of a remote resource

USAGE:
   waitfor [global options] command [command options] [arguments...]

DESCRIPTION:
   Tests and waits on the availability of a remote resource before executing a command with exponential backoff

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --resource value, -r value  -r http://localhost:8080 [$WAITFOR_RESOURCE]
   --attempts value, -a value  amount of attempts (default: 5) [$WAITFOR_ATTEMPTS]
   --interval value            interval between attempts (sec) (default: 5) [$WAITFOR_INTERVAL]
   --max-interval value        maximum interval between attempts (sec) (default: 60) [$WAITFOR_MAX_INTERVAL]
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)

```
