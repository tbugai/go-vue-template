NODE_BIN = $(shell npm bin)

IMPORT_PATH = $(shell echo `pwd` | sed "s|^$(GOPATH)/src/||g")
TARGET = $(shell echo $(IMPORT_PATH) | sed 's:.*/::')
PID = $(TARGET).pid

GO_FILES = $(shell find . -type f -name "*.go")
ASSETS = $(shell find assets -type f)

FRONTEND_PORT ?= 8000
BACKEND_PORT ?= 9000

build: clean $(BUNDLE) $(TARGET)

clean:
	@rm -rf public/bundles
	@rm -rf $(TARGET)

$(TARGET): $(GO_FILES)
	@printf "Buiding '$(TARGET)' binary ......"
	@go build -race -o $@

kill:
	@kill `cat $(PID)` || true

dev: clean restart
	FRONTEND_PORT=$(FRONTEND_PORT) BACKEND_PORT=$(BACKEND_PORT) $(NODE_BIN)/webpack-dev-server --config webpack.config.js &
	@printf "\n\nWaiting for the file change\n\n"
	@fswatch --event=Updated $(GO_FILES) | xargs -n1 -I{} make restart || make kill

restart: kill $(TARGET)
	@printf "\n\nrestart the app .........\n\n"
	@./$(TARGET) -port $(BACKEND_PORT) -debug & echo $$! > $(PID)

