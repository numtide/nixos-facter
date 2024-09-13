{
  pkgs,
  inputs,
  system,
  perSystem,
  ...
}: let
  # we have to import diskoLib like this because there are some impure default imports e.g. <nixpkgs>
  diskoLib = import "${inputs.disko}/lib" {
    inherit (pkgs) lib;
    makeTest = import "${inputs.nixpkgs}/nixos/tests/make-test-python.nix";
    eval-config = import "${inputs.nixpkgs}/nixos/lib/eval-config.nix";
  };
in
  # for now we only run the tests in x86_64-linux since we don't have access to a bare-metal ARM box or a VM that supports nested
  # virtualisation which makes the test take forever and ultimately fail
  pkgs.lib.optionalAttrs pkgs.stdenv.isx86_64 {
    basic = diskoLib.testLib.makeDiskoTest {
      inherit pkgs;
      name = "basic";
      disko-config = ./disko.nix;
      extraSystemConfig = {
        environment.systemPackages = [
          perSystem.self.nixos-facter
        ];

        systemd.services = {
          create-swap-files = {
            serviceConfig.Type = "oneshot";
            wantedBy = ["multi-user.target"];
            path = with pkgs; [
              coreutils
              util-linux
            ];
            script = ''
              # create some swap files
              mkdir -p /swap
              for (( i=1; i<=3; i++ )); do
                out="/swap/swapfile-$i"
                dd if=/dev/zero of="$out" bs=1MB count=10
                chmod 600 "$out"
                mkswap "$out"
                swapon "$out"
              done
            '';
          };
        };
      };

      extraTestScript = ''
        import json

        report = json.loads(machine.succeed("nixos-facter -e"))

        with subtest("Capture system"):
            assert report['system'] == '${system}'

        with subtest("Capture virtualisation"):
            virt = report['virtualisation']
            # kvm for systems that support it, otherwise the vm test should present itself as qemu
            # todo double-check this is the same for intel
            assert virt in ("kvm", "qemu"), f"expected virtualisation to be either kvm or qemu, got {virt}"

        with subtest("Capture swap entries"):
            assert 'swap' in report, "'swap' not found in the report"
            assert report['swap'] == [
                { 'path': '/dev/vda4', 'type': 'partition', 'size': 1048572, 'used': 0, 'priority': -2 },
                { 'path': '/dev/dm-0', 'type': 'partition', 'size': 10236, 'used': 0, 'priority': 100 },
                { 'path': '/swap/swapfile-1', 'type': 'file', 'size': 9760, 'used': 0, 'priority': -3 },
                { 'path': '/swap/swapfile-2', 'type': 'file', 'size': 9760, 'used': 0, 'priority': -4 },
                { 'path': '/swap/swapfile-3', 'type': 'file', 'size': 9760, 'used': 0, 'priority': -5 }
            ], "swap entries did not match what we expected"
            # todo is there a nice way of showing a diff?
      '';
    };
  }
