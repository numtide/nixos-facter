{
  flake,
  inputs,
  system,
  pkgs,
  ...
}:
let
  inherit (pkgs) go lib musl;
in
inputs.gomod2nix.legacyPackages.${system}.buildGoApplication rec {
  pname = "nixos-facter";
  # there's no good way of tying in the version to a git tag or branch
  # so for simplicity's sake we set the version as the commit revision hash
  # we remove the `-dirty` suffix to avoid a lot of unnecessary rebuilds in local dev
  version = lib.removeSuffix "-dirty" (flake.shortRev or flake.dirtyShortRev);

  # ensure we are using the same version of go to build with
  inherit go;

  src =
    let
      filter = inputs.nix-filter.lib;
    in
    filter {
      root = ../../.;
      include = [
        "go.mod"
        "go.sum"
        "gomod2nix.toml"
        (filter.matchExt "go")
      ];
    };

  modules = ./gomod2nix.toml;

  CGO_ENABLED = 0;

  ldflags = [
    "-s"
    "-w"
    "-X git.numtide.com/numtide/treefmt/build.Name=${pname}"
    "-X git.numtide.com/numtide/treefmt/build.Version=v${version}"
  ];

  meta = with lib; {
    description = "nixos-facter: declarative nixos-generate-config";
    homepage = "https://github.com/numtide/nixos-facter";
    license = licenses.mit;
    mainProgram = "nixos-facter";
  };
}
