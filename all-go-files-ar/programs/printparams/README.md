# شرح تمرين printparams

## فكرة التمرين

المطلوب كتابة برنامج يطبع كل argument تم تمريره في سطر الأوامر، بحيث يظهر كل argument في سطر مستقل.

مثال:

```bash
go run . choumi is the best cat
```

الناتج:

```text
choumi
is
the
best
cat
```

---

## الملف المطلوب

```text
printparams/main.go
```

## الكود

```go
package main

import (
    "os"

    "github.com/01-edu/z01"
)

func main() {
    for _, argument := range os.Args[1:] {
        for _, r := range argument {
            z01.PrintRune(r)
        }

        z01.PrintRune('\n')
    }
}
```

---

## فهم `os.Args`

`os.Args` هي Slice من النصوص تحتوي على اسم البرنامج ثم الـ arguments.

عند تنفيذ:

```bash
go run . choumi is the best cat
```

تكون الفكرة كالتالي:

```go
os.Args[0] // مسار الملف التنفيذي المؤقت الذي أنشأه go run
os.Args[1] // "choumi"
os.Args[2] // "is"
os.Args[3] // "the"
os.Args[4] // "best"
os.Args[5] // "cat"
```

المطلوب لا يشمل اسم البرنامج، لذلك نستخدم:

```go
os.Args[1:]
```

هذه تسمى عملية slicing، ومعناها: خذ العناصر من index رقم `1` حتى النهاية.

---

## الحلقة الخارجية

```go
for _, argument := range os.Args[1:] {
```

تمر على كل argument.

- استخدمنا `_` بدل index لأننا لا نحتاج رقم العنصر.
- `argument` يحتوي النص الحالي.

في المثال السابق ستكون قيمته بالتتابع:

```text
choumi
is
the
best
cat
```

لكن الحلقة الخارجية لا تطبع النص دفعة واحدة؛ لأن الدالة الوحيدة المسموحة للطباعة هي `z01.PrintRune`، وهي تطبع حرفًا واحدًا.

---

## الحلقة الداخلية

```go
for _, r := range argument {
    z01.PrintRune(r)
}
```

تمر على أحرف الـ argument الحالي.

لو كان:

```go
argument = "cat"
```

ستكون قيم `r`:

```text
'c'
'a'
't'
```

ثم تتم طباعة كل حرف.

استخدام `range` مهم لأنه يتعامل مع Unicode. لذلك يستطيع البرنامج طباعة:

```bash
go run . مرحبا ☺
```

دون تفكيك الحروف إلى بايتات مشوهة.

---

## النزول إلى سطر جديد

```go
z01.PrintRune('\n')
```

بعد انتهاء طباعة argument واحد، يطبع البرنامج newline. لذلك يظهر الـ argument التالي في سطر جديد.

لاحظ أن newline موجود **داخل الحلقة الخارجية**؛ لأننا نحتاج newline بعد كل argument.

---

## ماذا يحدث دون arguments؟

عند تنفيذ:

```bash
go run .
```

تكون:

```go
os.Args[1:]
```

Slice فارغة. لذلك لا تدخل الحلقة ولا يطبع البرنامج شيئًا، وهو السلوك الطبيعي لهذا التمرين.

---

## تتبع كامل

للتشغيل:

```bash
go run . hello world
```

1. الحلقة الخارجية تأخذ `hello`.
2. الحلقة الداخلية تطبع `h`, ثم `e`, ثم `l`, ثم `l`, ثم `o`.
3. يطبع `\n`.
4. الحلقة الخارجية تأخذ `world`.
5. الحلقة الداخلية تطبع أحرفها.
6. يطبع `\n`.

الناتج:

```text
hello
world
```

---

## الاختبار

```bash
cd printparams
gofmt -w main.go
go run . choumi is the best cat
```

اختبار Unicode:

```bash
go run . مرحبا ☺
```

فحص newline:

```bash
go run . hello | cat -e
```

الناتج المتوقع:

```text
hello$
```

---

## أخطاء شائعة

1. البدء من `os.Args[0]`، فيتم طباعة اسم البرنامج.
2. طباعة مسافة بين arguments بدل newline.
3. وضع `z01.PrintRune('\n')` خارج الحلقة، فتظهر كل arguments متصلة.
4. استخدام فهرسة البايتات مع نصوص Unicode.
5. نسيان newline بعد آخر `}` في الملف نفسه.
