# minefield
Random assortment of test code for migrating Elastic -> CH nothing here is set in stone.

#### cmd/csv2CH
This will convert a given CSV that was generated via vspec to a new CH table and matching go struct. This can be run with `go generate` via generate.go

#### cmd/esVsch
This binary will start a gernerate fake vehicle data start a test elastic container and a test clickhouse container then insert the fake vehicle data into both databases in their respective formats that mimick what is in production. The program will wait for ctrl-c to exit so you can exec into the containers and run any test queries you may have.

#### cmd/es-ch 
This a main.go I am using to test various CH client code with the different libraies.

#### internal/protbuf
This contains random code from when we were investigating infering the schema from protobuf.

