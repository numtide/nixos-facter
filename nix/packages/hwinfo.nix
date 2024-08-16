{pkgs, ...}:
# we use this fork for updated pci and usb id resolution until https://github.com/openSUSE/hwinfo/pull/144 is
# merged upstream
pkgs.hwinfo.overrideAttrs (old: {
  src = pkgs.fetchFromGitHub {
    owner = "brianmcgee";
    repo = "hwinfo";
    rev = "6fa35c4e8404217a6dc57cff1f5639f7b596be18";
    hash = "sha256-HYLoTArnwZsIVzNFtPeJus9inDKx3kYJSKWIu5IPp14=";
  };

  # patch runtime dependencies
  # todo remove once https://github.com/NixOS/nixpkgs/pull/334633 has been merged into nixpkgs
  postPatch =
    old.postPatch
    + ''
      substituteInPlace src/hd/hd_int.h --replace "/sbin/" ""
      substituteInPlace src/hd/hd_int.h --replace "/usr/bin/" ""
    '';

  nativeBuildInputs =
    old.nativeBuildInputs
    ++ [
      pkgs.makeBinaryWrapper
    ];

  postFixup = let
    runtimeDeps = with pkgs; [
      kmod # modprobe
      systemdMinimal # udevadm
    ];
  in ''
    wrapProgram $out/bin/hwinfo \
        --prefix PATH : ${pkgs.lib.makeBinPath runtimeDeps}
  '';
})
