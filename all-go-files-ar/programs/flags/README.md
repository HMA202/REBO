# شرح تمرين flags

## فكرة التمرين

البرنامج يستقبل نصًا، ويمكنه تنفيذ عمليتين عليه:

1. إضافة نص آخر في نهايته بواسطة:
   - `--insert=VALUE`
   - `-i=VALUE`
2. ترتيب جميع أحرف النتيجة وفق ASCII بواسطة:
   - `--order`
   - `-o`

كما يطبع صفحة المساعدة إذا:

- لم تُمرر arguments.
- استُخدم `--help`.
- استُخدم `-h`.

---

## أمثلة

```bash
go run . --insert=4321 --order asdad
```

الخطوات:

```text
asdad + 4321 = asdad4321
```

ثم الترتيب:

```text
1234aadds
```

مثال دون ترتيب:

```bash
go run . --insert=4321 asdad
```

الناتج:

```text
asdad4321
```

---

## أهمية تنسيق المساعدة

المطلوب تنسيق دقيق:

```text
--insert
  -i
<TAB><SPACE>This flag inserts the string into the string passed as argument.
--order
  -o
<TAB><SPACE>This flag will behave like a boolean, if it is called it will order the argument.
```

قبل `-i` و`-o` توجد مسافتان.

وقبل جملة الشرح يوجد:

1. Tab واحد `\t`
2. ثم مسافة واحدة

---

## الكود

```go
package main

import (
    "fmt"
    "os"

    "github.com/01-edu/z01"
)

func printHelp() {
    fmt.Print("--insert\n")
    fmt.Print("  -i\n")
    fmt.Print("\t This flag inserts the string into the string passed as argument.\n")
    fmt.Print("--order\n")
    fmt.Print("  -o\n")
    fmt.Print("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func hasPrefix(s, prefix string) bool {
    if len(s) < len(prefix) {
        return false
    }

    return s[:len(prefix)] == prefix
}

func sortASCII(s string) string {
    chars := []rune(s)

    for i := 0; i < len(chars)-1; i++ {
        for j := 0; j < len(chars)-1-i; j++ {
            if chars[j] > chars[j+1] {
                chars[j], chars[j+1] = chars[j+1], chars[j]
            }
        }
    }

    return string(chars)
}

func printString(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }

    z01.PrintRune('\n')
}

func main() {
    if len(os.Args) == 1 {
        printHelp()
        return
    }

    insert := ""
    order := false
    text := ""

    for i := 1; i < len(os.Args); i++ {
        arg := os.Args[i]

        if arg == "--help" || arg == "-h" {
            printHelp()
            return
        }

        if arg == "--order" || arg == "-o" {
            order = true
            continue
        }

        if hasPrefix(arg, "--insert=") {
            insert = arg[len("--insert="):]
            continue
        }

        if hasPrefix(arg, "-i=") {
            insert = arg[len("-i="):]
            continue
        }

        if arg == "--insert" || arg == "-i" {
            if i+1 < len(os.Args) {
                i++
                insert = os.Args[i]
            }
            continue
        }

        text = arg
    }

    result := text + insert

    if order {
        result = sortASCII(result)
    }

    printString(result)
}
```

---

# شرح الدوال

## `printHelp`

تطبع نص المساعدة بالتنسيق المطلوب.

استخدمنا `fmt.Print` لأن `fmt.*` مسموح، ولأن طباعة سطور ثابتة طويلة بـ `z01.PrintRune` ستكون أطول وأقل وضوحًا.

لاحظ الفرق:

```go
"  -i\n"
```

يبدأ بمسافتين عاديتين.

أما:

```go
"\t This flag..."
```

فيبدأ بـ Tab ثم مسافة.

## `hasPrefix`

تتحقق هل النص يبدأ بمقدمة معينة.

مثال:

```go
hasPrefix("--insert=4321", "--insert=")
```

يعيد `true`.

أولًا:

```go
if len(s) < len(prefix)
```

يحمي من محاولة أخذ جزء أطول من النص، مما قد يسبب panic.

