{
  perSystem,
  pkgs,
  ...
}:
perSystem.self.nixos-facter.overrideAttrs (old: {
  GOROOT = "${old.go}/share/go";
  nativeBuildInputs =
    old.nativeBuildInputs
    ++ [
      perSystem.gomod2nix.default
      pkgs.enumer
      pkgs.delve
      pkgs.pprof
      pkgs.golangci-lint
      pkgs.cobra-cli
      perSystem.self.hwinfo
    ];
  shellHook = ''
    # this is only needed for hermetic builds
    unset GO_NO_VENDOR_CHECKS GOSUMDB GOPROXY GOFLAGS
  '';
})
