{ pkgs ? import <nixpkgs> }:

with pkgs;
let
  fitgen = pkgs.callPackage ./fitgen.nix { inherit pkgs; };
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
    openjdk
    podman-compose
    postgresql

    fitgen
    gnumeric
    sqlc
    jmtpfs

    unzip
  ];
  shellHook = ''
    GOROOT=${go_1_19}/share/go
  '';
}
