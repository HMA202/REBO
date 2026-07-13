# شرح تمرين nbrconvertalpha

## فكرة التمرين

يستقبل البرنامج أرقامًا كنصوص عبر سطر الأوامر، ويحوّل كل رقم صحيح من `1` إلى `26` إلى الحرف المقابل له في الأبجدية اللاتينية:

```text
1  → a
2  → b
3  → c
...
26 → z
```

أي argument غير صحيح يطبع مكانه مسافة واحدة.

وعند وجود العلم:

```text
--upper
```

كأول argument، تُطبع الحروف كبيرة.

---

## أمثلة

```bash
go run . 8 5 12 12 15
```

الناتج:

```text
hello
```

```bash
go run . 12 5 7 5 14 56 4 1 18 25
```

الناتج:

```text
legen dary
```

الرقم `56` غير صالح، لذلك طُبعت مسافة مكانه.

```bash
go run . --upper 8 5 25
```

الناتج:

```text
HEY
```

---

## الكود

```go
package main

import (
    "os"

    "github.com/01-edu/z01"
)

func getPosition(s string) int {
    if len(s) == 0 {
        return 0
    }

    number := 0

    for _, r := range s {
        if r < '0' || r > '9' {
            return 0
        }

        number = number*10 + int(r-'0')

        if number > 26 {
            return 0
        }
    }

    return number
}

func main() {
    upper := false
    start := 1

    if len(os.Args) > 1 && os.Args[1] == "--upper" {
        upper = true
        start = 2
    }

    if start >= len(os.Args) {
        return
    }

    for i := start; i < len(os.Args); i++ {
        position := getPosition(os.Args[i])

        if position < 1 || position > 26 {
            z01.PrintRune(' ')
        } else if upper {
            z01.PrintRune('A' + rune(position-1))
        } else {
            z01.PrintRune('a' + rune(position-1))
        }
    }

    z01.PrintRune('\n')
}
```

---

# شرح الدالة `getPosition`

## الهدف

`os.Args` تحتوي نصوصًا. حتى عندما يكتب المستخدم:

```bash
8
```

فإن البرنامج يستلمه كـ:

```go
"8"
```

نحتاج تحويل هذا النص يدويًا إلى `int` لأن استخدام `strconv.Atoi` غير مطلوب أو غير مسموح في الحل المقترح.

## التحقق من النص الفارغ

```go
if len(s) == 0 {
    return 0
}
```

النص الفارغ ليس رقمًا صالحًا. نستخدم `0` كقيمة تشير إلى invalid لأن المواضع الصحيحة تبدأ من `1`.

## المرور على الأحرف

```go
for _, r := range s {
```

تمر الحلقة على كل rune في النص.

## التحقق أنه رقم

```go
if r < '0' || r > '9' {
    return 0
}
```

أي حرف خارج المجال من `'0'` إلى `'9'` يجعل argument غير صحيح.

أمثلة غير صحيحة:

```text
h
12a
-5
+3
2.5
```

## تحويل rune رقمي إلى قيمته

```go
r - '0'
```

الأرقام متتابعة في Unicode/ASCII. لذلك:

```text
'5' - '0' = 5
'9' - '0' = 9
```

ثم نحول الناتج إلى `int`:

```go
int(r - '0')
```

## بناء عدد متعدد الخانات

```go
number = number*10 + int(r-'0')
```

مثال تحويل `"12"`:

### البداية

```text
number = 0
```

### قراءة `'1'`

```text
number = 0*10 + 1 = 1
```

### قراءة `'2'`

```text
number = 1*10 + 2 = 12
```

## التوقف عند تجاوز 26

```go
if number > 26 {
    return 0
}
```

لا توجد حاجة لإكمال التحويل؛ لأن أي عدد أكبر من 26 غير صالح.

---

# شرح `main`

## متغير حالة الأحرف الكبيرة

```go
upper := false
```

في الوضع الافتراضي نطبع حروفًا صغيرة.

## موقع أول رقم

```go
start := 1
```

index `0` هو اسم البرنامج، لذلك الأرقام تبدأ عادة من index `1`.

## فحص `--upper`

```go
if len(os.Args) > 1 && os.Args[1] == "--upper" {
    upper = true
    start = 2
}
```

يجب فحص الطول أولًا قبل الوصول إلى `os.Args[1]`، وإلا قد يحدث panic عندما لا توجد arguments.

لأن التعليمات تقول إن العلم سيكون دائمًا أول argument، لا نحتاج البحث عنه في بقية القائمة.

عند وجوده:

- تصبح `upper = true`.
- تبدأ الأرقام من index `2`.

## عدم وجود أرقام

```go
if start >= len(os.Args) {
    return
}
```

إذا لم توجد أرقام بعد اسم البرنامج أو بعد `--upper`، ينتهي البرنامج دون طباعة.

## تحويل كل argument

```go
position := getPosition(os.Args[i])
```

إذا عاد `0`، فالمدخل غير صحيح. وإذا عاد رقمًا من 1 إلى 26، نحوله إلى حرف.

## طباعة مسافة للقيمة غير الصحيحة

```go
if position < 1 || position > 26 {
    z01.PrintRune(' ')
}
```

كل argument ينتج **حرفًا واحدًا** أو **مسافة واحدة**، لذلك يحتفظ الناتج بموقع كل إدخال.

## حساب الحرف الصغير

```go
'a' + rune(position-1)
```

لأن:

```text
position 1 يجب أن يعطي a
```

نطرح 1:

```text
'a' + 0 = 'a'
'a' + 1 = 'b'
'a' + 7 = 'h'
```

## حساب الحرف الكبير

```go
'A' + rune(position-1)
```

نفس الفكرة، لكن البداية من `'A'`.

---

## تتبع `8 5 25`

| argument | position | الحرف الصغير | الحرف الكبير |
|---:|---:|:---:|:---:|
| `8` | 8 | `h` | `H` |
| `5` | 5 | `e` | `E` |
| `25` | 25 | `y` | `Y` |

---

## الاختبار

```bash
cd nbrconvertalpha
gofmt -w main.go
go run . 8 5 12 12 15 | cat -e
go run . 12 5 7 5 14 56 4 1 18 25 | cat -e
go run . 32 86 h | cat -e
go run . --upper 8 5 25 | cat -e
```

---

## أخطاء شائعة

1. استخدام قيمة `position` مباشرة دون طرح 1، فيصبح `1 → b`.
2. قبول `0` أو `27` كقيم صحيحة.
3. تجاهل arguments غير الرقمية بدل طباعة مسافة.
4. طباعة newline بعد كل رقم؛ المطلوب كلمة واحدة ثم newline واحد.
5. فحص `os.Args[1]` دون التأكد من طول القائمة.
