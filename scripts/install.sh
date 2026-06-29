#!/usr/bin/env bash

set -e

APP_NAME="taskpilot"
INSTALL_DIR="$HOME/.local/bin"

echo "==> Installing ${APP_NAME}"

detect_package_manager() {
	if command -v pacman >/dev/null 2>&1; then
		echo "pacman"
	elif command -v apt >/dev/null 2>&1; then
		echo "apt"
	elif command -v dnf >/dev/null 2>&1; then
		echo "dnf"
	elif command -v brew >/dev/null 2>&1; then
		echo "brew"
	else
		echo ""
	fi
}

PKG_MANAGER=$(detect_package_manager)

if [ -z "$PKG_MANAGER" ]; then
	echo "Unsupported package manager."
	exit 1
fi

echo "Detected package manager: $PKG_MANAGER"

install_go() {
	case "$PKG_MANAGER" in
	pacman)
		sudo pacman -Sy --needed go
		;;
	apt)
		sudo apt update
		sudo apt install -y golang
		;;
	dnf)
		sudo dnf install -y golang
		;;
	brew)
		brew install go
		;;
	esac
}

install_taskwarrior() {
	case "$PKG_MANAGER" in
	pacman)
		sudo pacman -Sy --needed task
		;;
	apt)
		sudo apt update
		sudo apt install -y taskwarrior
		;;
	dnf)
		sudo dnf install -y task
		;;
	brew)
		brew install task
		;;
	esac
}

echo
echo "==> Checking Go"

if ! command -v go >/dev/null 2>&1; then
	echo "Go not found. Installing..."
	install_go
else
	echo "Go found: $(go version)"
fi

echo
echo "==> Checking Taskwarrior"

if ! command -v task >/dev/null 2>&1; then
	echo "Taskwarrior not found. Installing..."
	install_taskwarrior
else
	echo "Taskwarrior found"
fi

echo
echo "==> Building ${APP_NAME}"

go build -o "${APP_NAME}" ./cmd/task

echo
echo "==> Installing binary"

mkdir -p "$INSTALL_DIR"

cp "${APP_NAME}" "${INSTALL_DIR}/${APP_NAME}"

chmod +x "${INSTALL_DIR}/${APP_NAME}"

echo
echo "==> Checking optional assistants"

if command -v claude >/dev/null 2>&1; then
	echo "✓ Claude CLI detected"
else
	echo "⚠ Claude CLI not found"
fi

if command -v ollama >/dev/null 2>&1; then
	echo "✓ Ollama detected"
else
	echo "⚠ Ollama not found"
fi

echo
echo "==> Installation complete"

echo
echo "Binary installed at:"
echo "  ${INSTALL_DIR}/${APP_NAME}"

echo
echo "Ensure this directory is in your PATH:"
echo "  ${INSTALL_DIR}"

echo
echo "Try:"
echo "  taskpilot sync"
echo "  taskpilot work claude JAT-1"
