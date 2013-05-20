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

+ `git clone git://github.com/boopathi/pastebin.git`
+ `cd pastebin`
+ `source set_env.sh` #Sets the environment
+ `make && make install`
+ `./bin/pb` (or) just `pb`, as `./bin/` will be in your `$PATH`

## Contributing

+ Clone the repository, make changes and `make install && ./bin/pb` to re-compile the source and run the server
