build:
	go build -o devops-tools ./cmd/devops-tools
run:
	./devops-tools

clean:
	rm -f devops-tools