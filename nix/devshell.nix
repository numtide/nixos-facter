{
  pkgs,
  perSystem,
  ...
}:
pkgs.mkShell {
  GOROOT = "${pkgs.go}/share/go";

  packages = with pkgs; [
    go
    gotools
    enumer
    delve
    pprof
    graphviz
    libusb1.dev
    gcc
    pkg-config
    util-linux.dev
    pciutils
    hwinfo
    perSystem.gomod2nix.default
    golangci-lint
    cobra-cli
  ];
}
