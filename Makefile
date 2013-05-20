all:
	git submodule init
	git submodule update

install:
	cd src/pb && go install
	./utils/pastebin

start:
	./utils/pastebin start

stop:
	./utils/pastebin stop

restart:
	./utils/pastebin restart
