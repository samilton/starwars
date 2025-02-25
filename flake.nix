{
  description = "A Go project";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            golangci-lint
            delve
            go-tools
          ];

          # Set environment variables if needed
          shellHook = ''
            export GOPATH="$PWD/.go"
            export PATH="$GOPATH/bin:$PATH"
          '';
        };

        # Optional: Add a package definition if you want to build your Go project
        packages.default = pkgs.buildGoModule {
          pname = "starwars";
          version = "0.1.0";
          src = ./.;
          
          # Update this with the actual hash of your Go dependencies
          # Use 'vendorSha256 = null;' first to get the correct hash
          vendorSha256 = null;
        };
      }
    );
}