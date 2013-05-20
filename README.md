# pastebin

Hacker's Pastebin (curl to post your pastes)

## Usage

+ `cat SomeFile | curl -X POST --data-urlencode paste@- localhost:8080`
+ `curl -X POST --data-urlencode paste@myfile.txt localhost:8080`
+ `iostat | curl -X POST --data-urlencode paste@- localhost:8080`

## Installation

+ `git clone git://github.com/boopathi/pastebin.git`
+ `cd pastebin`
+ `source set_env.sh` #Sets the environment
+ `make && make serve`
+ `./bin/pb` (or) just `pb`, as `./bin/` will be in your `$PATH`

## Contributing

+ Clone the repository, make changes and `make serve && ./bin/pb` to re-compile the source and run the server
