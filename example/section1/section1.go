// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"

	formatalign "github.com/coolstina/format-align"
)

func main() {
	align := formatalign.NewFormatAlign()
	buffer := bytes.Buffer{}

	buffer.Write(append([]byte(align.Format("Username", "helloshaohua")), '\n'))
	buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
	buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

	fmt.Println(buffer.String())
}
