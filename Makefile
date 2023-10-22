command_name = boko
command_source = cmd/main/boko.go
install_path = $(GOBIN)/$(command_name)
build_path = bin/$(command_name)
home_path = ~/.boko

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
	mkdir -p $(home_path)/icons
	cp resources/default.png $(home_path)/icons/default.png
	ln -nfs $(install_path) /usr/local/bin/boko

$(build_path) : $(wildcard pkg/*) $(command_source)
	go build -o $(build_path) $(command_source)

.PHONY : setup
setup :
	go install github.com/golang/mock/mockgen

.PHONY : uninstall
uninstall: install
	rm $(install_path)
	rm /usr/local/bin/boko