ثم:

```go
s[:len(prefix)] == prefix
```

يأخذ بداية `s` بطول `prefix` ويقارنها.

## `sortASCII`

```go
chars := []rune(s)
```

تحول string إلى Slice من runes حتى يمكن تبديل الأحرف.

الـ string في Go غير قابل لتعديل أحرفه مباشرة، فلا يمكن كتابة:

```go
s[0] = 'A'
```

لذلك نستخدم Slice قابلة للتعديل.

بعد Bubble Sort نعيدها إلى string:

```go
return string(chars)
```

## `printString`

تطبع كل rune ثم newline واحدًا.

فصل الطباعة في دالة يجعل `main` أوضح ويمنع تكرار كود الطباعة.

---

# شرح تحليل arguments

## حالة عدم وجود arguments

```go
if len(os.Args) == 1 {
    printHelp()
    return
}
```

عندما يكون الطول 1، يوجد اسم البرنامج فقط.

## المتغيرات

```go
insert := ""
order := false
text := ""
```

- `insert`: النص المطلوب إضافته.
- `order`: هل طُلب الترتيب؟
- `text`: النص الأساسي.

## المرور على arguments

```go
for i := 1; i < len(os.Args); i++ {
```

نبدأ من 1 لتجاهل اسم البرنامج.

## المساعدة لها الأولوية

```go
if arg == "--help" || arg == "-h" {
    printHelp()
    return
}
```

إذا ظهر help، نطبع المساعدة وننهي البرنامج فورًا.

## علم الترتيب

```go
if arg == "--order" || arg == "-o" {
    order = true
    continue
}
```

`order` علم Boolean، لذلك يكفي حفظ `true`.

`continue` ينقل التنفيذ مباشرة إلى argument التالي حتى لا يُعامل العلم كنص أساسي.

## استخراج `--insert=...`

```go
insert = arg[len("--insert="):]
```

إذا كان:

```text
--insert=4321
```

فالجزء الذي يبدأ بعد طول `--insert=` هو:

```text
4321
```

## دعم الشكل المنفصل

الكود يدعم أيضًا:

```bash
--insert 4321
-i 4321
```

عند العثور على العلم، يزيد `i` يدويًا حتى ينتقل للقيمة التالية:

```go
i++
insert = os.Args[i]
```

مع فحص الحدود:

```go
if i+1 < len(os.Args)
```

## النص الأساسي

أي argument غير معروف كعلم يُحفظ في:

```go
text = arg
```

وفق أمثلة التمرين يوجد نص أساسي واحد.

---

# ترتيب العمليات

```go
result := text + insert
```

الإضافة تحدث أولًا.

ثم:

```go
if order {
    result = sortASCII(result)
}
```

الترتيب يحدث على النص الكامل بعد الإضافة.

هذا مهم؛ لأن المثال:

```bash
--insert=4321 --order asdad
```

يتوقع ترتيب `asdad4321` كاملًا، وليس ترتيب `asdad` وحدها.

---

## اختبار المسافات بدقة

```bash
go run . | cat -eT
```

- `-e` يظهر نهاية السطر `$`.
- `-T` يظهر tab على شكل `^I`.

يجب أن ترى بداية سطر الشرح هكذا تقريبًا:

```text
^I This flag...
```

---

## الاختبارات

```bash
cd flags
gofmt -w main.go
go run . --insert=4321 --order asdad
go run . --insert=4321 asdad
go run . asdad
go run . --order 43a21
go run .
go run . -h
go run . --help
```

---

## أخطاء شائعة

1. ترتيب النص قبل الإضافة بدل بعدها.
2. كتابة 8 مسافات يدويًا بدل Tab ثم مسافة.
3. نسيان المسافتين قبل `-i` و`-o`.
4. معاملة `--order` كنص عادي.
5. محاولة تعديل string مباشرة أثناء الترتيب.
6. أخذ slice من نص أقصر من prefix دون فحص الطول.
7. نسيان `continue` بعد معالجة العلم.
