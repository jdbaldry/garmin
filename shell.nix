{ pkgs ? import <nixpkgs> }:

with pkgs;
let
  gci = pkgs.callPackage ./gci.nix { inherit pkgs; };
  sqlc = pkgs.callPackage ./sqlc.nix { inherit pkgs; };
in
mkShell {
  buildInputs = [
    gci
    go_1_19
    gofumpt
    golangci-lint
    gopls
    gotools

    jsonnet

    podman-compose
    postgresql

    sqlc
    jmtpfs

    unzip
  ];
  shellHook = ''
    GOROOT=${go_1_19}/share/go
  '';
}
