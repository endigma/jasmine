build:
	@go build -ldflags "-s -w" -gcflags="-trimpath=${HOME}" -asmflags="-trimpath=${HOME}" -o out/
