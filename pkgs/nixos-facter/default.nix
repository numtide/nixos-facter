{
  flake,
  inputs,
  system,
  pkgs,
  ...
}:
let
  inherit (pkgs) go lib;
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
        "cmd"
        "pkg"
        "go.mod"
        "go.sum"
        "main.go"
      ];
    };

  modules = ./gomod2nix.toml;

  buildInputs = [
    pkgs.libusb1
  ];

  nativeBuildInputs = with pkgs; [
    gcc
    pkg-config
  ];

  runtimeInputs = [
    pkgs.libusb1
  ];

  ldflags = [
    "-s"
    "-w"
    "-X git.numtide.com/numtide/nixos-facter/build.Name=${pname}"
    "-X git.numtide.com/numtide/nixos-facter/build.Version=v${version}"
  ];

  meta = with lib; {
    description = "nixos-facter: declarative nixos-generate-config";
    homepage = "https://github.com/numtide/nixos-facter";
    license = licenses.mit;
    mainProgram = "nixos-facter";
  };
}
