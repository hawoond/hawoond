module github.com/hawoond/hawoond

go 1.22.3

replace (
	github.com/hawoond/hawoond/cmd/utils/converter => ./cmd/utils/converter
	github.com/hawoond/hawoond/cmd/utils/parser/json => ./cmd/utils/parser/json
	github.com/hawoond/hawoond/cmd/worker => ./cmd/worker
	github.com/hawoond/hawoond/internal/common/errorkit => ./internal/common/errorkit
	github.com/hawoond/hawoond/internal/common/quote => ./internal/common/quote
	github.com/hawoond/hawoond/internal/generator => ./internal/generator
)
