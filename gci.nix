{ pkgs ? import <nixpkgs> }:

with pkgs;
buildGoModule rec {
  pname = "gci";
  version = "0.9.0";

  src = fetchFromGitHub {
    owner = "daixiang0";
    repo = pname;
    rev = "v${version}";
    sha256 = "sha256-qWEEcIbTgYmGVnnTW+hM8e8nw5VLWN1TwzdUIZrxF3s=";
  };

  vendorSha256 = "sha256-dlt+i/pEP3RzW4JwndKTU7my2Nn7/2rLFlk8n1sFR60=";

  meta = with lib; {
    description = "GCI, a tool that control golang package import order and make it always deterministic.";
    homepage = "https://github.com/daixiang0/gci";
    license = with licenses; bsd3;
    maintainers = with maintainers; [ jdbaldry ];
  };
}
