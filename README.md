# Github Gist Query Application
This application queries the Github API and retrieves the public gists for a particular user. When a new public gist is available for the user, a desktop notification is trigered notifying you that a new public gist is available.

## Run the application
At the root of this directory there are binaries included for windows, linux and OSX Operating Systems.

The following example runs the application on OSX:
```
./github-gist-query-osx
2019/08/27 12:58:14 Started Github Public Gist Query Application on port [:8080]
```


## How to query public gists of a github user:
There are 3 ways to query this application to retrieve the public gists of a user namely:

- Via a browser
- Via an API client
- Via the command line

### browser

Go to a broswer window and enter the url `http://localhost:8080`

Enter your github username and click submit. All the public gists for that username would be returned as shown here:

![Alt text](query.png?raw=true "public gists results")

Next go to github gists and create a new github `public` gist.

Within 10 seconds, you should receive a desktop message notifying you that a new github public gist is available like the following:


![Alt text](notification.png?raw=true "desktop notification")

You can refresh the application page in the browser if it is still open, or re-submit the query.

### API client

The application exposes a simple API endpoint `http://localhost:8080/api` that a http API client can query to retrieve the public gists for a particular github user. The following example uses curl to query the API:
```
curl -X POST -H "Content-Type: application/json" --silent  -d '{"user":"pidah"}' http://localhost:8080/api  | jq
[
  "https://gist.githubusercontent.com/pidah/a8f672883901768e2cc06ddd6b0e0348/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/test.json",
  "https://gist.githubusercontent.com/pidah/72853c10800a4a19eda480187a3d5efb/raw/66d97484f3adb27d09a7cfa58eac4aeb1a621070/test.yaml",
  "https://gist.githubusercontent.com/pidah/da1e2607d4c510ede51dde6e6a633dd5/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/test.c",
  "https://gist.githubusercontent.com/pidah/09d1d7f1a4c679d78445e430d69833fd/raw/30d74d258442c7c65512eafab474568dd706c430/test.java",
  "https://gist.githubusercontent.com/pidah/b46ba9190fc90c6a56a27551ba610d3c/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/test.xml",
  "https://gist.githubusercontent.com/pidah/0ed118e9430af45852dcc02ad2076f04/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/testing2.go",
  "https://gist.githubusercontent.com/pidah/f1019fb170ce38b1efa731dbc20e1344/raw/6d5cc1559f7d179c3f12d1017e2de0522e7bcabe/testing.go",
  "https://gist.githubusercontent.com/pidah/0d37ba9575b4a33b4d9060eb4d8e7fee/raw/b92196d03ab98f62ace625bace5ee987687ab0d9/gistfile1.txt",
  "https://gist.githubusercontent.com/pidah/26b66e1a516bf8ad33df5e75a1d07c1a/raw/9a8dc5b16600500552b69a59ebf508fb07b8e12e/gistfile1.txt",
  "https://gist.githubusercontent.com/pidah/cc015785914ed77299f2e1527020d5fe/raw/90031ee71b577f4539f9e0d0f13def876f9ec396/Kafka-minikube-ingress",
  "https://gist.githubusercontent.com/pidah/c208c56d7258d5e44589fab013c136ae/raw/2eafcf52fbb157382122d6a2136c7c4bb5b1eb40/minikube-kafka",
  "https://gist.githubusercontent.com/pidah/6a581a7633129fdf428184c2769df159/raw/6bc693be73bc788cfef0ec82b0ecdb115491ba3a/signal.go",
  "https://gist.githubusercontent.com/pidah/6ccc0f2a76c1731ece020eb2dc3bdd3a/raw/fb4816a8c0d122e90a17d7fc44bdead343d01110/gistfile1.txt",
  "https://gist.githubusercontent.com/pidah/2393a02487642d818031/raw/53efd2da003eba89a7f0cb3a5096ad5cd7847b13/gistfile1.txt",
  "https://gist.githubusercontent.com/pidah/b00267aff067fa1e3064/raw/c85e9045ce67ab23e43a59d3980fc21bd63a64eb/gistfile1.txt",
  "https://gist.githubusercontent.com/pidah/610169767f7335b4a8c9/raw/294ddb80b2cda05b2c031ceab2f0d3fa9086a37b/keybase.md",
  "https://gist.githubusercontent.com/pidah/0bc52f80ebda826b641a/raw/61efcd25b29733da556ec780c9865451ace044aa/keybase.md",
  "https://gist.githubusercontent.com/pidah/10369211/raw/cc95bf377a3ad942501f0c1770df8addc60dab44/gistfile1.txt",
  "https://gist.githubusercontent.com/pidah/5974672/raw/709d7624dfa4fcffcec86f52d61a49ed36428ea8/multi_aws_region_vagrantfile"
]
```
Create a new public gist for the user. Within 10 seconds, you should receive a desktop message notifying you that a new github public gist is available for the github  user. Re-run the API client query to retrieve the updated results.

