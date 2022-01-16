# https://www.client9.com/self-documenting-makefiles/
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help

test: ## Test all files
	go test -v ./...
lint: ## Lint and format all the files
	for d in $$(go list -f {{.Dir}} ./...); do \
		gofmt -w $${d}/*.go; \
	done
delete-merged-branches: ## Delete all local branches merged to main, unless they start with dev
	git branch --merged | grep -i -v -E "main|master|dev"| xargs git branch -d
