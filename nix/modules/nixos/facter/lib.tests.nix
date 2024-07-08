let
  # todo is there a better way of doing this?
  pkgs = import <nixpkgs> { system = "x86_64-linux"; };
  lib = pkgs.lib.extend (import ./lib.nix);
in
{

  pci = with lib.pci; {
    testIsMassStorageController = {
      expr = isMassStorageController {
        base_class = {
          value = 1;
        };
      };
      expected = true;
    };
    testIsNetworkController = {
      expr = isNetworkController {
        base_class = {
          value = 2;
        };
      };
      expected = true;
    };
    testIsDisplayController = {
      expr = isDisplayController {
        base_class = {
          value = 3;
        };
      };
      expected = true;
    };
    testIsMultimediaController = {
      expr = isMultimediaController {
        base_class = {
          value = 4;
        };
      };
      expected = true;
    };
    testIsMemoryController = {
      expr = isMemoryController {
        base_class = {
          value = 5;
        };
      };
      expected = true;
    };
    testIsBridge = {
      expr = isBridge {
        base_class = {
          value = 6;
        };
      };
      expected = true;
    };
    testIsCommunicationController = {
      expr = isCommunicationController {
        base_class = {
          value = 7;
        };
      };
      expected = true;
    };
    testIsGenericSystemPeripheral = {
      expr = isGenericSystemPeripheral {
        base_class = {
          value = 8;
        };
      };
      expected = true;
    };
    testIsInputDeviceController = {
      expr = isInputDeviceController {
        base_class = {
          value = 9;
        };
      };
      expected = true;
    };
    testIsDockingStation = {
      expr = isDockingStation {
        base_class = {
          value = 10;
        };
      };
      expected = true;
    };
    testIsProcessor = {
      expr = isProcessor {
        base_class = {
          value = 11;
        };
      };
      expected = true;
    };
    testIsSerialBusController = {
      expr = isSerialBusController {
        base_class = {
          value = 12;
        };
      };
      expected = true;
    };
    testIsWirelessController = {
      expr = isWirelessController {
        base_class = {
          value = 13;
        };
      };
      expected = true;
    };
    testIsIntelligentController = {
      expr = isIntelligentController {
        base_class = {
          value = 14;
        };
      };
      expected = true;
    };
    testIsSatelliteCommunicationsController = {
      expr = isSatelliteCommunicationsController {
        base_class = {
          value = 15;
        };
      };
      expected = true;
    };
    testIsEncryptionController = {
      expr = isEncryptionController {
        base_class = {
          value = 16;
        };
      };
      expected = true;
    };
    testIsSignalProcessingController = {
      expr = isSignalProcessingController {
        base_class = {
          value = 17;
        };
      };
      expected = true;
    };
    testIsProcessingAccelerator = {
      expr = isProcessingAccelerator {
        base_class = {
          value = 18;
        };
      };
      expected = true;
    };
    testIsNonEssentialInstrumentation = {
      expr = isNonEssentialInstrumentation {
        base_class = {
          value = 19;
        };
      };
      expected = true;
    };
    testIsCoprocessor = {
      expr = isCoprocessor {
        base_class = {
          value = 64;
        };
      };
      expected = true;
    };
  };

  initrd =
    with lib.initrd;
    let

      sataController = {
        base_class = {
          value = 1;
          name = "Mass storage controller";
        };
        sub_class = {
          value = 6;
          name = "SATA controller";
        };
        driver = "ahci";
        driver_module = "ahci";
      };
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
    {
      testAvailableKernelModules = {
        expr = availableKernelModules {
          hardware = {
            storage_ctrl = [
              sataController
              sataController
              {
                bus = {
                  name = "None";
                };
                base_class = {
                  value = 1;
                  name = "Mass storage controller";
                };
                sub_class = {
                  value = 2;
                  name = "Floppy disk controller";
                };
              }
            ];
            usb_ctrl = [
              usbController
              usbController
              usbController
            ];
            sound = [
              {
                base_class = {
                  value = 4;
                  name = "Multimedia controller";
                };
                sub_class = {
                  value = 3;
                  name = "Audio device";
                };
                driver = "snd_hda_intel";
                driver_module = "snd_hda_intel";
              }
            ];
          };

        };
        expected = [
          "ahci"
          "xhci_pci"
        ];
      };
    };

}
