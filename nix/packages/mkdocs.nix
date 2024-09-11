{
  pkgs,
  flake,
  ...
}:
pkgs.runCommand "mkdocs"
{
  buildInputs = with pkgs; [
    mkdocs
    python3Packages.mkdocs-material
    python3Packages.mike
  ];
} ''
  mkdir -p $out/bin
  cat <<MKDOCS > $out/bin/mkdocs
  #!${pkgs.bash}/bin/bash
  set -euo pipefail
  export PYTHONPATH=$PYTHONPATH
  export MKDOCS_NIXOS_FACTER=${flake}/mkdocs.yml
  export MKDOCS_NIXOS_FACTER_THEME="${flake}/docs/theme"
  exec ${pkgs.mkdocs}/bin/mkdocs "\$@"
  MKDOCS
  chmod +x $out/bin/mkdocs
''
