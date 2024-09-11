{
  pkgs,
  perSystem,
  ...
}:
pkgs.mkShellNoCC {
  packages = [
    perSystem.self.mkdocs
    pkgs.python3Packages.mike
  ];
}
