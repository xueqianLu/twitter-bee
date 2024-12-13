# twitter-bee
A twitter scraper client to simulate a real account and provide a restfull api to receive task.

# How to run
1. prepare `userlib.json`, see `userlib.json.tml` for example.
```shell
cp userlib.json.tml userlib.json
```
Fill in the `userlib.json` with your twitter account information.

2. build the service.
```shell
make
```
3. run the service.
```shell
./build/bin/twitter-bee --user <username in userlib.json>
```
# Advance
1. You can use a proxy to avoid being banned by twitter.
```shell
./build/bin/twitter-bee --user <username in userlib.json> --proxy http://127.0.0.1:2340
```
