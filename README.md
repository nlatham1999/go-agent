# go-agent

[![forthebadge](https://forthebadge.com/images/featured/featured-built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/featured/featured-made-with-crayons.svg)](https://forthebadge.com)

Go Agent is a really cool agent based modelling library for go

based off of netlogo

consists of three parts. 
- library to run models
- api complete with frontend to interface with models

For a demo run `go run main.go`

## Model

This is the library that is used to create and interact with a model. For an exhaustive list of all the functions look here: https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md

### High Level Overview

The [Model](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#func-newmodel) struct holds the state of the model and comes with a bunch of functions to interact with the different agents 

There are three layers of agents: [Patches](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#type-patch) which are the stationary grid of the model, [Turtles](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#type-turtle) which can move around, and [Links](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#type-link) which represent relationships between turtles and can be directed or undirected, enabling graph type models where the links are edges and turtles are nodes.

Collections of agents are called agentsets, there are three types, corresponding to the three types of agents [PatchAgentSet](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#type-linkagentset), [TurtleAgentSet](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#TurtleAgentSet), and [LinkAgentSet](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#type-linkagentset)

Breeds are subsets of Turtles and Links and come with their own functions: [TurtleBreed](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#type-turtlebreed) and [LinkBreed](https://github.com/nlatham1999/go-agent/blob/main/pkg/model/doc.md#type-linkbreed)

Coordinate System: The coordinate system follows cartesian cordinates with each integer coordinate being the center of a patch. So patch (1,1) stretches from .5 to 1.5 in both x and y

### Sample

below is a sample of a wolf and sheep model where patches are grass or dirt and the turtle breeds are wolves and sheep

```Go

	sheep := model.NewTurtleBreed("sheep", "", map[string]interface{}{
		"health": 0,
	})
	wolves := model.NewTurtleBreed("wolves", "", map[string]interface{}{
		"hunger": 0,
	})

    // create a model with the breeds and extra properties
    // and patches from -15 to 15 in both x and y for a world of -15.5 to 15.5 in both directions (since patches are centered)
	modelSettings := model.ModelSettings{
		TurtleBreeds: []*model.TurtleBreed{sheep, wolves}, // add the turtle breeds
		TurtleProperties: map[string]interface{}{ // add the turtle properties
			"energy": 0,
		},
		PatchProperties: map[string]interface{}{ // add the patch properties
			"grassOrDirt": "grass",
		},
		MinPxCor:   -15,
		MaxPxCor:   15,
		MinPyCor:   -15,
		MaxPyCor:   15,
		RandomSeed: 10,
	}

	// create the model
	m := model.NewModel(modelSettings)

	// get the agentset of sheep
	sheepAgentSet := sheep.Agents()

	// for each sheep attempt to eat grass and increase the engergy if grass is found
	sheepAgentSet.Ask(
		func(t *model.Turtle) {
			energy := t.GetProperty("energy").(int)
			patch := t.PatchHere()
			if patch.GetProperty("grassOrDirt").(string) == "grass" {
				energy += 2
				patch.SetProperty("grassOrDirt", "dirt")
			} else {
				energy--
			}
			t.SetProperty("energy", energy)
		},
	)

	// for each sheep set the heading to a random direction and move forward in that direction
	sheepAgentSet.Ask(
		func(t *model.Turtle) {
			t.SetHeading(m.RandomFloat(360))
			t.Forward(1)
		},
	)

    // wolf movements
	...
```

### Concurrency
Concurrency is currently supported only for turtles via AskConcurrent, SetPropertySafe and GetPropertySafe  

For simulation forward type models concurrency will not ruin determinism, two sets of conccurent actions can be created, a read group and a write group that happens after all the reads are done. For continuous models however, concurrency ruins determinism.

## TODO
  
-  [x] switch the breedname string params to be the turtlebreed and linkbreed types  
-  [x]  add a function for turtlesinradius at the model level - for speed purposes, should first get all patches, and then turtles on the patch  
    - this should instead be a seperate package on top
    - collision detection layer they can add on
-  [x]  patches, turtles "own" should be renamed to "properties"  
-  [x] there should be a way to pass in the mouse x and y, maybe as a dynamic variable?  
-  [ ] build widgets in js  
-  [ ] implement subtract-headings  
-  [ ] implement dx and dy  
-  [x] make sure that distance works with wrap around on the horizontal  
-  [ ] add slider to change render speed on frontend   
-  [x] documentation   
-  [x] move folder structure so packages are in pkg folder
-  [x] have main screen and then model running on second
-  [x] make link types in threejs render properly
-  [x] concurrency on patches and links
-  [ ] 3D
-  [x] collision detection
-  [ ] labels on turtles and links
-  [ ] setting the color should be a setcolor function
-  [x] make moving on world wrapping on left and down work 🤦‍♂️
-  [x] make setxy able to be called concurrently
-  [x] remove tie modes
-  [x] make change patch ownership able to be called concurrently
-  [x] button widgets
-  [ ] enable setting x,y, height(?) and width on widgets
-  [x] update readme on concurrency
-  [x] remove collision detection
-  [ ] make a seperate package for collision detection
-  [ ] speed improvements on rendering
-  [x] on sliders, if the value is set, than use that value when loading
-  [ ] agentset copy funcs
-  [ ] stats should just be widgets
-  [x] loading unloading model type to enable storing into json
-  [ ] threejs should not be setting the css values
-  [ ] Should be a createLink func at the model level
-  [x] rename pcolor to color
-  [x] labelColor on turtles should be public, same with label
-  [x] seperate concurrency package