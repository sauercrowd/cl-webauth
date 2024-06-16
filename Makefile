libwebauthn.so:
	go build -o libwebauthn.so -buildmode=c-shared main.go

.PHONY: clean

clean:
	rm -f libwebauthn.*
