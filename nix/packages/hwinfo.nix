{pkgs, ...}:
# we use this fork for updated pci and usb id resolution until https://github.com/openSUSE/hwinfo/pull/144 is
# merged upstream
pkgs.hwinfo.overrideAttrs {
  src = pkgs.fetchFromGitHub {
    owner = "brianmcgee";
    repo = "hwinfo";
    rev = "6fa35c4e8404217a6dc57cff1f5639f7b596be18";
    hash = "sha256-HYLoTArnwZsIVzNFtPeJus9inDKx3kYJSKWIu5IPp14=";
  };
}
