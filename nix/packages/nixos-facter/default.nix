{
  inputs,
  perSystem,
  pkgs,
  ...
}:
pkgs.callPackage ./package.nix {
  # The hwinfo flake does not export packages for all of our systems
  # (riscv64-linux). Build its package.nix with our pkgs instead.
  hwinfo = import "${inputs.hwinfo}/nix/package.nix" { inherit pkgs; };
}
// {
  passthru.tests = import ./tests {
    inherit pkgs inputs perSystem;
  };
}
