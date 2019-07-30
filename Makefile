null  :=
space := $(null) #
comma := ,

PKGSLIST = mx mx/internal
COVERPKGS= $(subst $(space),$(comma),$(strip $(foreach i,$(PKGSLIST),github.com/sudachen/go-mxnet/$(i))))

build:
	cd mx; go build

run-tests:
	mkdir -p .tmp/artifacts
	cd tests && go test -coverprofile=../.tmp/c.out -coverpkg=$(COVERPKGS)
	go tool cover -html=.tmp/c.out -o .tmp/artifacts/coverage.html

