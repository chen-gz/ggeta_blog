.PHONY: all front_end back_end

all: front_end back_end

front_end:
	cd front_end && yarn build

back_end:
	GOOS=linux GOARCH=amd64 go build

