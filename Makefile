OUTFILE = bin/quorum-challenge-backend

all: staticcheck build

fmtcheck:
	"$(CURDIR)/scripts/gofmtcheck.sh"

staticcheck:
	"$(CURDIR)/scripts/staticcheck.sh"

build:
	"$(CURDIR)/scripts/build.sh" -o $(OUTFILE)

clean:
	"$(CURDIR)/scripts/build.sh" -c

run:
	./$(OUTFILE)

out:
	echo "$(OUTFILE)"

test:
	go test -p 1 -v ./... 

debug: build run

.PHONY: *
