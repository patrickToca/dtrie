# dtrie
Dynamic Persistent Hash Array Mapped Trie

a persistent hash trie that dynamically expands or shrinks to provide efficient memory allocation.

## Big O
- O(log32(n)) get, remove, and update
- O(n) insertion

## Based on the following papers and talks:
- Ideal Hash Trees by Phil Bagwell
- Optimizing Hash-Array Mapped Tries for Fast and Lean Immutable JVM Collections by Michael J. Steindorfer and Jurgen J. Vinju
- Extreme Cleverness by Daniel Spiewak
