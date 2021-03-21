# https://www.client9.com/self-documenting-makefiles/
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help

run: ## Run the script
	go run .
test: ## Test all files
	ginkgo -r
lint: ## Lint all the files
	golint ./...
delete-merged-branches: ## Delete all local branches merged to main, unless they start with dev
	git branch --merged | grep -i -v -E "main|master|dev"| xargs git branch -d
