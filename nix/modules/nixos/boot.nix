{
  lib,
  config,
  facterLib,
  ...
}: let
  cfg = config.facter.boot;
  inherit (config.facter) report;
in {
  options.facter.boot.enable =
    lib.mkEnableOption "Enable the Facter Boot module"
    // {
      default = true;
    };

  config = lib.mkIf cfg.enable (facterLib.canonicalize {
    boot.initrd.availableKernelModules =
      builtins.filter
      (dm: dm != null)
      (
        builtins.map
        ({driver_module ? null, ...}: driver_module)
        (
          builtins.filter (with facterLib;
            isOneOf [
              # Needed if we want to use the keyboard when things go wrong in the initrd.
              isUsbController
              # A disk might be attached.
              isFirewireController
              # definitely important
              isMassStorageController
            ])
          report.hardware
        )
      );
  });
}
