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

package cmd

import (
	"bytes"
	"fmt"
	"github.com/hamster21/knot/lib"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// identCmd represents the ident command
var identCmd = &cobra.Command{
	Use:   "ident",
	Short: "Creates an identification string for the provided file",
	Long: `I'm going to fill this in later :)
`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := range args {
			buf, err := ioutil.ReadFile(args[i])
			if err != nil {
				panic(err)
			} else {
				fmt.Println(lib.NewKObject(bytes.NewBuffer(buf)))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(identCmd)
}
