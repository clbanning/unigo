<h2>unigo - encode unicode characters in slice as string</h2>
Utility functions for encoding ASCII unicode code-points (UTF-8/UTF-16) to runes.

<h4>Usage</h4>

<pre>
data := []byte(`\u044d\u0442\u043e \u0442\u0435\u0441\u0442 \u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u0435`)
fmt.Println(unigo.EncodeToString(data) // prints: это тест сообщение
</pre>

or:
<pre>
fmt.Println(unigo.EncodeToString([]byte(`{"key":"wasn\u0027t"}`))) // prints: {"key":"wasn't"}
</pre>

<h4>Documentation</h4>

At the [usual place](https://godoc.org/github.com/clbanning/unigo).

<h4>Motivation</h4>

There was a discussion on [go-nuts](https://groups.google.com/forum/#!topic/golang-nuts/KO1yubIbKpU) that seemed to call out for such a utility.
