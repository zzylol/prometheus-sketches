go test -v  -count=1  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=10000
go test -v  -count=1  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=100000
