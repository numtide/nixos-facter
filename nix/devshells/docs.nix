{
  pkgs,
  perSystem,
  ...
}:
pkgs.mkShellNoCC {
  packages = with pkgs;
    [
      perSystem.godoc.default
      (pkgs.writeScriptBin "gen-reference" ''
        out="./docs/content/reference/go_types"
        godoc -c -o $out .
        git add $out
      '')
      (pkgs.writeScriptBin "mkdocs" ''
        # generate reference docs first
        gen-reference
        # execute the underlying command
        ${pkgs.mkdocs}/bin/mkdocs "$@"
      '')
    ]
    ++ (with pkgs.python3Packages; [
      mike
      mkdocs-material
    ]);
}
