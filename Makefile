plantuml: 
	goplantuml -aggregate-private-members -show-aggregations -show-connection-labels -recursive ./ > serverClassDiagram.puml
	plantuml serverClassDiagram.puml 

runcli:
	go run cmd/cli/main.go localhost:4242

runserver:
	go run cmd/server/main.go

testcoverage:
	go test ./... -coverprofile testcoverage.out