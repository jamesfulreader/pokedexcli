# Pokedex CLI

A command-line interface Pokedex application written in Go that interacts with the PokeAPI. This application simulates the experience of exploring the Pokemon world, discovering Pokemon in different locations, and attempting to catch them.

## Features

The Pokedex CLI provides several commands that allow you to:

- Explore different location areas in the Pokemon world
- View Pokemon that can be encountered in each location
- Attempt to catch Pokemon with a probability-based catching system
- Inspect detailed information about caught Pokemon
- View your collection of caught Pokemon

## Installation

To install the Pokedex CLI, you'll need Go installed on your system. Then follow these steps:

```bash
# Clone the repository
git clone https://github.com/jamesfulreader/pokedexcli.git

# Navigate to the project directory
cd pokedexcli

# Build the application
go build

# Run the application
./pokedexcli
```

## Commands

Here are the available commands in the Pokedex CLI:

- `help`: Displays a list of all available commands
- `map`: Shows the next 20 locations in the Pokemon world
- `mapb`: Shows the previous 20 locations
- `explore [location]`: Lists all Pokemon that can be found in a specific location
- `catch [pokemon]`: Attempts to catch a Pokemon
- `inspect [pokemon]`: Shows detailed information about a caught Pokemon
- `pokedex`: Lists all Pokemon you have caught
- `exit`: Exits the application

## Usage Examples

Here's how to use some of the main features:

### Exploring Locations

```bash
Pokedex > map
// Displays 20 location names
Pokedex > explore canalave-city
// Shows Pokemon that can be found in Canalave City
```

### Catching Pokemon

```bash
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
```

### Inspecting Pokemon

```bash
Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  -defense: 40
  -special-attack: 50
  -special-defense: 50
  -speed: 90
Types:
  - electric
```

## Technical Details

### Project Structure

The project is organized into several packages:

- `cmd`: Contains command implementations and configuration
- `internal/pokeapi`: Handles communication with the PokeAPI
- `internal/pokecache`: Implements caching for API responses
- `internal/models`: Defines shared data structures
- `internal/pokedex`: Manages the user's Pokemon collection

### Key Features Implementation

1. **API Integration**: The application uses the PokeAPI (https://pokeapi.co/) to fetch Pokemon data. All API calls are handled through a custom client that includes error handling and response caching.

2. **Caching System**: To improve performance and reduce API calls, the application includes a custom caching system that stores API responses for 5 minutes.

3. **Concurrent Safety**: The cache and Pokedex implementations use mutex locks to ensure thread safety when accessing shared resources.

4. **Pokemon Catching Mechanism**: The catching system uses a probability-based approach where the catch rate is influenced by the Pokemon's base experience points.

## Error Handling

The application includes robust error handling for various scenarios:

- Invalid commands
- Network errors during API calls
- Invalid location or Pokemon names
- Attempting to inspect uncaught Pokemon

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Pokemon data provided by [PokeAPI](https://pokeapi.co/)
- Inspired by the Pokemon game series by Game Freak and Nintendo

## Future Enhancements

Potential future improvements could include:

- Adding Pokemon battles between caught Pokemon
- Implementing a leveling system
- Adding evolution mechanics
- Storing caught Pokemon persistently
- Adding more detailed location information
- Implementing Pokemon trading functionality
