Apache Monitor
====

A tool for extract Apache metrics and stored it at Elasticsearch. This is important, because you may have a history of those performance information, and then track if something unusual happened

If you are familiar with Apache Server Status and JK Status, you should know the info above

<dl>
<dt>Current Time: Friday, 21-Nov-2014 14:51:37 BRT</dt>
<dt>Restart Time: Sunday, 16-Nov-2014 04:02:03 BRT</dt>
<dt>Parent Server Generation: 6</dt>
<dt>Server uptime:  5 days 10 hours 49 minutes 34 seconds</dt>
<dt>Total accesses: 8588127 - Total Traffic: 16.3 GB</dt>
<dt>CPU Usage: u52.38 s15.89 cu0 cs0 - .0145% CPU load</dt>
<dt>18.2 requests/sec - 36.3 kB/second - 2038 B/request</dt>
<dt>4 requests currently being processed, 11 idle workers</dt>
<pre>_______W_W_____W_...............................................
................................................................
................................................................
................................................................
</pre>
<p>Scoreboard Key:<br />
"<b><code>_</code></b>" Waiting for Connection, 
"<b><code>S</code></b>" Starting up, 
"<b><code>R</code></b>" Reading Request,<br />
"<b><code>W</code></b>" Sending Reply, 
"<b><code>K</code></b>" Keepalive (read), 
"<b><code>D</code></b>" DNS Lookup,<br />
"<b><code>C</code></b>" Closing connection, 
"<b><code>L</code></b>" Logging, 
"<b><code>G</code></b>" Gracefully finishing,<br /> 
"<b><code>I</code></b>" Idle cleanup of worker, 
"<b><code>.</code></b>" Open slot with no current process</p>
<p />
</dl>


This project aims to extract those informations and store them at Elasticsearch, like this:
```json
{
    "@source": "http://127.0.0.1/server-status/",
    "@tags": [
      "extended"
    ],
    "@timestamp": "2014-11-21T17:56:55.702Z",
    "@fields": {
      "CPULoad": 0.0111,
      "ReqPerSec": 18.2,
      "BytesPerSec": 36.3,
      "BytesPerReq": 2038,
      "BusyWorkers": 14,
      "IdleWorkers": 0,
      "ScoreBoard": "_______W_W_____W_...............................................\n................................................................\n................................................................\n................................................................\n",
      "HostName": "127.0.0.1",
      "Waiting": 14,
      "Starting": 0,
      "Reading": 0,
      "Replying": 3,
      "Keepalive": 0,
      "Dns": 0,
      "Closing": 0,
      "Logging": 0,
      "Grace": 0,
      "Idle": 0,
      "Available": 239
    }
  }
```

But as I said, not only server-status. The JK Status is also supported.

example
=========
To start Apache Monitor, you have to specify a configuration file. Here comes an example

```javascript
{

    "elasticsearch": "http://10.10.2.12:9200",
    "status": [
        {
            "type": "jk",
            "url": "http://hostname1.com/jkstatus",
            "interval": "1m"
        },
        {
            "type": "extended",
            "url": "http://hostname1.com/server-status/",
            "interval": "1m"
        },
        {
            "type": "extended",
            "url": "http://hostname2.com/server-status/",
            "interval": "1m"
        }
    ]

}
```
