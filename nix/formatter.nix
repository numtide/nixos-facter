{
  flake,
  pkgs,
  system,
  inputs,
  ...
}:
inputs.treefmt-nix.lib.mkWrapper pkgs {
  projectRootFile = ".git/config";

  programs = {
    prettier.enable = true;
    nixfmt-rfc-style.enable = true;
  };

  settings.formatter.prettier = {
    options = [
      "--tab-width"
      "4"
    ];
    includes = [
      "*.css"
      "*.html"
      "*.js"
      "*.json"
      "*.jsx"
      "*.md"
      "*.mdx"
      "*.scss"
      "*.ts"
      "*.yaml"
    ];
  };
}
