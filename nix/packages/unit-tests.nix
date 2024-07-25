{
  pname,
  pkgs,
  flake,
  inputs,
  ...
}:
pkgs.runCommandLocal pname {
  nativeBuildInputs = [pkgs.nix-unit];
} ''
  export HOME="$(realpath .)"
  find ${flake} -type f -name "*.unit.nix" \
      -exec nix-unit --eval-store "$HOME" \
      -I nixpkgs=${inputs.nixpkgs} {} \;
  touch $out
''
