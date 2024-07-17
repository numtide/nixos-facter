lib: let
  inherit (lib) fold mapAttrsRecursive sort unique;
in rec {
  isAllOf = filters: device: fold (next: memo: memo && (next device)) true filters;
  isOneOf = filters: device: fold (next: memo: memo || (next device)) false filters;

  isMassStorageController = {base_class ? {}, ...}:
    (base_class.value or null) == 1;
  isNetworkController = {base_class ? {}, ...}:
    (base_class.value or null) == 2;
  isDisplayController = {base_class ? {}, ...}:
    (base_class.value or null) == 3;
  isMultimediaController = {base_class ? {}, ...}:
    (base_class.value or null) == 4;
  isMemoryController = {base_class ? {}, ...}:
    (base_class.value or null) == 5;
  isBridge = {base_class ? {}, ...}:
    (base_class.value or null) == 6;
  isCommunicationController = {base_class ? {}, ...}:
    (base_class.value or null) == 7;
  isGenericSystemPeripheral = {base_class ? {}, ...}:
    (base_class.value or null) == 8;
  isInputDeviceController = {base_class ? {}, ...}:
    (base_class.value or null) == 9;
  isDockingStation = {base_class ? {}, ...}:
    (base_class.value or null) == 10;
  isProcessor = {base_class ? {}, ...}:
    (base_class.value or null) == 11;
  isSerialBusController = {base_class ? {}, ...}:
    (base_class.value or null) == 12;
  isWirelessController = {base_class ? {}, ...}:
    (base_class.value or null) == 13;
  isIntelligentController = {base_class ? {}, ...}:
    (base_class.value or null) == 14;
  isSatelliteCommunicationsController = {base_class ? {}, ...}:
    (base_class.value or null) == 15;
  isEncryptionController = {base_class ? {}, ...}:
    (base_class.value or null) == 16;
  isSignalProcessingController = {base_class ? {}, ...}:
    (base_class.value or null) == 17;
  isProcessingAccelerator = {base_class ? {}, ...}:
    (base_class.value or null) == 18;
  isNonEssentialInstrumentation = {base_class ? {}, ...}:
    (base_class.value or null) == 19;
  isCoprocessor = {base_class ? {}, ...}:
    (base_class.value or null) == 64;

  isFirewireController = isAllOf [
    isSerialBusController
    (
      {sub_class ? {}, ...}:
        (sub_class.value or 9999) == 0
    )
  ];

  isUsbController = isAllOf [
    isSerialBusController
    (
      {sub_class ? {}, ...}:
        (sub_class.value or 9999) == 3
    )
  ];

  isCpuAmd = {hardware, ...}:
    builtins.filter
    (device: device.hardware_class == "cpu" && device.detail.vendor_name == "AuthenticAMD")
    hardware;

  isCpuIntel = {hardware, ...}:
    builtins.filter
    (device: device.hardware_class == "cpu" && device.detail.vendor_name == "GenuineIntel")
    hardware;

  canonicalize = attrs:
    mapAttrsRecursive (_: value:
      if builtins.isList value
      then canonicalSort value
      else value)
    attrs;

  canonicalSort = list: sort (a: b: a < b) (unique list);

  devicesFilter = {
    vendorId,
    deviceIds,
  }:
    isAllOf [
      (item: (item.vendor.value or null) == vendorId)
      (item: builtins.elem (item.device.value or null) deviceIds)
    ];

  pci.devices = import ./pci/devices.nix;

  nixosModules = reportPath: [
    (import ../modules/nixos/facter.nix)
    {config.facter.report = builtins.fromJSON (builtins.readFile reportPath);}
  ];
}
