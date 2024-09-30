go test -v  -count=2  -timeout 0 -run TestAppendThroughput ./ -numts=100
go test -v  -count=2  -timeout 0 -run TestAppendThroughput ./ -numts=1000
go test -v  -count=2  -timeout 0 -run TestAppendThroughput ./ -numts=10000
