{lib, ...}: {
  imports = [
    ./boot.nix
    ./networking
    ./virtualisation.nix
  ];

  options.facter = with lib; {
    report = mkOption {
      type = types.raw;
    };
  };
}
