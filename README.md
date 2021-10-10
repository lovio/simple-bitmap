# simple-bitmap

bitmap(位图) 是一种常见的数据结构，在 introsect(求交)，union(求并)计算时有非常好
的性能。常用语大量数据的快速排序、快速去重、快速查找等场景。

bitmap 有很多的实现，比如 redis 就内置了 bitmap 的数据结构。也有很多的库实现了
bitmap。

我们在选择 bitmap 的时候需要考虑具体的使用场景。如果你的数据十分稀疏，那么使用传
统的 bitmap 就会非常浪费内存。这个时候我们就需要靠考虑 compressed bitmap 来高效
利用空间了。比如 roaringbitmap 就是一种，他从 Bitmap 的一层连续存储转换为一个二
级的存储结构（Chunk + Container）。

## ThreadSafe

在实现的时候底层存储使用了 byte，并发情况下进行 SetBit 是会出现 lost update 的，
所以我们通过 atomic 的 CAS 来解决这个问题。比 RWLock 要高效很多。既然使用了
atomic，那么 unsafe 的方法可能必要性就没有那么高了。

## TODO

- [x] 使用 RWLock 的效率太低了，改成 atomic 之后，需要把存储的类型从 byte 改为
      uint32，这样就可以直接使用 atomic.LoadUint32
- [x] Remove thread unsafe bitmap
