.PHONY: clean
clean: ## remove build artifacts
	rm -rf ./build/*

.PHONY: binary-osx
binary-osx: ## build executable for macOS
	./scripts/build/osx.sh

.PHONY: binary-windows
binary-windows: ## build executable for Window 
	./scripts/build/windows.sh

.PHONY: binary-linux
binary-linux: ## build executable for linuxs
	./scripts/build/linux.sh
