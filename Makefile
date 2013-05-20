all:
	git submodule init
	git submodule update

serve:
	cd src/pb && go install
	./utils/pastebin

start:
	./utils/pastebin start

stop:
	./utils/pastebin stop

restart:
	./utils/pastebin restart
