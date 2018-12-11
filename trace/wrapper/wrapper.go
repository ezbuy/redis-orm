package wrapper

import (
	"context"
)

type Wrapper interface {
	Do(ctx context.Context, statement string)
	Close()
}
