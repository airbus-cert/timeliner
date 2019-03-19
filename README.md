# ⏳📈 timeliner

A rewrite of mactime, an ancient Perl tool that has (almost) 0 feature.

## Why another tool?

The mactime's capabilities to filter events based on the time are limited to only a date filter. timeliner uses a [real expression engine](https://github.com/Knetic/govaluate) to parse and apply the filtering logic. The following queries can be expressed:

* Show events that happened between 01:00am and 05:00am: `hour >= 1 && hour < 5`
* Show events that happened on Saturday or Sunday: `weekday == 'Sunday || weekday == 'Saturday'`
* Show events that happened between 2018-12-01 and 2018-12-31: `date >= '2018-12-31' && date <= '2018-12-01'`
* Show events that happened between 01:00am and 05:00am on Sundays or Saturday between 2018-12-01 and 2018-12-31: `(hour >= 1 && hour < 5) && (weekday == 'Sunday || weekday == 'Saturday') && (date >= '2018-12-31' && date <= '2018-12-01)`

You get the idea :)

