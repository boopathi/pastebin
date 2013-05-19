all:
	git submodule init
	git submodule update

serve:
	cd src/pb && go install
