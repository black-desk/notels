# LSP package design note

This file is a note to explain why the pkg/lsp package design like this.

## Why we need a new lsp package?

Currently, we have three choice:

- [glsp][==link1==]
- [go-lsp][==link2==]
- [gopls internal][==link3==]

These packages are all hard to maintain when the protocol itself upgrade, it
result that they are all stucking at lsp 3.15, which is quite old.

But the LSP maintainer now provide [a json file][==link4==] generating from the
reference implementation of LSP protocol to describe the protocol itself,
including the message and types protocol using.

So if we can generate go code from that json, we will have a easy-to-update LSP
SDK.

Peter Weinberger is working on this for gopls too.[^fn1]

## Why place this SDK in this project?

When it ready for using in other project, I will create a new git repository for
this package. But now, as I working on my language server at the same time, it
is more convenient to fix bugs in the SDK itself.

## Why everything is in a structure?

This idea come from the in-progress work of Peter Weinberger. I found it will
simplify the code generator.

[^fn1]:
    - <https://go-review.googlesource.com/c/tools/+/430915>
    - <https://go-review.googlesource.com/c/tools/+/443055>
    - <https://go-review.googlesource.com/c/tools/+/446595>

[==link1==]: https://github.com/tliron/glsp
[==link2==]: https://github.com/sourcegraph/go-lsp
[==link3==]: https://cs.opensource.google/go/x/tools/+/master:gopls/internal/lsp/
[==link4==]: https://github.com/microsoft/vscode-languageserver-node/blob/main/protocol/metaModel.json
