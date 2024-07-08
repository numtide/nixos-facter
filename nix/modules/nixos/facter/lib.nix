final: prev: {

  pci = rec {
    isMassStorageController = { base_class, ... }: base_class.value == 1;
    isNetworkController = { base_class, ... }: base_class.value == 2;
    isDisplayController = { base_class, ... }: base_class.value == 3;
    isMultimediaController = { base_class, ... }: base_class.value == 4;
    isMemoryController = { base_class, ... }: base_class.value == 5;
    isBridge = { base_class, ... }: base_class.value == 6;
    isCommunicationController = { base_class, ... }: base_class.value == 7;
    isGenericSystemPeripheral = { base_class, ... }: base_class.value == 8;
    isInputDeviceController = { base_class, ... }: base_class.value == 9;
    isDockingStation = { base_class, ... }: base_class.value == 10;
    isProcessor = { base_class, ... }: base_class.value == 11;
    isSerialBusController = { base_class, ... }: base_class.value == 12;
    isWirelessController = { base_class, ... }: base_class.value == 13;
    isIntelligentController = { base_class, ... }: base_class.value == 14;
    isSatelliteCommunicationsController = { base_class, ... }: base_class.value == 15;
    isEncryptionController = { base_class, ... }: base_class.value == 16;
    isSignalProcessingController = { base_class, ... }: base_class.value == 17;
    isProcessingAccelerator = { base_class, ... }: base_class.value == 18;
    isNonEssentialInstrumentation = { base_class, ... }: base_class.value == 19;
    isCoprocessor = { base_class, ... }: base_class.value == 64;

    isFirewireController =
      item@{ base_class, sub_class, ... }: (isSerialBusController item) && sub_class.value == 0;
    isUsbController =
      item@{ base_class, sub_class, ... }: (isSerialBusController item) && sub_class.value == 3;
  };

  initrd = {
    availableKernelModules =
      { hardware }:
      let
        filter =
          item:
          # Has a driver module
          (builtins.hasAttr "driver_module" item)
          &&
            # Definitely important
            (
              (final.pci.isMassStorageController item)
              ||
                # A disk might be attached.
                (final.pci.isFirewireController item)
              ||
                # Needed if we want to use the keyboard when things go wrong in the initrd.
                (final.pci.isUsbController item)
            );
      in
      # sort the final list to ensure a stable output
      final.sort (a: b: a < b) (
        # de-duplicate
        final.unique (
          final.flatten (
            final.mapAttrsToList (
              class: items: builtins.map (item: item.driver_module) (builtins.filter filter items)
            ) hardware
          )
        )
      );
  };
}
