# شرح تمرين rotatevowels

## فكرة التمرين

المطلوب جمع جميع حروف العلة الموجودة في جميع arguments، ثم عكس ترتيب حروف العلة فقط وإعادتها إلى مواقع حروف العلة الأصلية.

حروف العلة في هذا التمرين:

```text
a e i o u
A E I O U
```

الحرف `y` ليس حرف علة.

الحروف الساكنة والمسافات تبقى في أماكنها كما هي.

---

## مثال أساسي

```bash
go run . "Hello World"
```

حروف العلة بالترتيب:

```text
e o o
```

بعد عكسها:

```text
o o e
```

نعيدها إلى مواقع حروف العلة:

```text
Hello World
Hollo Werld
```

---

## الكود

```go
package main

import (
    "os"

    "github.com/01-edu/z01"
)

func isVowel(r rune) bool {
    return r == 'a' ||
        r == 'e' ||
        r == 'i' ||
        r == 'o' ||
        r == 'u' ||
        r == 'A' ||
        r == 'E' ||
        r == 'I' ||
        r == 'O' ||
        r == 'U'
}

func main() {
    vowels := []rune{}

    for _, argument := range os.Args[1:] {
        for _, r := range argument {
            if isVowel(r) {
                vowels = append(vowels, r)
            }
        }
    }

    vowelIndex := len(vowels) - 1

    for argumentIndex, argument := range os.Args[1:] {
        for _, r := range argument {
            if isVowel(r) {
                z01.PrintRune(vowels[vowelIndex])
                vowelIndex--
            } else {
                z01.PrintRune(r)
            }
        }

        if argumentIndex < len(os.Args[1:])-1 {
            z01.PrintRune(' ')
        }
    }

    z01.PrintRune('\n')
}
```

---

# لماذا نحتاج مرورين على النص؟

لا يمكن معرفة أول حرف علة بديل قبل معرفة آخر حرف علة في جميع arguments.

لذلك نستخدم مرحلتين:

1. المرور الأول: جمع جميع حروف العلة.
2. المرور الثاني: إعادة طباعة النص، لكن أخذ حروف العلة من نهاية القائمة إلى بدايتها.

---

# شرح `isVowel`

```go
func isVowel(r rune) bool
```

تستقبل حرفًا واحدًا وتعيد `true` إذا كان حرف علة، وإلا تعيد `false`.

الرمز:

```go
||
```

يعني OR المنطقية.

الدالة تقرأ كالتالي:

> أعد true إذا كان `r` يساوي `a` أو `e` أو `i` أو `o` أو `u`، أو أحد نظائرها الكبيرة.

أمثلة:

```go
isVowel('a') // true
isVowel('E') // true
isVowel('y') // false
isVowel('b') // false
```

---

# المرحلة الأولى: جمع حروف العلة

## إنشاء Slice فارغة

```go
vowels := []rune{}
```

ستحتوي حروف العلة بنفس ترتيب ظهورها.

## المرور على كل arguments

```go
for _, argument := range os.Args[1:] {
```

نستبعد اسم البرنامج باستخدام `[1:]`.

## المرور على كل حرف

```go
for _, r := range argument {
```

يقرأ كل حرف كـ rune، فيعمل مع Unicode بصورة صحيحة.

## إضافة حرف العلة

```go
if isVowel(r) {
    vowels = append(vowels, r)
}
```

`append` تضيف عنصرًا إلى نهاية Slice.

مثال:

```text
Hello World
```

تتطور Slice هكذا:

```go
[]rune{'e'}
[]rune{'e', 'o'}
[]rune{'e', 'o', 'o'}
```

---

# تحديد نقطة البداية العكسية

```go
vowelIndex := len(vowels) - 1
```

آخر index لأي Slice يساوي طولها ناقص 1.

إذا كانت:

```go
vowels = []rune{'e', 'o', 'o'}
```

فإن:

```text
len(vowels) = 3
vowelIndex = 2
```

أي نبدأ من آخر `o`.

إذا لم توجد حروف علة، تصبح القيمة `-1`. هذا آمن هنا لأن الكود لا يستخدم `vowels[vowelIndex]` إلا عندما يعثر على حرف علة. وإذا لم توجد حروف علة فلن يدخل ذلك الفرع.

---

# المرحلة الثانية: الطباعة والاستبدال

```go
for argumentIndex, argument := range os.Args[1:] {
```

نحتاج `argumentIndex` حتى نعرف هل نطبع مسافة بين arguments.

## عند حرف علة

```go
if isVowel(r) {
    z01.PrintRune(vowels[vowelIndex])
    vowelIndex--
}
```

لا نطبع حرف العلة الأصلي. نطبع بدلًا منه الحرف الموجود في نهاية `vowels`، ثم نرجع index خطوة إلى الخلف.

## عند حرف عادي

```go
else {
    z01.PrintRune(r)
}
```

يطبع كما هو دون تغيير.

---

# وضع المسافات بين arguments

عند تشغيل:

```bash
go run . "HEllO World" "problem solved"
```

Shell يرسل argumentين منفصلين دون الاحتفاظ بالمسافة التي كانت بينهما في الأمر. لذلك يجب أن يعيد البرنامج طباعة مسافة بينهما.

```go
if argumentIndex < len(os.Args[1:])-1 {
    z01.PrintRune(' ')
}
```

نطبع المسافة فقط إذا لم نكن عند آخر argument، حتى لا توجد مسافة إضافية في نهاية السطر.

---

# تتبع مثال كامل

المدخلات:

```text
"aEi" "Ou"
```

حروف العلة المجمعة:

```text
a E i O u
```

الترتيب العكسي:

```text
u O i E a
```

نضعها في المواقع الأصلية:

```text
uOi Ea
```

لاحظ أن حالة الحرف الكبيرة أو الصغيرة تنتقل مع الحرف نفسه؛ لا نحافظ على حالة الموقع الأصلي.

---

# حالة لا توجد فيها حروف علة

```bash
go run . "str" "shh" "psst"
```

`vowels` تبقى فارغة، وكل الأحرف تدخل فرع `else`، لذلك الناتج مطابق للمدخلات:

```text
str shh psst
```

---

# حالة لا توجد arguments

عند:

```bash
go run .
```

لا تعمل الحلقتان، ثم يصل البرنامج إلى:

```go
z01.PrintRune('\n')
```

فيطبع newline فقط كما تطلب التعليمات.

---

## الاختبارات

```bash
cd rotatevowels
gofmt -w main.go
go run . "Hello World" | cat -e
go run . "HEllO World" "problem solved"
go run . "str" "shh" "psst"
go run . "happy thoughts" "good luck"
go run . "aEi" "Ou"
go run . | cat -e
```

---

## أخطاء شائعة

1. عكس النص كله بدل حروف العلة فقط.
2. معالجة كل argument منفصلًا؛ المطلوب عكس حروف العلة عبر جميع arguments كمجموعة واحدة.
3. اعتبار `y` حرف علة.
4. نسيان الحروف الكبيرة `A E I O U`.
5. نسيان المسافة بين arguments.
6. إضافة مسافة بعد آخر argument.
7. محاولة الوصول إلى `vowels[-1]` دون التأكد من وجود حرف علة؛ الكود الحالي يتجنب ذلك طبيعيًا.
