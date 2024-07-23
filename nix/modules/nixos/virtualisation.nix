{
  lib,
  config,
  facterLib,
  ...
}: let
  inherit (config.facter) report;
  cfg = config.facter.virtualisation;
in {
  options.facter.virtualisation = {
    virtio.enable =
      lib.mkEnableOption "Enable the Facter Virtualisation Virtio module"
      // {
        default = builtins.any (facterLib.devicesFilter facterLib.pci.devices.virtio_scsi) report.hardware;
        defaultText = "hardware dependent";
      };
    oracle.enable =
      lib.mkEnableOption "Enable the Facter Virtualisation Oracle module"
      // {
        default = report.virtualisation == "oracle";
        defaultText = "environment dependent";
      };
    parallels.enable =
      lib.mkEnableOption "Enable the Facter Virtualisation Parallels module"
      // {
        default = report.virtualisation == "parallels";
        defaultText = "environment dependent";
      };
    qemu.enable =
      lib.mkEnableOption "Enable the Facter Virtualisation Qemu module"
      // {
        default = builtins.elem report.virtualisation ["qemu" "kvm" "bochs"];
        defaultText = "environment dependent";
      };
    hyperv.enable =
      lib.mkEnableOption "Enable the Facter Virtualisation Hyper-V module"
      // {
        default = report.virtualisation == "microsoft";
        defaultText = "environment dependent";
      };
    systemd_nspawn.enable =
      lib.mkEnableOption "Enable the Facter Virtualisation Systemd NSpawn module"
      // {
        default = report.virtualisation == "systemd-nspawn";
        defaultText = "environment dependent";
      };
    # no virtualisation detected
    none.enable =
      lib.mkEnableOption "Enable the Facter Virtualisation None module"
      // {
        default = report.virtualisation == "none";
        defaultText = "environment dependent";
      };
  };

  config = {
    # virtio & qemu
    boot.initrd = {
      kernelModules = lib.optionals cfg.qemu.enable [
        "virtio_balloon"
        "virtio_console"
        "virtio_rng"
        "virtio_gpu"
      ];
      availableKernelModules =
        (
          lib.optionals cfg.qemu.enable
          [
            "virtio_net"
            "virtio_pci"
            "virtio_mmio"
            "virtio_blk"
            "virtio_scsi"
            "9p"
            "9pnet_virtio"
          ]
        )
        ++ (lib.optionals cfg.virtio.enable ["virtio_scsi"]);
    };

    virtualisation = {
      # oracle
      virtualbox.guest.enable = cfg.oracle.enable;
      # hyper-v
      hypervGuest.enable = cfg.hyperv.enable;
    };

    # parallels
    hardware.parallels.enable = cfg.parallels.enable;
    nixpkgs.config.allowUnfreePredicate = lib.mkIf cfg.parallels.enable (pkg: builtins.elem (lib.getName pkg) ["prl-tools"]);

    # systemd-nspawn
    boot.isContainer = cfg.systemd_nspawn.enable;

    # none (e.g. bare-metal)
    # provide firmware for devices that might not have been detected by nixos-facter
    hardware.enableRedistributableFirmware = lib.mkDefault cfg.none.enable;

    # update microcode
    hardware.cpu.amd.updateMicrocode = lib.mkIf (cfg.none.enable && facterLib.isCpuAmd report) (lib.mkDefault config.hardware.enableRedistributableFirmware);
    hardware.cpu.intel.updateMicrocode = lib.mkIf (cfg.none.enable && facterLib.isCpuIntel report) (lib.mkDefault config.hardware.enableRedistributableFirmware);
  };
}
