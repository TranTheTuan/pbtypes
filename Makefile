.PHONY: build install clean upload

build:
	mkdir -pv build/go
	./regenerate_go.sh
