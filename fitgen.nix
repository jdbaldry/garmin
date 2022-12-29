{ pkgs ? import <nixpkgs> }:

with pkgs;
buildGoModule rec {
  pname = "fit";
  version = "0.14.0";

  src = fetchFromGitHub {
    owner = "tormoder";
    repo = pname;
    rev = "v${version}";
    sha256 = lib.fakeSha256;
  };
  vendorSha256 = lib.fakeSha256;

  meta = with lib; {
    description = "A Go package for decoding and encoding Garmin FIT files.";
    homepage = "https://github.com/tormoder/fit";
    license = with licenses; mit;
    maintainers = with maintainers; [ jdbaldry ];
  };
}
