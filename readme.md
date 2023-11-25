Implementing simple full text search engine

Flow of text processing

```mermaid
flowchart LR
    text-- tokenizer ---> tokens--processing ---> processed_tokens-- indexer---> inverted_index
```

Processing
```mermaid
flowchart LR
    token-->lowercase-->filter_stop_words-->steming
```

Run example 
```
Loaded 672388 documents in 19.173416973s
Indexed 672388 documents in 11.977604974s
Search found 2 documents in 38.599µs
```



