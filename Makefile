PLANTUML_JAR_URL = https://sourceforge.net/projects/plantuml/files/plantuml.jar/download
DIAGRAMS_SRC := $(wildcard docs/diagrams/*.plantuml)
DIAGRAMS_PNG := $(addsuffix .png, $(basename $(DIAGRAMS_SRC)))
DIAGRAMS_SVG := $(addsuffix .svg, $(basename $(DIAGRAMS_SRC)))

.PHONY: help clean diagrams png-diagrams svg-diagrams
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

clean: ## Cleans plantuml.jar and generated diagrams
	@rm -f plantuml.jar $(DIAGRAMS_PNG) $(DIAGRAMS_SVG)

diagrams: svg-diagrams png-diagrams ## Generate diagrams in SVG and PNG format
svg-diagrams: plantuml.jar $(DIAGRAMS_SVG) ## Generate diagrams in SVG format
png-diagrams: plantuml.jar $(DIAGRAMS_PNG) ## Generate diagrams in PNG format

plantuml.jar:
	@echo Downloading $@....
	@curl -sSfL $(PLANTUML_JAR_URL) -o $@

docs/diagrams/%.svg: docs/diagrams/%.plantuml
	@echo Generating $^ from plantuml....
	@java -jar plantuml.jar -tsvg $^

docs/diagrams/%.png: docs/diagrams/%.plantuml
	@echo Generating $^ from plantuml....
	@java -jar plantuml.jar -tpng $^
