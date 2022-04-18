build-blueprint-linter:
	go build -o bin/gc_linter
	zip -j gc_linter.zip ./bin/gc_linter ./blueprint.rule.json

run:
	go run .

test:
	go test -v ./...
