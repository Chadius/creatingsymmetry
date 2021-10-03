# https://www.client9.com/self-documenting-makefiles/
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help

FORMULA_FILENAME="data/formula.yml"

run: ## Run the script with default arguments
	go run . -f $(FORMULA_FILENAME)
test: ## Test all files
	go test -v ./...
lint: ## Lint and format all the files
	golint ./...
	gofmt -w ./..
delete-merged-branches: ## Delete all local branches merged to main, unless they start with dev
	git branch --merged | grep -i -v -E "main|master|dev"| xargs git branch -d
