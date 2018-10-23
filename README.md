# Flopper - pick and place massaging tool

This is a tool to rearrange the PnP output from Altium into a coordinate
system the PnP software can work with. The coordinate system Altium uses is
usually not the one you'd want to use when running the PnP machine. This tools
allows you to arbitrarily rotate the X, Y coordinates for each component into
whatever orientation you need. Use the offset parameters to adjust the coordinates
since the points will be rotated around (0, 0).

To rotate something 90 degrees CCW:

````
./flopper -input_file <input file> -output_file sample90.csv -rotate 90 -xoffset <board width> -yoffset 0
````

...and to rotate something 90 degrees CW:

```
./flopper -input_file <input file> -output_file sample90cw.csv -rotate -90 -xoffset 0 -yoffset <board height>
````

