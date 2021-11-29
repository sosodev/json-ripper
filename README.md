# json-ripper

It rips JSON from text and dumps all of it to an array.
Also emits the given text, so you can still monitor your output.

```bash
echo "what {\"is\": \"this?\"} wow {\"that's\": \"cool\"}" | json-ripper
what {"is": "this?"} wow {"that's": "cool"}

[{"is":"this?"},{"that's":"cool"}]
```

or with the `-no-echo` flag you just get the ripped JSON without an echo of the input

```bash
echo "what {\"is\": \"this?\"} wow {\"that's\": \"cool\"}" | json-ripper

[{"is":"this?"},{"that's":"cool"}]
```

## Install

with a correctly configured Go installation run

`go install github.com/sosodev/json-ripper@latest`
