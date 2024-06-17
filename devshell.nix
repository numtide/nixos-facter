{
  inputs,
  pkgs,
  system,
  perSystem,
  ...
}:
let
  inherit (pkgs) lib;
  # we can't use perSystem.devshell.mkShell because it prefers ".packages" over ".legacyPackages" and the devshell
  # flake doesn't expose these utils, only the legacy compat stuff does
  inherit (inputs.devshell.legacyPackages.${system}) mkShell;
in
mkShell {

  env = [
    {
      name = "DEVSHELL_NO_MOTD";
      value = 1;
    }
    {
      name = "GOROOT";
      value = pkgs.go + "/share/go";
    }
    {
      name = "LD_LIBRARY_PATH";
      prefix = "$DEVSHELL_DIR/lib";
    }
    {
      name = "C_INCLUDE_PATH";
      prefix = "$DEVSHELL_DIR/include";
    }
    {
      name = "PKG_CONFIG_PATH";
      prefix = "$DEVSHELL_DIR/lib/pkgconfig";
    }
  ];

  packages = lib.mkMerge [
    (with pkgs; [
      go
      delve
      pprof
      graphviz
      libusb1.dev
      gcc
      pkg-config
    ])
  ];

  commands = [
    { package = perSystem.gomod2nix.default; }
    { package = pkgs.golangci-lint; }
    { package = pkgs.cobra-cli; }
  ];

}
