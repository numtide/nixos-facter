{
  inputs,
  pkgs,
  system,
  perSystem,
  ...
}:
let
  # we can't use perSystem.devshell.mkShell because it prefers ".packages" over ".legacyPackages" and the devshell
  # flake doesn't expose these utils, only the legacy compat stuff does
  inherit (inputs.devshell.legacyPackages.${system}) mkShell;
in
mkShell { }
