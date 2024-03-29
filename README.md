# advent-of-code-2023

My solutions to Advent of Code 2023, attempting to learn Go (from a minimal base).
Using Go version `go1.21.4`. 

![check code workflow](https://github.com/itsluketwist/advent-of-code-2023/actions/workflows/check.yaml/badge.svg)

<div>
    <!-- badges from : https://shields.io/ -->
    <!-- logos available : https://simpleicons.org/ -->
    <a href="https://adventofcode.com/2023/">
        <img alt="Advent of Code" src="https://img.shields.io/badge/Advent_of_Code-FFFF66?style=for-the-badge&logo=adventofcode&logoColor=black" />
    </a>
    <a href="https://go.dev/">
        <img alt="Go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
    </a>
</div>


## *notable days*
- [Day 8](day08/main.go) - GCD / LCM solution
- [Day 10](day10/main.go) / [Day 18](day18/main.go) - uses Shoelace formula and Pick's theorem
- [Day 12](day12/main.go) - dynamic programming example
- [Day 12](day12/main.go) - uses Djikstra's algorithm and a bucket heap implementation
- [Day 20](day20/main.go) - FIFO queue
- [Day 23](day23/main.go) - graph building + traversal

## *install*

Make sure you have go installed:

```shell
go version
```

Otherwise... [go get it](https://go.dev/doc/install)!

## *usage*

To run the code for day `x`, run:

```shell
make run d=x
```

You can also specify a day part, or whether to try the final input:

```shell
make run d=x p=1 t=1
```

To test the code for day `x`, run:

```shell
make test day=x
```

Format and lint the code with:

```shell
make clean
```

You can also create the template folder and code for day `x` using:

```shell
make new d=x
```
