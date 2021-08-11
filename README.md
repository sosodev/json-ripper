# json-ripper

It rips JSON from text and dumps all of it to an array.
Also emits the given text, so you can still monitor your output.

```bash
echo "what {\"is\": \"this?\"} wow {\"that's\": \"cool\"}" | json-ripper
what {"is": "this?"} wow {"that's": "cool"}

[{"is":"this?"},{"that's":"cool"}]
```
