{
  pname,
  pkgs,
  flake,
  inputs,
  perSystem,
  ...
}:
pkgs.runCommandLocal pname {
  nativeBuildInputs = [perSystem.nix-unit.default];
} ''
  export HOME="$(realpath .)"
  find ${flake} -type f -name "*.unit.nix" \
      -exec nix-unit --eval-store "$HOME" \
      -I nixpkgs=${inputs.nixpkgs} {} \;
  touch $out
''
