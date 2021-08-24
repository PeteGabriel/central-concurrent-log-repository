# Central and Concurrent Log Repository

Server application that receives info from clients (specified number) in concurrency and centralizes this info in a log file. Reports can be extracted from this application.

- [x] Accept at least 5 concurrent clients. Number can be specified as an env variable.
- [ ] Data saving actions and reports are written/read to/from a file named numbers.log.
- [x] Numbers and must be composed of exactly nine decimal digits are to be accepted by this server application. An incorrect input will terminate the connection with the client.
- [ ] No duplicate numbers may be written to the log file.
- [x] The command **terminate** is also accepted and terminates all the client connections plus this server application.
- [ ] Every 10 seconds, the Application must print a report to standard output:

The report should include the following points:
1. The difference since the last report of the count of new unique numbers that have
   been received.
2. The difference since the last report of the count of new duplicate numbers that
   have been received.
3. The total number of unique numbers received for this run of the Application.
4. Example: `Received 50 unique numbers, 2 duplicates. Unique total:
   234567`

more features to be added soon



## How to:

```
 docker build --tag docker-central-log . && docker run -p 4001:4001 docker-central-log
 
 ls //after any writing operation, 'numbers.log' file should appear
```

Clients can open connections using `telnet` for example.
