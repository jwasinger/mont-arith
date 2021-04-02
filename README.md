# mont-arith 

Library for Modular Addition, Subtraction and Montgomery Multiplication.  Supports values between 1 and 16384 bits.

Implementation of the arithmetic up to 704 bits is adapted slightly from [Goff](https://github.com/consensys/goff).  Widths above 704 bits do generic/non-interleaved Montgomery multiplication using Go's built-in big.Int.  This is a stop-gap until more performant implementations are created.

## Usage

Generate Go code for the arithmetic:
```
make
```

Unit tests:
```
make test
```

Benchmarks:
```
make benchmark
```
