# Pokedex CLI 

A REPL (Read-Eval-Print Loop) tool using PokedexAPI to explore and catch Pokémon in your Pokédex, with caching for faster access.

## Getting Started

### 1. Start the CLI
```sh
go build && ./pokedexcli
```

### 2. View Help
```sh
pokedex >help
```

## Commands

- **inspect {pokemon_name}**: 🔍 Get information about a Pokémon.
- **pokedex**: 📖 View all the Pokémon in your Pokédex.
- **exit**: 🚪 Exit the Pokedex.
- **help**: ℹ️ Display a help message.
- **map**: 🗺️ Display the names of 20 location areas in the Pokémon world. Each subsequent call to `map` shows the next 20 locations.
- **mapb**: 🔙 Similar to the `map` command, but displays the previous 20 locations, allowing you to go back.
- **explore {location_area}**: 🌲 List the Pokémon in a location area.
- **catch {pokemon_name}**: 🎣 Attempt to catch a Pokémon and add it to your Pokédex.

Happy exploring and catching! 🐾