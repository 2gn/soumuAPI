package zyloutils

import "zylo/reiwa"

```
F -> contest number
```

func raiseError(err error) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		reiwa.DisplayToast("Error in raiseError")
	}
	reiwa.DisplayToast(fmt.Sprintf("In %s, at %d: %s", file, line, err.Error()))
}