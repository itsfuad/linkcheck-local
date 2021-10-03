# linkcheck-local

Locate relative Markdown links for offline document validation.

Small, dependency-light Go standard library utility. Original implementation and history generated for a retrospective development exercise.

## Run

`go run . examples/input.txt --pretty`

Use `go run . -` for piped input. The program does not modify input files.

## Maintenance

Version 0.2 includes a small usability improvement and additional coverage for everyday workflows. No network dependencies are required.
