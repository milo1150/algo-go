run go test -count=1 ./...

NOTE:
-count=1 : Ensures that the tests are not cached and are executed fresh every time.
./... : Recursively tests all packages and subdirectories.