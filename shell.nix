{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  name = "hugo-shell";
  buildInputs = [
    pkgs.go_1_23
    pkgs.git
  ];
   shellHook = ''
  '';
}
