# doorbell
:bell:

## set up 
notes for myself when I have to deploy this again
#### env file

```
export IPHONE=iphone number
export TWPHONE=fake twiio number
export SID=twilio sid
export TOKEN=twilio token
export NAME=name of heroku app
export PORT=whatever
```

#### twilio number
set up a twilio number, and set its voice request URL to `<NAME>.herokuapp.com/call`.

#### heroku notes
i skipped all that buildpack business here, so you need to `go build ./...` before pushing to heroku.

#### reference

borrowed heavily from https://www.twilio.com/blog/2014/10/making-and-receiving-phone-calls-with-golang.html