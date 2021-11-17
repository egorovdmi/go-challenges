default:
	go build -o ./bin/speller ./cmd/speller
	go build -o ./bin/speller-trie ./cmd/speller-trie
	go build -o ./bin/string-builder ./cmd/string-builder
	go build -o ./bin/check-permutations ./cmd/check-permutations
	go build -o ./bin/urlify ./cmd/urlify

unit:
	go test ./... -v

bench:
	go test -bench . ./...
