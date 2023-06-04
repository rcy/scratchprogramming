with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    nodejs-18_x

    entr

    go
    golint
    gopls
    sqlite
    flyctl
    golangci-lint
  ];
}
