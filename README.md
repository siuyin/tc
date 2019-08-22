# tc manipulates timecodes including 29.97 dropframe timecode (30000/1001)

## Usage
```sh
go get github.com/siuyin/tc

```
See tc_test.go for usage examples.

## Long timecodes
The SMPTE standard covers timecode to "99:59:59:29".

This library has been tested to "120:00:00:00", beyond that
drop frame assumptions cause frame inaccuracies.

