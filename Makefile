all: go-libs
	clang -std=gnu11 c/main.c -Ic/include $(LDFLAGS) -L./out -ltest -o out/main

.PHONY: dirs
dirs:
	mkdir -p out

.PHONY: go-dyn
go-dyn:
	go install -buildmode=shared -linkshared -v -a -x std
	@cd go && go install -buildmode=shared -linkshared -v test

go-libs: dirs
	cd go && go build -o ../out/libtest.so --buildmode=c-shared -linkshared test
