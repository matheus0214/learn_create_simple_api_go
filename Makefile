list-dependencies:
	go list -f '{{ join .Imports "\n" }}'