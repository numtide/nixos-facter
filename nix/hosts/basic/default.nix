{
  flake,
  inputs,
  ...
}: let
  facterLib = flake.lib;
in
  inputs.nixpkgs.lib.nixosSystem {
    system = "x86_64-linux";
    specialArgs = {
      inherit flake inputs facterLib;
    };
    modules =
      [./config.nix]
      ++ (facterLib.nixosModules ./report.json);
  }
