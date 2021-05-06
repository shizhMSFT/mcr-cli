.PHONY: all
all: build

.PHONY: FORCE
FORCE:

bin/mcr: cmd/mcr FORCE
	go build -o $@ ./$<

.PHONY: build
build: bin/mcr

.PHONY: clean
clean:
	git status --ignored --short | grep '^!! ' | sed 's/!! //' | xargs rm -rf

.PHONY: install
install: bin/mcr
	cp $< ~/bin/
