{
  description = "tag generates artwork from your favourite terminal themes";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = { self, nixpkgs }:
    let
      version = "0.2.0";
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        with pkgs;
        {
          devShell = mkShell {
            buildInputs = [
              go
              gopls
              git-chglog
            ];
          };
          tag = buildGoModule {
            pname = "tag";
            inherit version;
            src = ./.;
            # vendorSha256 = lib.fakeSha256;
            vendorSha256 = "sha256-WyJkorlxaj6HC+RO8rFCbVpPvTr6la52vMqvG2EIKk8=";
          };
        });
      defaultPackage = forAllSystems (system: self.packages.${system}.tag);
    };
}
