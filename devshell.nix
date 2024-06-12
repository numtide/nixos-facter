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
      name = "GOROOT";
      value = pkgs.go + "/share/go";
    }
  ];

  packages = lib.mkMerge [
    (with pkgs; [
      go
      delve
      pprof
      graphviz
    ])
  ];

  commands = [
    {
      category = "development";
      package = perSystem.gomod2nix.default;
    }
  ];

}
