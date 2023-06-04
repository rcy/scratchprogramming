start:
	find . -name *.html -o -name *.go | entr -r go run .
