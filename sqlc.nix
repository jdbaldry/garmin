{ pkgs ? import <nixpkgs> }:

with pkgs;
buildGoModule rec {
  pname = "sqlc";
  version = "1.16.0";

  buildInputs = [ xxHash ];
  src = fetchFromGitHub {
    owner = "kyleconroy";
    repo = pname;
    rev = "v${version}";
    sha256 = "sha256-YxGMfGhcPT3Pcyxu1hAkadkJnEBMX26fE/rGfGSTsyc=";
  };

  proxyVendor = true;
  vendorSha256 = "sha256-cMYTQ8rATCXOquSxc4iZ2MvxIaMO3RG8PZkpOwwntyc=";

  doCheck = false;
  preCheck = ''
    export XDG_CACHE_HOME=$(mktemp -d)
  '';

  meta = with lib; {
    description = "Compile SQL to type-safe code.";
    homepage = "https://sqlc.dev/";
    license = with licenses; mit;
    maintainers = with maintainers; [ jdbaldry ];
  };
}
