pastebin
========

Hacker's Pastebin (curl to post your pastes)

Usage
=====

+ `cat SomeFile | curl -X POST --data-urlencode paste@- localhost:8080`
+ `curl -X POST --data-urlencode paste@myfile.txt localhost:8080`
+ `iostat | curl -X POST --data-urlencode paste@- localhost:8080`

