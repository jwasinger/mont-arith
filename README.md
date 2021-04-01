# Go-evm-modext

Pure-go arithmetic generator for modular addition subtraction and Montgomery multiplication across a range of bitwidths (256-4096bit)

Implementation of Montgomery multiplication for bitwidths between 256-768bit are adapted slightly from [Goff](https://github.com/consensys/goff).  Widths above 768bits use generic/non-interleaved Montgomery multiplication using Go's built-in big.Int.  This is a stop-gap until more performant implementations are created.

## Usage

Generate the arithmetic (placed under `build` directory) and a user-friendly interface for invoking it:
`make`

Run unit tests:
`make test`

Benchmarks
`make benchmark`
