# Alien Invasion!
An exercise in Go to simulate a mad alien invasion.

## Installation and usage
Create a map file and run the simulation:
```
go install
./bin/invade MAPFILE [-n ALIENS]
```


## Map files
Define a map as a text file with a line for each city.
Cities can have a neighbouring city in any cardinal direction, north, east, south and west.

### Map file syntax: 
```
NAME [north=NAME] [east=NAME] [south=NAME] [west=NAME]
```

### Example:
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```

Defines a map with 4 cities:

![Map](doc/map_rendered.png)