### Command line

You can also run the query via the command line for a particular user by passing the github username as a single argument as follows:
```
peteridah$ ./github-gist-query-osx pidah
2019/08/27 13:24:35 public gists for pidah: [
    "https://gist.githubusercontent.com/pidah/a8f672883901768e2cc06ddd6b0e0348/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/test.json",
    "https://gist.githubusercontent.com/pidah/72853c10800a4a19eda480187a3d5efb/raw/66d97484f3adb27d09a7cfa58eac4aeb1a621070/test.yaml",
    "https://gist.githubusercontent.com/pidah/da1e2607d4c510ede51dde6e6a633dd5/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/test.c",
    "https://gist.githubusercontent.com/pidah/09d1d7f1a4c679d78445e430d69833fd/raw/30d74d258442c7c65512eafab474568dd706c430/test.java",
    "https://gist.githubusercontent.com/pidah/b46ba9190fc90c6a56a27551ba610d3c/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/test.xml",
    "https://gist.githubusercontent.com/pidah/0ed118e9430af45852dcc02ad2076f04/raw/9a2c7732fab5bcd73ea3ed52d2d9599a4cc47666/testing2.go",
    "https://gist.githubusercontent.com/pidah/f1019fb170ce38b1efa731dbc20e1344/raw/6d5cc1559f7d179c3f12d1017e2de0522e7bcabe/testing.go",
    "https://gist.githubusercontent.com/pidah/0d37ba9575b4a33b4d9060eb4d8e7fee/raw/b92196d03ab98f62ace625bace5ee987687ab0d9/gistfile1.txt",
    "https://gist.githubusercontent.com/pidah/26b66e1a516bf8ad33df5e75a1d07c1a/raw/9a8dc5b16600500552b69a59ebf508fb07b8e12e/gistfile1.txt",
    "https://gist.githubusercontent.com/pidah/cc015785914ed77299f2e1527020d5fe/raw/90031ee71b577f4539f9e0d0f13def876f9ec396/Kafka-minikube-ingress",
    "https://gist.githubusercontent.com/pidah/c208c56d7258d5e44589fab013c136ae/raw/2eafcf52fbb157382122d6a2136c7c4bb5b1eb40/minikube-kafka",
    "https://gist.githubusercontent.com/pidah/6a581a7633129fdf428184c2769df159/raw/6bc693be73bc788cfef0ec82b0ecdb115491ba3a/signal.go",
    "https://gist.githubusercontent.com/pidah/6ccc0f2a76c1731ece020eb2dc3bdd3a/raw/fb4816a8c0d122e90a17d7fc44bdead343d01110/gistfile1.txt",
    "https://gist.githubusercontent.com/pidah/2393a02487642d818031/raw/53efd2da003eba89a7f0cb3a5096ad5cd7847b13/gistfile1.txt",
    "https://gist.githubusercontent.com/pidah/b00267aff067fa1e3064/raw/c85e9045ce67ab23e43a59d3980fc21bd63a64eb/gistfile1.txt",
    "https://gist.githubusercontent.com/pidah/610169767f7335b4a8c9/raw/294ddb80b2cda05b2c031ceab2f0d3fa9086a37b/keybase.md",
    "https://gist.githubusercontent.com/pidah/0bc52f80ebda826b641a/raw/61efcd25b29733da556ec780c9865451ace044aa/keybase.md",
    "https://gist.githubusercontent.com/pidah/10369211/raw/cc95bf377a3ad942501f0c1770df8addc60dab44/gistfile1.txt",
    "https://gist.githubusercontent.com/pidah/5974672/raw/709d7624dfa4fcffcec86f52d61a49ed36428ea8/multi_aws_region_vagrantfile"
]

```

Again if you create a new public gist for the user, within 10 seconds you will receive a desktop notification and a message in the log as follows:
```
2019/08/27 13:28:36 New Public gist found for pidah !
```

## Design Considerations

- As the requirement was only to retreive the public gists for a particular github user, this implementation does not support authentication against the github API with personal tokens or Basic Auth.

- This implementation polls the github API every 10 seconds to check for new Public gists for a previously queried github username. A better approach would be to use a Github webhook to trigger a http POST request to the application when a new public gist is detected. Unfortunately Github Gists API do not support Webhooks like Github Repositories API.

- For each user queried, an entry is made into a map (dictionary) object in memory, mapping the username to the number of publicly available gists for the user. Every 10 seconds, a new query is made to the github API to retrieve an updated public gist count for that user. If the new count value is greater than the existing value, a desktop notification is triggered. However, if the application is shutdown, all counts for previously queried users are lost, i.e the stored results are not persisted.

- The desktop notification feature works across various platforms, but was only tested on Mac OSX 10.13.6. Notifications can easily be extended to include other targets e.g slack, email, prometheus etc. 

 


