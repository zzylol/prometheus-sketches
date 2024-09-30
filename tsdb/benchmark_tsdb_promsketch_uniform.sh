rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=1
rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=10 
rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=100
rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=1000 
rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=10000
rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=100000
rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=150000
rm -r /mydata/TestDBInsertionThroughputZipf*
go test -v  -count=5  -timeout 0 -run TestDBInsertionThroughputZipf ./ -numts=200000
rm -r /mydata/TestDBInsertionThroughputZipf*