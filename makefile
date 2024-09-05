cli:
	go build -o cli ./cmd/cli

gui:
	go build -o gui ./cmd/gui

clean:
	rm cli gui
