# cli
Command Line Interface for waitfor

**waitfor** is a utility that tests and waits for the availability of remote resources before executing a command. It supports various resource types including HTTP endpoints, databases (PostgreSQL, MySQL, MongoDB), filesystem paths, and processes. It uses exponential backoff to retry connections, making it perfect for startup scripts, CI/CD pipelines, and container orchestration scenarios.

## Use Cases

- **Container Orchestration**: Wait for dependent services before starting your application
- **CI/CD Pipelines**: Ensure databases and services are ready before running tests
- **Development**: Start services in the correct order during local development
- **Health Checks**: Verify service availability before proceeding with deployments

## Supported Resource Types

waitfor supports the following resource types:

- **HTTP/HTTPS** - `http://` or `https://` - Tests HTTP endpoint availability
- **PostgreSQL** - `postgres://` - Tests PostgreSQL database connectivity
- **MySQL** - `mysql://` - Tests MySQL database connectivity  
- **MongoDB** - `mongodb://` - Tests MongoDB database connectivity
- **Filesystem** - `file://` - Tests file or directory existence
- **Process** - `proc://` - Tests if a process is running

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

### Install Script
```shell
curl https://raw.githubusercontent.com/go-waitfor/cli/master/install.sh | sudo sh
```

## Quick start

### Basic Usage
```bash
# Wait for a PostgreSQL database and HTTP service, then start npm
waitfor -r postgres://localhost:5432/mydb?user=user&password=test -r http://myservice:8080 npm start
```

### Examples by Resource Type

#### HTTP Endpoints
```bash
# Wait for a web service to be available
waitfor -r http://localhost:8080 npm start

# Wait for multiple services
waitfor -r http://localhost:8080 -r http://localhost:3000 docker-compose up
```

#### Database Connections
```bash
# PostgreSQL
waitfor -r postgres://user:password@localhost:5432/dbname ./start-app.sh

# MySQL
waitfor -r mysql://user:password@localhost:3306/dbname ./start-app.sh

# MongoDB
waitfor -r mongodb://localhost:27017/dbname ./start-app.sh
```

#### Filesystem
```bash
# Wait for a file to exist
waitfor -r file:///path/to/file ./start-app.sh

# Wait for a directory
waitfor -r file:///path/to/directory ./start-app.sh
```

#### Process
```bash
# Wait for a process to be running
waitfor -r proc://nginx ./start-app.sh
```

### Environment Variables
You can use environment variables instead of command-line flags:
```bash
export WAITFOR_RESOURCE="http://localhost:8080,postgres://localhost:5432/mydb"
export WAITFOR_ATTEMPTS=10
export WAITFOR_INTERVAL=3
export WAITFOR_MAX_INTERVAL=30

waitfor npm start
```

## Configuration

### Command Line Options

| Option | Short | Description | Default | Environment Variable |
|--------|-------|-------------|---------|---------------------|
| `--resource` | `-r` | Resource to wait for (can be specified multiple times) | - | `WAITFOR_RESOURCE` |
| `--attempts` | `-a` | Number of connection attempts | `5` | `WAITFOR_ATTEMPTS` |
| `--interval` | - | Initial interval between attempts (seconds) | `5` | `WAITFOR_INTERVAL` |
| `--max-interval` | - | Maximum interval between attempts (seconds) | `60` | `WAITFOR_MAX_INTERVAL` |
| `--help` | `-h` | Show help | - | - |
| `--version` | `-v` | Show version | - | - |

### Full Command Reference
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
