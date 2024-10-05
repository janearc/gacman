# gacman
unity front end, go back end (that's this!), what could go wrong?

# some quick notes

the hierarchy of stuff in this environment is:
- a Space (this contains all things)
  - a Dungeon (this is a set of Levels)
    - a Level (this is a set of Rooms)
      - a Room (this is a set of Cells)
        - a Cell (this can contain Shapes such as Entity or Stairs)
          - an Entity (this can be just about anything including an attack helicopter)

Additionally, there are:
- Chunks (these are used to store and retrieve Cells)
- Vector3s (these are used to store positions in 3D space and are used to locate Cells and Entities)

At present, this is just communicated in json over a websocket. But materially, this is procedurally
generated per interaction, and in this way resembles Nethack or Dwarf Fortress. Each Level
has an "up" and a "down" staircase, and the cursor can move between Levels by traversing
the stairs. Each level in the dungeon is remembered (per-instance) in the Space object
(which is not visible to the cursor/player). Theoretically, many Dungeons could exist and
there could be magical stairways or some kind of nefarious wizard or whichever that teleported
you to a different dungeon and set of levels, but it's not implemented, and I'm not building
a roguelike. At least, not necessarily.

The state this software is in presently proves procedural generation and logical containment
of a set of 2D maps in a 3D structure that exist within separate contexts in a larger space
(you might think of it as a "universe" or something). And, because all of this is *known*,
and we can communicate over a websocket, all of this same information can be communicated to
Unity in my cleverly named [What Could Go Wrong](https://github.com/janearc/wcgw) project. So, kind of,
this is a go backend for a unity front end, that replicates a nearly fifty year old game (Nethack
was released in 1987, and Rogue was released in 1980). [It's not stupid, it's *advanced*](https://www.youtube.com/watch?v=aiGjK64z7KI).

Anyways, it's just a toy, I had been thinking to myself, "what's something I absolutely should not, can not do in golang"
and the first thing that came to mind was "first person shooter/3d environment." And because I don't
like being told I can't do things, I decided to build this.

# how do I run this?

Conveniently, I have placed a makefile for you. You should just be able to run `make`, and
you'll have two binaries in `build/{client,server}`. Surprisingly, just running `./build/server`
will run the server and `./build/client` will run the client. At the moment, the client 
expects cardinal directions, and the server responds by moving your cursor around in your
level/dungeon/space. I don't plan to keep this text/console interactive layer around for very
long, but I might keep some kind of relic there for debugging purposes (for example, `./client --show-map`).

Have fun? I guess? What kind of person wants to use this?