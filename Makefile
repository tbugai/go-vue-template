TARGET = "server"
GO_FILES = $(shell find . -type f -name "*.go")
ASSETS = $(shell find assets -type f)
PID = .pid

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
	@printf "\n\nWaiting for the file change\n\n"
	@fswatch --event=Updated $(GO_FILES) | xargs -n1 -I{} make restart || make kill

restart: kill $(TARGET)
	@printf "\n\nrestart the app .........\n\n"
	@./$(TARGET) -debug & echo $$! > $(PID)

