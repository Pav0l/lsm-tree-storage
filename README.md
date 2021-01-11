# Log-Structured Merge-Tree storage engine

Pet project to play around with Go and LSM tree data store

### DB READS

- read request enters READ QUEUE
- workers (multiple) reads from READ QUEUE
- worker checks for the read value in memtable (red-black tree) (RAM)
    - if not present, check the in memory sparse index for the most recent segment (SSTable), repeat with older segments until found
    - if found, read value from segment based on the sparse index (key:offset)
    - if not found return "null"

### DB WRITES

- write request enters WRITE QUEUE
- worker (single process) reads from WRITE QUEUE
- worker writes data to memtable (red-black tree)
    - if memtable size bigger than X threshold -> write it to disk as SSTable file with it's sparse in-memory index as the latest segment
    - if segment count bigger than Y -> run merge and compaction process in background, then replace and clean up old files

### COMPONENTS

- queue (read, write)
- memtable implementation -> in memory key-value store; limited RAM availability; writes in red-black tree to sort writes
- SSTable (segment) writer -> array of key-value stores from memtable. creates in-memory key-value index of (key: offset)
- memtable manager -> checks the size of memtable and manages when it's written to disk as SSTable and creating of new memtable instance
- segment manager -> monitors the number of segments and initializes merging and compaction; deletes the old segments;
- merge and compact -> merge-sort algorithm that merges multiple segments and compacts them into one; updates sparse index in memory

### (temporary?) SIMPLIFICATIONS

- CSV as file format for simplicity
- Engine crash recovery (no memtable snapshots)
- No prevention of partial writes in case of crash
- No Bloom filter for faster reads
- 15664897489
