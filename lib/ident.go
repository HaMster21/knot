// Copyright © 2015 Hans Meyer <hamster.of.dev@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lib

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
	"io"
)

type KObject struct {
	Identifier [64]byte
	Data       []byte
}

func NewKObject(data io.Reader) *KObject {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(data); err != nil {
		panic(err)
	}

	var h [64]byte
	sha3.ShakeSum256(h[:], buf.Bytes())
	return &KObject{Identifier: h, Data: buf.Bytes()}
}

func (k *KObject) String() string {
	return fmt.Sprintf("%s\n\n%s", base58.Encode(k.Identifier[:]), k.Data)
}
