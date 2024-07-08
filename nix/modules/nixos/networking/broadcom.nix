{
  lib,
  config,
  facterLib,
  ...
}: let
  inherit (config.facter) report;
  cfg = config.facter.networking.broadcom;
in {
  options.facter.networking.broadcom = with lib; {
    full_mac.enable =
      mkEnableOption "Enable the Facter Broadcom Full MAC module"
      // {
        default =
          builtins.any
          (facterLib.devicesFilter facterLib.pci.devices.broadcom_full_mac)
          report.hardware;
        defaultText = "hardware dependent";
      };
    sta.enable =
      mkEnableOption "Enable the Facter Broadcom STA module"
      // {
        default =
          builtins.any
          (facterLib.devicesFilter facterLib.pci.devices.broadcom_sta)
          report.hardware;
        defaultText = "hardware dependent";
      };
  };

  config.hardware = lib.mkIf cfg.full_mac.enable {
    enableRedistributableFirmware = true;
  };

  config.boot = lib.mkIf cfg.sta.enable {
    kernelModules = ["wl"];
    extraModulePackages = [config.boot.kernelPackages.broadcom_sta];
  };
}
