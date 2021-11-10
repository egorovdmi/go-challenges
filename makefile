default:
	go build -o ./bin/speller ./cmd/speller
	go build -o ./bin/speller-trie ./cmd/speller-trie
	go build -o ./bin/string-builder ./cmd/string-builder
