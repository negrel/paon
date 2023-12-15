<h1 align="center">
    <img height="250" src="https://github.com/negrel/Animated-Fluent-Emojis/blob/master/Emojis/Animals/Peacock.png?raw=true">
</h1>

<p align="center">
	<a href="https://pkg.go.dev/github.com/negrel/paon">
		<img src="https://godoc.org/github.com/negrel/paon?status.svg">
	</a>
	<a href="https://goreportcard.com/report/github.com/negrel/paon">
		<img src="https://goreportcard.com/badge/github.com/negrel/paon">
	</a>
<a href="https://app.fossa.com/projects/git%2Bgithub.com%2Fnegrel%2Fpaon?ref=badge_shield" alt="FOSSA Status"><img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnegrel%2Fpaon.svg?type=shield"/></a>
	<a href="https://github.com/negrel/paon/raw/master/LICENSE">
		<img src="https://img.shields.io/github/license/negrel/paon">
	</a>
</p>

# :peacock: - Paon
*A framework to build efficient, thread-aware, highly composable widgets based terminal user interface.*

> **NOTE**: Paon is still under active development and is not ready yet.  
> **NOTE**: It's coming soon. :eyes:

## Widgets
- [ ] Text
- [ ] TextInput
- [ ] PasswordInput
- [ ] Button
- [ ] Drop-down
- [ ] Checkbox
- [ ] Radio

## Layouts
- [ ] HBox
- [ ] VBox
- [ ] Table
- [ ] ScrollView
- [ ] Form

# Project structure

```shell
┌─[negrel@matebook] - [~/code/golang/paon] - [4792]
└─[$] tree . -d
.
├── events                             # Public events
├── examples                           # Code examples
│   └── ...
├── internal                           # internal packages
│   └── ...
├── pdk                                # Paon Development Kit | Packages used to develop custom widgets
│   ├── backend                          # Interface to abstract terminal/console interaction.
│   │   └── tcell                          # tcell implementation
│   ├── draw                             # Drawing interfaces
│   ├── id                               # A package to generate unique ID (thread-safe)
│   ├── layout                           # Layout interfaces
│   ├── math                             # Basic numerical helpers
│   ├── tree                             # Generic node tree package
│   └── widgets                          # Widgets definition and basic implementations
├── scripts
├── styles                             # Widgets styling package
│   ├── property                         # Style property interface
│   └── value                            # Style value types
└── widgets                            # Built-in widgets (see list above)

```

## :stars: Show your support

Please give a :star: if this project helped you!

## :scroll: License

Apache2.0 © [Alexandre Negrel](https://www.negrel.dev/)


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnegrel%2Fpaon.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnegrel%2Fpaon?ref=badge_large)
