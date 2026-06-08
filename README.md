# Go Automata

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
rule := uint8(30)

binary := rules.RuleToBinary(rule)
// "00011110"

engine, err := rules.MakeRuleEngine(rule)
if err != nil {
    panic(err)
}

next := engine["100"]
// 1
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

- `rules/ruleset.go`: Returns the 8 neighborhood patterns in Wolfram order
- `rules/binaries.go`: Converts rule numbers to 8-bit binary strings
- `rules/rule_engine.go`: Builds neighborhood-to-next-state lookup maps

## Troubleshooting

- Rule values are `uint8`, so valid rule numbers are `0-255`
- Binary strings are always 8 bits (left-padded with zeros)
- If engine construction fails, check for invalid binary parsing or local code modifications

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
