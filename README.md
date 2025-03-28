# twitter-bee
A twitter scraper client to simulate a real account and provide a restfull api to receive task.

# How to run
1. register rapidapi account and subscribe to different twitter api.
```
https://rapidapi.com/alexanderxbx/api/twitter-api45/playground/apiendpoint_b4db8ee2-25bc-4743-934d-8b9b4e774bd7
https://rapidapi.com/davethebeast/api/twitter241/playground/apiendpoint_51903177-97a1-4953-b281-d4cb4691e7ac
https://rapidapi.com/Glavier/api/twitter135/playground/apiendpoint_ac816a07-3773-4462-bfd9-0b1f7b58e183
https://rapidapi.com/xtreme-apis-xtreme-apis-default/api/twitter293/playground/apiendpoint_c925b0f2-4f96-473d-a424-ac31cf330bc7
```
2. add your api key to `keylist.json.tml` and rename it to `keylist.json`.
```shell
cp keylist.json.tml keylist.json
```

3. build the service.
```shell
make
```
4. run the service.
```shell
./build/bin/tbee 
```
