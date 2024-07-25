{
  pkgs,
  flake,
  system,
  ...
}: let
  inherit (flake.packages.${system}) nixos-facter;
in
  # for now we only run the tests in x86_64-linux since we don't have access to a bare-metal ARM box or a VM that supports nested
  # virtualization which makes the test take forever and ultimately fail
  pkgs.lib.optionalAttrs pkgs.stdenv.isx86_64 {
    basic = pkgs.nixosTest {
      name = "basic";
      nodes.machine = {
        environment.systemPackages = [nixos-facter];
      };
      testScript = ''
        machine.succeed("nixos-facter generate report -o /report.json")
      '';
    };
  }
