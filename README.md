# tcid
Go Turkish Identity Number Generator &amp; Validator

## Usage

go get github.com/itp-atakan/tcid

```go
import "github.com/itp-atakan/tcid"

func main() {
    id := tcid.Generate()
    if !tcid.Validate(id) {
        log.Errorf("ID is invalid %s", id)
    }
}
```