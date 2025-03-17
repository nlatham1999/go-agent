# go-agent

agent based modelling library for go based off of netlogo

consists of three parts. 
- library to run models
- api to interface with models

## TODO
  
- [x] switch the breedname string params to be the turtlebreed and linkbreed types  
-  [ ]  add a function for turtlesinradius at the model level - for speed purposes, should first get all patches, and then turtles on the patch  
    - this should instead be a seperate package on top
    - collision detection layer they can add on
-  [x]  patches, turtles "own" should be renamed to "properties"  
-  [ ] there should be a way to pass in the mouse x and y, maybe as a dynamic variable?  
-  [ ] build widgets in js  
-  [ ] implement subtract-headings  
-  [ ] implement dx and dy  
-  [ ] make sure that distance works with wrap around on the horizontal  
-  [ ] add slider to change render speed on frontend   
-  [ ] documentation   
-  [ ] move folder structure so packages are in pkg folder
-  [x] have main screen and then model running on second
-  [ ] make link types in threejs render properly