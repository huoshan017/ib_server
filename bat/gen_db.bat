set GOPATH=%cd%/../../../../..

cd ../tools
code_generator.exe -c ../conf/db/game_db.json -d ../src/game -p protobuf/protoc.exe
code_generator.exe -c ../conf/db/account_db.json -d ../src/account -p protobuf/protoc.exe
