// The MIT License (MIT)
//
// # Copyright (c) 2016 xtaci
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package std

import (
	"io"

	"github.com/xtaci/qpp"
)

// QPPPort implements io.ReadWriteCloser interface for Quantum Permutation Pads
type QPPPort struct {
	underlying io.ReadWriteCloser // io.Writer is not enough, we need to close the underlying writer as well

	pad   *qpp.QuantumPermutationPad
	wprng *qpp.Rand
	rprng *qpp.Rand
}

func NewQPPPort(underlying io.ReadWriteCloser, pad *qpp.QuantumPermutationPad, seed []byte) *QPPPort {
	wprng := qpp.CreatePRNG(seed)
	rprng := qpp.CreatePRNG(seed)
	return &QPPPort{underlying, pad, wprng, rprng}
}

func (r *QPPPort) Read(p []byte) (n int, err error) {
	n, err = r.underlying.Read(p)
	r.pad.DecryptWithPRNG(p[:n], r.rprng)
	return
}

func (r *QPPPort) Write(p []byte) (n int, err error) {
	r.pad.EncryptWithPRNG(p, r.wprng)
	return r.underlying.Write(p)
}

func (r *QPPPort) Close() error {
	return r.underlying.Close()
}
