# freeport
A small CLI tool written in go for quickly determining whether a local port is available for use or not.

## NOTE
NOTE: This is not a professional project and has not been thoroughly tested. I take 
no responsibility for any negative outcomes of using this program.

## Build and Install
To build freeport, you must have [GO](https://go.dev/dl/) installed on your system.
1. Clone this repo `git clone git@github.com:gequalspisquared/freeport.git` and cd into it `cd freeport`
2. Build the executable `go build freeport.go`
3. Install the executable to somewhere in your path, on my Linux system this was done by `sudo cp freeport /usr/local/bin`

## Usage
To determine if a port is available for use or not, simple run `freeport <port>`, where '<port>' is replaced with 
the port number you are interested in.
