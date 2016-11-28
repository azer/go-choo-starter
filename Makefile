include .env

PROJECTNAME=$(shell basename "$(PWD)")
GOBASE=$(shell pwd)/server
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard server/*.go)
PID=/tmp/go-$(PROJECTNAME).pid
UIBASE=$(shell pwd)/ui
UIBIN=$(UIBASE)/node_modules/.bin

all: develop

start:
	@echo "  ▶️  Starting $(PROJECTNAME) at $(ADDR)"
	@-$(GOBIN)/$(PROJECTNAME) & echo $$! > $(PID)
	@echo "  🆗  Process ID: "$(shell cat $(PID))

stop:
	@echo "  ⏹  Stopping $(PROJECTNAME)"
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true

restart: stop start
build: go-build ui-build
clean: go-clean ui-clean

develop: setup build
	@DEVELOP=1 LOG=* $(MAKE) restart
	@echo "  👓  Watching for changes..."
	@fswatch server/. -e "server/bin" -e "server/pkg" | (while read; do DEVELOP=1 LOG=* make setup clean build restart; done)

setup:
	@echo "  🔄  Please wait while I'm getting the dependencies of $(PROJECTNAME) from internet."
	@$(MAKE) go-get
	@$(MAKE) ui-install

go-build:
	@echo "  🛠  Building server into $(GOBIN)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-get:
	@cd server && GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

go-install:
	@cd server && GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

go-clean:
	@echo "  🚿  Cleaning Go build cache"
	@cd server && GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

go-rebuild: go-clean go-build

ui-install:
	@cd ui && npm install --silent

ui-build:
	@echo "  🐣  Building UI into ./public"
	@mkdir -p ./public
	@cd ui && ./node_modules/.bin/browserify -r min-document --igv __filename,__dirname,_process -t [ babelify --presets [ es2015 ] ] client-side.js > ../public/dist.js
	@cd ui && cat *.css components/**/*.css > ../public/dist.css

ui-clean:
	@echo "  🚿  Cleaning UI builds"
	@-rm public/{dist.js,dist.css} 2> /dev/null || true

ui-build-serverside:
	@cd ui && ./node_modules/.bin/browserify -r min-document --igv __filename,__dirname,_process -t [ babelify --presets [ es2015 ] ] --debug server-side.js

export COMPONENT_INDEX
export COMPONENT_VIEW
create-component:
	define COMPONENT_INDEX
		import view from './view'
		import state from './state'
		import * as reducers from './reducers'
		import * as effects from './effects'

	export default {
		namespace: '$(name)',
		view,
		state,
		reducers,
		effects
	}
	endef

	define COMPONENT_VIEW
		import html from "choo/html"

		const view = (state, prev, send) => html`
			$(name)
		`

		export default view
	endef

	@mkdir ui/components/${name}
	@echo "$$COMPONENT_INDEX" > ui/components/${name}/index.js
	@echo "$$COMPONENT_VIEW" > ui/components/${name}/view.js
	@echo "export default {}" > ui/components/${name}/reducers.js
	@echo "export default {}" > ui/components/${name}/effects.js
	@echo "export default {}" > ui/components/${name}/state.js

commands:
	@cat docs/man

usage: commands
help: commands

.PHONY: default go-build go-get go-install go-run go-rebuild go-clean ui-build ui-clean ui-build-serverside start stop restart clean commands help usage
