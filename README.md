# package_indexer

## Description
A basic package indexer written in Go. It keeps track of package dependencies. Clients connect to the server and inform which packages should be indexed, and which dependencies they might have on other packages. The server does not index any package until all of its dependencies have been indexed first. The server also does not remove a package if any other packages depend on it.

The server opens a TCP socket on port 8080. It accepts connections from multiple clients at the same time, all trying to add and remove items to the index concurrently. Clients are independent of each other, and it is expected that they will send repeated or contradicting messages. New clients can connect and disconnect at any moment, and sometimes clients can behave badly and try to send broken messages.

Messages from clients follow this pattern: `<command>|<package>|<dependencies>\n`

Where:
* `<command>` is mandatory, and is either `INDEX`, `REMOVE`, or `QUERY`
* `<package>` is mandatory, the name of the package referred to by the command, e.g. `mysql`, `openssl`, `pkg-config`, `postgresql`, etc.
* `<dependencies>` is optional, and if present it will be a comma-delimited list of packages that need to be present before `<package>` is installed. e.g. `cmake,sphinx-doc,xz`
* The message always ends with the character `\n`

Here are some sample messages:
```
INDEX|cloog|gmp,isl,pkg-config\n
INDEX|ceylon|\n
REMOVE|cloog|\n
QUERY|cloog|\n
```

For each message sent, the client will wait for a response code from the server. Possible response codes are `OK\n`, `FAIL\n`, or `ERROR\n`. After receiving the response code, the client can send more messages.

The response code returned should be as follows:
* For `INDEX` commands, the server returns `OK\n` if the package can be indexed. It returns `FAIL\n` if the package cannot be indexed because some of its dependencies aren't indexed yet and need to be installed first. If a package already exists, then its list of dependencies is updated to the one provided with the latest command.
* For `REMOVE` commands, the server returns `OK\n` if the package could be removed from the index. It returns `FAIL\n` if the package could not be removed from the index because some other indexed package depends on it. It returns `OK\n` if the package wasn't indexed.
* For `QUERY` commands, the server returns `OK\n` if the package is indexed. It returns `FAIL\n` if the package isn't indexed.
* If the server doesn't recognize the command or if there's any problem with the message sent by the client it should return `ERROR\n`.

## Built With 
* Golang: I chose to use Golang because the last project I coded in was written in Go as well. Go is also great for and easy to use when testing. 
* Docker: A Dockerfile is used to build a Docker image. 

## How to Run 
1. To build an executable, run 
```
$ go build 
```
Then to run the executable/server, run 
```
$ ./package_indexer
```

2. Alternatively, you can create a Docker image using the Dockerfile. First, create the docker image using
```
$ docker build -t package-indexer . 
```
Then to run, use
```
$ docker run -p 8080:8080 package-indexer
```

3. Or, pull and run the docker image from my repository. This can be done by doing
```
$ docker run -p 8080:8080 juleskim/package-indexer:latest
```

To run the tests, run
```
$ go test
```
To run the tests and see its coverage and to see more information about tests run, do 
```
$ go test -cover -v
```

## Testing 
Tests are included in the `*_test.go` files. These tests cover parsing and indexing requests. See How to Run section for how to run the tests. 
The given test suite was also running during development. All tests passed. 

## Security 
* Format string vulnerabilities: If a format string were not given to a function dealing with format strings, this could lead to format string vulnerabilities. Thus, all such calls to the function were given the proper amount of format strings and arguments. 