{
  pkgs,
  flake,
  system,
  ...
}: let
  inherit (flake.packages.${system}) nixos-facter;
in {
  basic = pkgs.nixosTest {
    name = "basic";
    nodes.machine = {
      environment.systemPackages = [nixos-facter];
    };
    testScript = ''
      machine.succeed("nixos-facter generate report -p -o /report.json")
    '';
  };
}
