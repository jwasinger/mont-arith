# mont-arith 

Library for Modular Addition, Subtraction and Montgomery Multiplication.  Supports values between 0 and 16384 bits.

Implementation of Montgomery multiplication for bitwidths between 256-768bit are adapted slightly from [Goff](https://github.com/consensys/goff).  Widths above 768bits use generic/non-interleaved Montgomery multiplication using Go's built-in big.Int.  This is a stop-gap until more performant implementations are created.

## Usage

Run unit tests:
`make test`

Benchmarks
`make benchmark`
