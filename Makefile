.DEFAULT_GOAL := run

d := -1  # day to run
p := 0  # part to run
t := 0  # try input

# run code
run:
ifeq ($(shell expr $(d) \>= 10), 1)
	go run github.com/itsluketwist/advent-of-code-2023/day$(d) -part=$(p) -try=$(t)
else ifeq ($(shell expr $(d) \>= 0), 1)
	go run github.com/itsluketwist/advent-of-code-2023/day0$(d) -part=$(p) -try=$(t)
endif

# run tests
test:
ifeq ($(shell expr $(d) \>= 10), 1)
	go test -v github.com/itsluketwist/advent-of-code-2023/day$(n)
else ifeq ($(shell expr $(d) \>= 0), 1)
	go test -v github.com/itsluketwist/advent-of-code-2023/day0$(n)
else
	go test -v ./...
endif

# format code
format:
	go fmt ./...

# lint code
lint:
	go vet ./...

# do both!
clean:
	make format lint
