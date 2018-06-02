#
# Simple Makefile for conviently testing, building and deploying experiment.
#
PROJECT = scriptreplayer

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif


scriptreplayer$(EXT): bin/scriptreplayer$(EXT)


bin/scriptreplayer$(EXT): scriptreplayer.go
	go build -o bin/scriptreplayer$(EXT) scriptreplayer.go

build: $(PROJECT)$(EXT)

install: 
	env GOBIN=$(GOPATH)/bin go install scriptreplayer.go

website: page.tmpl README.md nav.md INSTALL.md LICENSE css/site.css
	bash mk-website.bash

#test: clean bin/scriptreplayer$(EXT)
#	go test

format:
	gofmt -w scriptreplayer.go

lint:
	golint scriptreplayer.go

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/scriptreplayer scriptreplayer.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/scriptreplayer.exe scriptreplayer.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/scriptreplayer scriptreplayer.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/scriptreplayer scriptreplayer.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	if [ -d dist ]; then rm -fR dist; fi
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	bash package-versions.bash > dist/package-versions.txt

update_version:
	./update_version.py --yes

release: clean scriptreplayer.go distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7
	

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish:
	bash mk-website.bash
	bash publish.bash

