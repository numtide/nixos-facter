# nixos-facter

> [!NOTE]
> **Status: alpha**

Instead of generating Nix code we would be generating a JSON-like data structure of facts about the machine.
This can be recorded and loaded by Nix. Or added to the nixos-hardware repo.
It can be used for user diagnostics.
It inverts the control where the NixOS module system can now make decisions on which kernel modules to load over time.

## How the project works

The project has two parts:

-   The binary that scans the system and outputs JSON or TOML.
-   A series of [NixOS modules](https://github.com/numtide/nixos-facter-modules) that can load that data and make decisions.

That's it.

## Quick Start

To generate a report:

```console
# you must run as root
❯ sudo nix run github:numtide/nixos-facter -- -o report.json

# you can use fx to view the report in the terminal
❯ fx report.json 
```

## Some more considerations

The generated data should be generic enough to be embeddable in public Git repos.
We can show things like hardware models, what type of graphics card is running, ...
Avoid serial numbers and MAC addresses.

Another way to look at it is that the generated data should be the same if two people have the same model of hardware.

## Downsides of this approach

The main downside is that it adds a layer of indirection. New hardware detection requires updating both the tool and the Nix schema.
