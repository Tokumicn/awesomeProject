# 设置变量
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GODEPS := $(GOCMD) mod download
GOGENERATE := $(GOCMD) generate
GOLINTER := golangci-lint run
BINARY_NAME := yourprojectname
MAIN_FILE := main.go

# 设置要排除的目录列表（根据实际情况更改）
EXCLUDE_DIRS := ./vendor ./.git ./.idea ./examples ./test

# 查找所有非排除目录的目录
SRC_DIRS := $(shell find . -type d $(foreach dir,$(EXCLUDE_DIRS),-not -path "$(dir)*"))

# 构建目标：生成代码
generate:
    @for dir in $(SRC_DIRS); do \
        echo "Generating code in directory: $$dir"; \
        (cd $$dir && $(GOGENERATE) -v); \
    done

# 构建目标：代码格式检测
lint:
    $(GOLINTER) ./...

# 构建目标：运行测试
test:
    $(GOTEST) ./...

# 构建目标：编译代码
build:
    $(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# 构建目标：清理项目
clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)

# 构建目标：安装依赖
deps:
    $(GODEPS)

# 构建目标：执行所有构建步骤
all: generate lint test build

# 声明所有目标，确保它们被视为伪目标而不是实际的文件名
.PHONY: generate lint test build clean deps all