test:
	go test -count 20 -failfast -run TestPatate/run . || true
	go test -count 20 -failfast -run TestPatate/runctx . || true 
	go test -count 100 -failfast -run TestPatate/runctxselect .
