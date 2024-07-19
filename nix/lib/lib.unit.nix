let
  # todo is there a better way of doing this?
  pkgs = import <nixpkgs> {};
  inherit (pkgs) lib;

  facterLib = import ./lib.nix lib;

  usbController = {
    base_class = {
      value = 12;
      name = "Serial bus controller";
    };
    sub_class = {
      value = 3;
      name = "USB Controller";
    };
    driver = "xhci_hcd";
    driver_module = "xhci_pci";
  };
in
  with facterLib; {
    testIsMassStorageController = {
      expr = builtins.map isMassStorageController [
        {}
        {base_class = {};}
        {base_class.value = 2;}
        {base_class.value = 1;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsNetworkController = {
      expr = builtins.map isNetworkController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 2;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsDisplayController = {
      expr = builtins.map isDisplayController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 3;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsMultimediaController = {
      expr = builtins.map isMultimediaController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 4;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsMemoryController = {
      expr = builtins.map isMemoryController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 5;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsBridge = {
      expr = builtins.map isBridge [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 6;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsCommunicationController = {
      expr = builtins.map isCommunicationController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 7;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsGenericSystemPeripheral = {
      expr = builtins.map isGenericSystemPeripheral [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 8;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsInputDeviceController = {
      expr = builtins.map isInputDeviceController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 9;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsDockingStation = {
      expr = builtins.map isDockingStation [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 10;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsProcessor = {
      expr = builtins.map isProcessor [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 11;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsSerialBusController = {
      expr = builtins.map isSerialBusController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 12;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsWirelessController = {
      expr = builtins.map isWirelessController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 13;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsIntelligentController = {
      expr = builtins.map isIntelligentController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 14;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsSatelliteCommunicationsController = {
      expr = builtins.map isSatelliteCommunicationsController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 15;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsEncryptionController = {
      expr = builtins.map isEncryptionController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 16;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsSignalProcessingController = {
      expr = builtins.map isSignalProcessingController [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 17;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsProcessingAccelerator = {
      expr = builtins.map isProcessingAccelerator [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 18;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsNonEssentialInstrumentation = {
      expr = builtins.map isNonEssentialInstrumentation [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 19;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
    testIsCoprocessor = {
      expr = builtins.map isCoprocessor [
        {}
        {base_class = {};}
        {base_class.value = 1;}
        {base_class.value = 64;}
      ];
      expected = [
        false
        false
        false
        true
      ];
    };

    testIsAllOf = {
      expr = [
        (isAllOf [] usbController)
        (isAllOf [isMassStorageController] usbController)
        (isAllOf [isUsbController] usbController)
        (isAllOf [
            isUsbController
            (item: item.driver == "foo")
          ]
          usbController)
        (isAllOf [
            isUsbController
            (item: item.driver == "xhci_hcd")
          ]
          usbController)
        (isAllOf [
            isUsbController
            (item: item.driver == "xhci_hcd")
            (item: item.driver_module == "xhci_pci")
          ]
          usbController)
      ];
      expected = [
        true
        false
        true
        false
        true
        true
      ];
    };

    testIsOneOf = {
      expr = [
        (isOneOf [] usbController)
        (isOneOf [isMassStorageController] usbController)
        (isOneOf [
            isMassStorageController
            isFirewireController
          ]
          usbController)
        (isOneOf [
            isMassStorageController
            isFirewireController
            isUsbController
          ]
          usbController)
      ];
      expected = [
        false
        false
        false
        true
      ];
    };
  }
