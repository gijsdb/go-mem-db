plantuml: 
	goplantuml -aggregate-private-members -show-aggregations -show-connection-labels -recursive ./ > serverClassDiagram.puml
	plantuml serverClassDiagram.puml 
