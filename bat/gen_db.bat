set GOPATH=%cd%/../../../../..

cd ../conf/db
md proto

cd ../../tools
code_generator.exe -c ../conf/db/game_db.json -d ../src/game -p ../conf/db/proto/game_db.proto
code_generator.exe -c ../conf/db/account_db.json -d ../src/account -p ../conf/db/proto/account_db.proto

cd protobuf
protoc.exe --go_out=../../src/game/game_db --proto_path=../../conf/db/proto game_db.proto
protoc.exe --go_out=../../src/account/account_db --proto_path=../../conf/db/proto account_db.proto
cd ../../proxy/test
