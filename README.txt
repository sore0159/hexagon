
   \__/  \__/  \__/  \__/  \__/  \__/  \__/  \__/  
   /  \__/  \__/  \__/  \__/  \__/  \__/  \__/  \_
   ================ GOLANG PACKAGE ===============
   /  \__/  \__/  \=== HEXAGON ===  \__/  \__/  \_
   \__/  \__/  \__/  \__/  \__/  \__/  \__/  \__/  
   /  \__/  \__/  \__/  \__/  \__/  \__/  \__/  \_
   \__/  \__/  \__/  \__/  \__/  \__/  \__/  \__/  


Hexagon is a library for hexagon grid geometry.  

It includes a type Coord{int,int} with methods
 such as adding, scaling, stepwise distance, and pathfinding.

Also it has a Pixel{float64, float64} type in conjunction
with a Viewport type to facilitate rendering maps of
hexagon grids.  The Viewport constructor MakeViewport takes
the hex radius (in pixels), and boolians for if the display
should be isometric, and if the display should be 
flat-end up/pointy-end up.

After calling the constructor, setting the anchored hexagon/pixels
and the viewable frame (upper left and lower right pixels) creates
a functional Viewport.  Call VisList() for a slice of the hexagons
within the frame, and for any hexagon the viewport methods CenterOf
and CornersOf give the pixels for rendering.

Finally Viewport has a HexContaining method to translate a pixel
into the hexagon coordinate that contains that pixel.
