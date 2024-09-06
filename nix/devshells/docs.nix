{pkgs, ...}:
pkgs.mkShellNoCC {
  packages = with pkgs;
    [
      mkdocs
    ]
    ++ (with pkgs.python3Packages; [
      mike
      mkdocs-material
    ]);
}
