{
  lib,
  config,
  facterLib,
  ...
}: let
  inherit (config.facter) report;
  cfg = config.facter.networking.intel;
in {
  options.facter.networking.intel = with lib; {
    _2200BG.enable =
      mkEnableOption "Enable the Facter Intel 2200BG module"
      // {
        default =
          builtins.any
          (facterLib.devicesFilter facterLib.pci.devices.intel_2200BG)
          report.hardware;
        defaultText = "hardware dependent";
      };
    _3945ABG.enable =
      mkEnableOption "Enable the Facter Intel 3945ABG module"
      // {
        default =
          builtins.any
          (facterLib.devicesFilter facterLib.pci.devices.intel_3945ABG)
          report.hardware;
        defaultText = "hardware dependent";
      };
  };

  config.networking = lib.mkIf cfg._2200BG.enable {
    enableIntel2200BGFirmware = true;
  };

  config.hardware = lib.mkIf cfg._3945ABG.enable {
    enableRedistributableFirmware = true;
  };
}
