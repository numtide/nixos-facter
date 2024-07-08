{
  lib,
  config,
  facterLib,
  ...
}: let
  inherit (config.facter) report;
  cfg = config.facter.virtualization;
in {
  options.facter.virtualization.enable =
    lib.mkEnableOption "Enable the Facter Virtualization module"
    // {
      default = builtins.any (facterLib.devicesFilter facterLib.pci.devices.virtio_scsi) report.hardware;
      defaultText = "hardware dependent";
    };

  config = lib.mkIf cfg.enable {
    boot.initrd.availableKernelModules = ["virtio_scsi"];
  };
}
