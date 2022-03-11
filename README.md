I was having an issue with homebrew not signing binaries properly.
The crash report for an invalid signature (in Monterey at least) does not include basic information
like which dylib it is that actually has the invalid signature.

I thought I found everything, but was then having problems with part of ffmpeg, which has a lot of dependencies,
so going through them manually is not at all practical.
I whipped this up to traverse the linked binaries and verify their signature.

(it ended up being libharfbuzz)

so this is hardly a mature tool or anything, but is a bit more than just a throwaway shell script
