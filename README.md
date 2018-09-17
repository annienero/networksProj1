## High Level Approach

Build up our functionality from basic capabilities like sending messages to server, to reading them, etc.
Only move on to next piece of functionality once we are confident we can do the more foundational pieces.
As far as design, we have a main loop that will keep reading and replying to FIND messages until it gets
A BYE message, at which point we print out the secret flag. We have helpers to execute and reply to FIND.

## Challenges We Faced

Figuring out how to read variable length messages from the socket that weren't terminated by EOF. This
was only a hard problem because we are bad at reading the assignment and didn't realize messages were
terminated with newlines.

## How We Tested

We used school wifi or ssh'd to CCIS machines to test. Didn't do anything fancy, just did manual testing.