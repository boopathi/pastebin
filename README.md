# pastebin

Hacker's Pastebin (curl to post your pastes)

## Usage

### Random URI

+ Send a `POST` or `PUT` request to paste.
+ `cat SomeFile | curl -X PUT --data-urlencode paste@- localhost:8080`
+ `curl -X PUT --data-urlencode paste@myfile.txt localhost:8080`
+ `iostat | curl -X PUT --data-urlencode paste@- localhost:8080`

### Custom URI

+ `-d name=customURI` is to be included
+ `cat somefile | curl -X PUT --data-urlencode paste@- -d name=customURI localhost:8080`

## Installation

+ `go get github.com/boopathi/pastebin`
+ `cd $GOPATH/src/github.com/boopathi/pastebin`
+ `go build`
+ Start your redis-server on `localhost:6379`
+ Run `./pastebin`

## LICENSE

http://boopathi.mit-license.org/
