go-ean
======

go-ean is a simple utility library for calculating EAN checksums and validating EAN-8, EAN-13 and UPC numbers.

## Installation

    go get github.com/nicholassm/go-ean/ean

## Usage

To calculate a checksum use the `ChecksumEan8`, `ChecksumEan13` or `ChecksumUpc` functions:

    package main

    import "github.com/nicholassm/go-ean/ean"

    func main() {
      c, err := ean.ChecksumUpc("012345678905")
      println(c)  // Prints 5

      c, err := ean.ChecksumEan13("629104150021")
      println(c)  // Prints 3

      c2, err2 := ean.ChecksumEan8("7351353")
      println(c2) // Prints 7
    }

To check the validity of a string as EAN-8, EAN-13 or UPC use `Valid`:

    println(ean.Valid("96385074")) // Prints true
    println(ean.Valid("abc"))      // Prints false

## License

(The MIT License)

Copyright (c) 2013 Nicholas Schultz-Møller

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
