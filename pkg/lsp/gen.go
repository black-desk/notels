package protocol

// Download metaModel.json

//go:generate wget -O metaModel.json https://raw.githubusercontent.com/microsoft/vscode-languageserver-node/f2864f8fb946ad8ff82928d5087d84210f77aa2d/protocol/metaModel.json
//go:generate wget -O metaModel.schema.json https://raw.githubusercontent.com/microsoft/vscode-languageserver-node/f2864f8fb946ad8ff82928d5087d84210f77aa2d/protocol/metaModel.schema.json

// The original metaModel.json has some problems.
// - https://github.com/microsoft/vscode-languageserver-node/issues/1125
// - https://github.com/microsoft/vscode-languageserver-node/issues/1103
// fix-metaModel.patch will fix them

//go:generate patch -u metaModel.json -i fix-metaModel.patch

// Generate lsp protocol types and interface

//go:generate go run github.com/black-desk/notels/tools/lspgen
//go:generate go run github.com/segmentio/golines@latest -m 80 -t 8 --no-ignore-generated --shorten-comments -w .
