package protocol

// Download metaModel.json

//go:generate wget -O metaModel.json https://raw.githubusercontent.com/microsoft/vscode-languageserver-node/205a43a011713e8d5d668f76d1b666d40dfc1ede/protocol/metaModel.json

// Generate lsp protocol types and interface

//go:generate go run github.com/black-desk/notels/tools/lspgen
//go:generate go run github.com/segmentio/golines@latest -m 80 -t 8 --no-ignore-generated --shorten-comments -w .
