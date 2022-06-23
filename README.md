# format-align

Golang string format align implement

## Installation

```shell
go get -u github.com/coolstina/format-align
```

## Feature

- Alignment: Support Right/Left align
- Placeholder: Support specify placeholder length
- Separator: Support specify separator

## Example

### [Use default argument.](example/section1/section1.go)

Use default argument.

```shell
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
```

```shell
Username                  : helloshaohua
Sex                       : male
IsStudent                 : false
```


### [Alignment of right.](example/section2/section2.go)

```go
package main

import (
	"bytes"
	"fmt"

	formatalign "github.com/coolstina/format-align"
)

func main() {
	align := formatalign.NewFormatAlign(formatalign.WithAlignment(formatalign.AlignmentOfRight))
	buffer := bytes.Buffer{}

	buffer.Write(append([]byte(align.Format("Username", "helloshaohua")), '\n'))
	buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
	buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

	fmt.Println(buffer.String())
}
```

```shell
                  Username: helloshaohua
                       Sex: male
                 IsStudent: false
```


### [Specify placeholder.](example/section3/section3.go)


```go
package main

import (
	"bytes"
	"fmt"

	formatalign "github.com/coolstina/format-align"
)

func main() {
	align := formatalign.NewFormatAlign(
		formatalign.WithAlignment(formatalign.AlignmentOfRight),
		formatalign.WithPlaceholder(12),
	)
	buffer := bytes.Buffer{}

	buffer.Write(append([]byte(align.Format("Username", "helloshaohua")), '\n'))
	buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
	buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

	fmt.Println(buffer.String())
}
```

```shell
    Username: helloshaohua
         Sex: male
   IsStudent: false
```

### [Specify separator.](example/section3/section3.go)

```go
package main

import (
	"bytes"
	"fmt"

	formatalign "github.com/coolstina/format-align"
)

func main() {
	align := formatalign.NewFormatAlign(
		formatalign.WithAlignment(formatalign.AlignmentOfRight),
		formatalign.WithPlaceholder(12),
		formatalign.WithSeparator("-"),
	)
	buffer := bytes.Buffer{}

	buffer.Write(append([]byte(align.Format("Username", "helloshaohua")), '\n'))
	buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
	buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

	fmt.Println(buffer.String())
}
```

```shell
    Username- helloshaohua
         Sex- male
   IsStudent- false
```