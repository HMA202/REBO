# شرح تمرين revparams

## فكرة التمرين

المطلوب طباعة arguments سطر الأوامر بترتيب معكوس. لا نعكس أحرف كل كلمة، بل نعكس **ترتيب الكلمات نفسها**.

مثال:

```bash
go run . choumi is the best cat
```

الناتج:

```text
cat
best
the
is
choumi
```

---

## الكود

```go
package main

import (
    "os"

    "github.com/01-edu/z01"
)

func main() {
    for i := len(os.Args) - 1; i >= 1; i-- {
        for _, r := range os.Args[i] {
            z01.PrintRune(r)
        }

        z01.PrintRune('\n')
    }
}
```

---

## لماذا نبدأ من `len(os.Args) - 1`؟

في Go، أول index دائمًا هو `0`. لذلك آخر index يساوي:

```go
len(slice) - 1
```

لو كانت `os.Args` تحتوي 6 عناصر، فإن indexes تكون:

```text
0 1 2 3 4 5
```

إذًا آخر عنصر هو:

```go
os.Args[5]
```

وهذا يساوي:

```go
os.Args[len(os.Args)-1]
```

---

## شرح الحلقة العكسية

```go
for i := len(os.Args) - 1; i >= 1; i-- {
```

تتكون من ثلاثة أجزاء:

### البداية

```go
i := len(os.Args) - 1
```

ابدأ من آخر argument.

### الشرط

```go
i >= 1
```

استمر حتى index رقم `1`.

لا نصل إلى `0` لأن:

```go
os.Args[0]
```

هو اسم البرنامج، وليس argument أدخله المستخدم.

### التناقص

```go
i--
```

بعد كل دورة نرجع عنصرًا واحدًا للخلف. وهي مكافئة تقريبًا لـ:

```go
i = i - 1
```

---

## طباعة argument الحالي

```go
for _, r := range os.Args[i] {
    z01.PrintRune(r)
}
```

`os.Args[i]` هو الـ argument الحالي. نمر على حروفه ونطبعها واحدًا واحدًا.

نستخدم `range` لدعم Unicode.

---

## مثال تتبع

عند التشغيل:

```bash
go run . one two three
```

تكون البيانات تقريبًا:

```go
os.Args[1] = "one"
os.Args[2] = "two"
os.Args[3] = "three"
```

الحلقة الخارجية:

| قيمة `i` | العنصر المطبوع |
|---:|---|
| 3 | `three` |
| 2 | `two` |
| 1 | `one` |

الناتج:

```text
three
two
one
```

---

## الفرق بين عكس arguments وعكس الأحرف

الصحيح:

```text
cat
best
the
is
choumi
```

غير الصحيح:

```text
tac
tseb
eht
si
imuohc
```

التمرين يريد تغيير ترتيب الـ arguments فقط، مع بقاء أحرف كل argument كما هي.

---

## حالة عدم وجود arguments

عند:

```bash
go run .
```

غالبًا يكون طول `os.Args` هو `1`؛ لأن القائمة تحتوي اسم البرنامج فقط.

تبدأ `i` بقيمة `0`، والشرط:

```go
0 >= 1
```

خطأ، ولذلك لا يطبع البرنامج شيئًا.

---

## الاختبار

```bash
cd revparams
gofmt -w main.go
go run . choumi is the best cat
```

اختبار مع arguments تحتوي مسافات داخلية:

```bash
go run . "first argument" "second argument"
```

الناتج:

```text
second argument
first argument
```

---

## أخطاء شائعة

1. استخدام شرط `i >= 0`، فيتم طباعة اسم البرنامج أيضًا.
2. استخدام `i++` بدل `i--`، فتحدث حلقة غير صحيحة أو panic.
3. عكس أحرف كل كلمة بدل عكس ترتيب الكلمات.
4. نسيان newline بعد كل argument.
5. الوصول إلى index خارج حدود `os.Args`.
