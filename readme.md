# Alien Invasion!
An exercise in Go to simulate a mad alien invasion.

## Installation and usage
Create a map file and run the simulation:
```
go install
./bin/invade MAPFILE [-n ALIENS]
```

## End conditions
A simulation ends once one of the conditions has been reached:
* No more aliens in the simulation
* Each alien has reached 10,000 moves

## Map files
Define a map as a text file with a line for each city.
Cities can have a neighbouring city in any cardinal direction, north, east, south and west.

### Map file syntax: 
```
NAME [north=NAME] [east=NAME] [south=NAME] [west=NAME]
```

### Example:
Define a map with 4 cities:  
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```    
![Map](doc/map_rendered.png)

## Open questions and assumptions
> If cities are defined neighbours by implication only (i.e. defined as neighbours by their neighbours but not directly by themselves), there is no route between them.  
  
Example: Helsinki has no route to Tallinn despite being direct neighbours by implication.
```
Helsinki east=Vyborg
Vyborg south=St.Petersburg west=Helsinki
St.Petersburg north=Vyborg east=Tallinn
Tallinn east=St.Petersburg
```



