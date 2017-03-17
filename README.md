# my-hash-map

Performance:

Average read speed compared to STD-Lib for 178693 words anagrams:
15% faster if no cache misses
15% slower if all cache misses.

Write / Build speed: 2x slower than STD-Lib (although haven't optimised this)

Memory:
Currently it actually uses **LESS memory than would be required to store the keys themselves**.
This is because it stores only the parts of the keys that are needed to resolve collisions.
To allow for inserts, we would need to store the whole key, and upon insertion, figure out if we need an extra bit in bitNums to distinguish.



Currently assumes list of keys is known ahead of time, although simple enough changes could be made to allow growing, inserts etc.

It works on 64 bit numbers for now.
It buckets 64bit numbers using the first 20 bits.
If a bucket has collisions, it figures out, of this group of colliding numbers, what is the minimum number of bits we need to look at of each of these numbers to be able to distinguish between them.
see http://stackoverflow.com/questions/37294648/algorithm-minimal-number-of-bits-to-distinguish-a-set-of-given-binary-numbers/42714847#42714847 for example
if we had We can use the key and the important bits to create an offset to lookup within the bucket.
to generate the offset, we sum **2^i**, only when the key contains the **ith** important bit

I stored these offsets in **BitScoreCache**, but it improved performance by less than 5%. However, making this cache a special readonly piece of memory that's shared by the operating system, assuming that the hashmap was used by multiple programs, this could speed up the hashmap and be constant memory footprint.


Note: runtime hashmap is here: https://golang.org/src/runtime/hashmap_fast.go?s=2939:3009#L100

Why it's hard to beat the standard library:
https://www.goinggo.net/2013/12/macro-view-of-map-internals-in-go.html

https://github.com/golang/go/blob/master/src/runtime/hashmap.go

would be interesting to bench against minimal perfect hashing e.g. https://github.com/alecthomas/mph

good link explaining std-lib implementation https://www.youtube.com/watch?v=Tl7mi9QmLns