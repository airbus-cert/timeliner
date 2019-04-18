# ⏳📈 timeliner

A rewrite of mactime, an ancient Perl tool that has (almost) 0 feature.

## Why another tool?

The mactime's capabilities to filter events based on the time are limited to only a date filter. timeliner uses a [real expression engine](https://github.com/Knetic/govaluate) to parse and apply the filtering logic. The following queries can be expressed using a BPF syntax:

* Show events that happened between 01:00am and 05:00am: `hour >= 1 && hour < 5`
* Show events that happened on Saturday or Sunday: `weekday == 'Sunday || weekday == 'Saturday'`
* Show events that happened between 2018-12-01 and 2018-12-31: `date >= '2018-12-31' && date <= '2018-12-01'`
* Show events that happened between 01:00am and 05:00am on Sundays or Saturday between 2018-12-01 and 2018-12-31: `(hour >= 1 && hour < 5) && (weekday == 'Sunday || weekday == 'Saturday') && (date >= '2018-12-31' && date <= '2018-12-01)`

You get the idea :)

The project is still ⍺ and 👼and is missing a few must-have features, but the killer feature is its expression engine which is ready.

## How to use it?

```
$ timeliner -h
Usage of timeliner:
	timeliner [options] MFT.txt

  -color
    	Enable color output
  -filter string
    	Event filter, like "hour > 14"
  -strict
    	Only show the entries maching the date restrictions

$ timeliner -filter 'hour >= 1 && hour < 5' MFT.txt
2006-10-10 02:15:35: \.\Users\xxx\AppData\Local\Temp\eo117895978tm
           02:16:07: \.\Users\xxx\AppData\Local\Temp\eo117895980tm
2007-05-24 03:24:43: \.\Users\xxx\AppData\Local\Temp\eo130872105tm
           03:24:43: \.\Users\xxx\AppData\Local\Temp\eo113046312tm
           03:24:43: \.\Users\xxx\AppData\Local\Temp\eo112784182tm
           03:24:43: \.\Users\xxx\AppData\Local\Temp\eo112063273tm
```

There is a `-strict` flag to limit the output to only the matching event. For example, for one file, its modification time could be in 2015 while the creation in 2013, if we filter events after 2015:
* without the strict mode, both events (in 2013 and 2015) would show up.
* with strict mode, only the 2015 event would be kept:

```
$ timeliner MFT.txt
2013-04-10 08:42:37: \.\Windows\System32\winevt\Logs\Setup.evtx
2015-02-16 15:58:27: \.\Windows\System32\winevt\Logs\Setup.evtx

$ timeliner -filter 'date > "2015-01-01"' MFT.txt
2013-04-10 08:42:37: \.\Windows\System32\winevt\Logs\Setup.evtx
2015-02-16 15:58:27: \.\Windows\System32\winevt\Logs\Setup.evtx

$ timeliner -strict -filter 'date > "2015-01-01"' MFT.txt
2015-02-16 15:58:27: \.\Windows\System32\winevt\Logs\Setup.evtx
```

