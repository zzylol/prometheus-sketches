go test -v  -count=10  -timeout 0 -run TestAppendThroughput ./ -numts=100
go test -v  -count=10  -timeout 0 -run TestAppendThroughput ./ -numts=1000
go test -v  -count=5  -timeout 0 -run TestAppendThroughput ./ -numts=10000
go test -v  -count=5  -timeout 0 -run TestAppendThroughput ./ -numts=50000
go test -v  -count=5  -timeout 0 -run TestAppendThroughput ./ -numts=75000
go test -v  -count=5  -timeout 0 -run TestAppendThroughput ./ -numts=100000
go test -v  -count=5  -timeout 0 -run TestAppendThroughput ./ -numts=150000
go test -v  -count=5  -timeout 0 -run TestAppendThroughput ./ -numts=200000
go test -v  -count=5  -timeout 0 -run TestAppendThroughput ./ -numts=500000