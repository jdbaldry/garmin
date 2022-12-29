{ pkgs ? import <nixpkgs> }:

with pkgs;
buildGoModule rec {
  pname = "fit";
  version = "0.14.0";

  src = fetchFromGitHub {
    owner = "tormoder";
    repo = pname;
    rev = "v${version}";
    sha256 = "sha256-uR70Q+fgSZGjy24sy6BGiHX7rgiGHqwIxkhZ3QX76/4=";
  };
  vendorSha256 = "sha256-M7af046D2pb2I0THCaOhO9e3MyI/ssJQ9/sO2wh8HT4=";

  meta = with lib; {
    description = "A Go package for decoding and encoding Garmin FIT files.";
    homepage = "https://github.com/tormoder/fit";
    license = with licenses; mit;
    maintainers = with maintainers; [ jdbaldry ];
  };
}
