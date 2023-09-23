command_name = boko
command_source = cmd/main/boko.go
install_path = $(GOBIN)/$(command_name)
build_path = bin/$(command_name)

.PHONY : build
build : $(build_path)

.PHONY : generate
generate :
	go generate ./...

.PHONY : test
test : generate $(wildcard pkg/*)
	go test ./... --cover

.PHONY : install
install : $(install_path)

$(install_path) : $(wildcard pkg/*) $(command_source)
	go install $(command_source)

$(build_path) : $(wildcard pkg/*) $(command_source)
	go build -o $(build_path) $(command_source)
