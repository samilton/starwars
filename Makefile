.PHONY: all xwing deathstar clean

all: xwing deathstar

xwing:
	go build -o services/xwing/xwing ./services/xwing/main.go

deathstar:
	go build -o services/deathstar/deathstar ./services/deathstar/main.go

clean:
	rm -f services/xwing/xwing services/deathstar/deathstar