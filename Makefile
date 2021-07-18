all: go-libs
	clang -fPIC -std=gnu11 c/main.c -Ic/include $(LDFLAGS) ./out/libgsego.so -o out/gsec

.PHONY: dirs
dirs:
	mkdir -p out

.PHONY: go-dyn
go-dyn:
	go install -buildmode=shared -linkshared -v -a -x std

go-libs: dirs
	cd go && go build -o ../out/libgsego.so --buildmode=c-shared -v gsego
