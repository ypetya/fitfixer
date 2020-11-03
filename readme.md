Usage
=====

```
fitfixer <target.fit> <source.fit>
```

Creates or overwrites a file `<target.fit>.withHR.fit`.

Everything is taken from the first file - target.
The HR data is taken from the second file - source.

For every samle in "target" the HeartRate is updated from the "source" file
following the rule to take the first sample having timestamp >= in source.

What it soles
-------------

When you have multiple main devices for example a trainer via bluetooth + android and an ANT+ only strap monitored on GARMIN,
two files can be merged this way.
