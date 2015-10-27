
   \__/  \__/  \__/  \__/  \__/  \__/  \__/  \__/  
   /  \__/  \__/  \__/  \__/  \__/  \__/  \__/  \_
   ================ GOLANG PACKAGE ===============
   /  \__/  \__/  \=== HEXAGON ===  \__/  \__/  \_
   \__/  \__/  \__/  \__/  \__/  \__/  \__/  \__/  
   /  \__/  \__/  \__/  \__/  \__/  \__/  \__/  \_
   \__/  \__/  \__/  \__/  \__/  \__/  \__/  \__/  


Hexagon is a library for hexagon grid geometry.  

It includes a type Coord{int,int} with basic methods
 such as adding, subtracting, and scaling.  Coord has
geometric methods for determining if a value lies
on a line from the origin, determining which sector
a value is in, stepwise distance, creating a ring, 
and pathfinding.

It includes a type Polar{int, int} representing 
polar coordinates on a hexagon grid.  Polar and
Coord have methods to convert to each other (with
0,0 = 0,0).  Polar does not have as robust a feature
set as Coord, but is nice for certain kinds of
calculations.

It has a Pixel{float64, float64} type in conjunction
with a Viewport type to facilitate rendering maps of
hexagon grids.  The Viewport constructor MakeViewport takes
the hex radius (float64), and boolians for if the display
should be isometric, and if the display should be 
flat-end up/pointy-end up.

After calling the constructor, setting the anchored hexagon/pixels
and the viewable frame (upper left and lower right pixels) creates
a functional Viewport.  Call VisList() for a slice of the hexagons
within the frame, and for any hexagon the viewport methods CenterOf
and CornersOf give the pixels for rendering.

Finally Viewport has a HexContaining method to translate a pixel
into the hexagon coordinate that contains that pixel.

Design Calls:
Viewport has an inverted y-axis for pixels: higher y Coords will 
have lower y pixel values.  This is because of how golang's image
library works.

The Coordinate system uses the three directions {0,1}, {1,0}, 
and {1,-1}.  An alternate approach would be for  {1,-1} to be
 replaced with {1,1}, but I found certain computations worked 
better with {1,-1}. The grid uses an "axial" coordinate
system; offset coordinate systems look foolish to me.

The Sector and Axis features of Coord are nonstandard, but I
think they will be pretty useful.

Pathfinding does not use the slick "draw a line then check which
hexes it goes through" method, which might be better, but I
didn't learn of it until late, and wanted Coords to be able to
create paths without using Pixels (as Pixels are implimented to
be dependant on Viewports for meaning).
