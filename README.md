# go-protobuf-flatbuffers-benchmarks
An exploration of protocol buffers and flatbuffers with Go.

| data                          | total fields | set fields | len  | marshal (ns/op) | unmarshal (ns/op) |
|-------------------------------|--------------|------------|------|-----------------|-------------------|
| protobuf nested               | 4            | 4          | 18   | 164.4           | 142.3             |
| protobuf flat                 | 4            | 4          | 19   | 136.1           | 90.95             |
| protobuf large                | 119          | 119        | 565  | 1450            | 1080              |
| protobuf xlarge               | 599          | 599        | 3109 | 7545            | 5826              |
| protobuf xlarge partial 8     | 599          | 8          | 35   | 4619            | 226.6             |
| protobuf xlarge partial 64    | 599          | 64         | 266  | 5004            | 692.6             |
| protobuf grouped              | 600          | 600        | 2797 | 8339            | 8140              |
| protobuf grouped partial      | 600          | 64         | 293  | 1242            | 1012              |
| protobuf grouped large        | 2480         | 140        | 652  | 2264            | 1769              |
| flatbuffers nested            | 4            | 4          | 72   | 365.0           | 0.3516            |
| flatbuffers flat              | 4            | 4          | 64   | 363.3           | 0.3541            |
| flatbuffers large             | 119          | 119        | 1040 | 2358            | 0.3528            |
| flatbuffers xlarge            | 599          | 599        | 5072 | 11748           | 0.3512            |
| flatbuffers xlarge partial 8  | 599          | 8          | 96   | 1049            | 0.3518            |
| flatbuffers xlarge partial 64 | 599          | 64         | 584  | 1744            | 0.3562            |
