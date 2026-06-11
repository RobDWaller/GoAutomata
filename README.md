# Go Automata

![Go Version](https://img.shields.io/badge/Go-1.24.4-00ADD8?logo=go)

Go Automata is a small Go library for working with **Wolfram elementary cellular automata**.
It converts a rule number (`0-255`) into an 8-bit representation and a neighborhood lookup table.

In this model, each cell's next state depends on a 3-cell neighborhood:

- `111`
- `110`
- `101`
- `100`
- `011`
- `010`
- `001`
- `000`

## Usage

```go
import "github.com/robdwaller/go-automata/automata"

steps := uint(5)
rule := uint8(30)
seed := "0000000000000000000100000000000000000000"

rows := automata.Generate(steps, rule, seed)

// rows[0] is the seed row, rows[steps] is the final generation
for step := uint(0); step <= steps; step++ {
    fmt.Println(rows[step])
}
```

You can also work with the lower-level API:

```go
engines := automata.MakeRuleEngine(30)
// enginge is a []RuleEngine — one entry per neighbourhood pattern

nextRow := automata.Step(engines, []uint8{1, 0, 0})
// nextRow = []uint8{1, 1, 1}
```

Bit mapping follows Wolfram order from left to right (`111` to `000`).
For rule 30 (`00011110`):

- `111 -> 0`
- `110 -> 0`
- `101 -> 0`
- `100 -> 1`
- `011 -> 1`
- `010 -> 1`
- `001 -> 1`
- `000 -> 0`

## Architecture

All source files live in the `automata/` package:

- `automata/automata.go` — `Generate()` convenience wrapper that accepts a seed string
- `automata/rules.go` — `RuleEngine` struct, `GetRuleset()`, `RuleToBinary()`, `MakeRuleEngine()`
- `automata/neighbourhood.go` — `GetNeighbourhood()`, `FindNeighbourhood()`
- `automata/steps.go` — `Step()` (advance one row), `Steps()` (advance multiple rows)

## Troubleshooting

- Rule values are `uint8`, so valid rule numbers are `0-255`
- Binary slices are always 8 bits (most significant bit first)
- Neighbourhoods are compared struct-by-struct; the wildcard patterns are the 8 Wolfram-order triples

## Development

Run tests:

```bash
go test -v ./...
```

Validate code:

```bash
gofmt -l ./
go vet ./...
golangci-lint run
golangci-lint run --output.html.path [file-or-directory]
```

## License

MIT — see [LICENSE](LICENSE) for details.
