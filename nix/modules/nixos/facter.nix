{lib, ...}: {
  imports = [
    ./boot.nix
    ./networking
    ./virtualization.nix
  ];

  options.facter = with lib; {
    report = mkOption {
      type = types.raw;
    };
  };
}
