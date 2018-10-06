PROJECT := testK8s

BUILD_DIR := out
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

$(BUILD_DIR)/$(PROJECT): $(BUILD_DIR) $(GO_FILES)
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(PROJECT) --ldflags '-extldflags "-static"' .

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)