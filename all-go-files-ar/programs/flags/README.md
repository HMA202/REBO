# شرح تمرين `flags`

# Detailed Explanation of the `flags` Exercise

---

# القسم الأول: الشرح باللغة العربية

## 1. فكرة التمرين

المطلوب كتابة برنامج بلغة Go يستقبل نصًا من سطر الأوامر، ويمكنه تنفيذ عمليتين اختياريتين على هذا النص:

1. **إضافة نص آخر إلى نهاية النص الأساسي** باستخدام أحد العلمين:

```text
--insert=VALUE
-i=VALUE
```

2. **ترتيب جميع أحرف النص حسب ترتيب ASCII** باستخدام أحد العلمين:

```text
--order
-o
```

كما يجب على البرنامج طباعة صفحة المساعدة في حالتين:

* لم يرسل المستخدم أي arguments.
* استخدم المستخدم أحد علمي المساعدة:

```text
--help
-h
```

---

## 2. معنى Flag

الـ **Flag** هو argument خاص يبدأ غالبًا بشرطة واحدة أو شرطتين، ويستخدم لتغيير طريقة عمل البرنامج.

أمثلة:

```text
--order
--insert=4321
-h
-o
```

العلم الطويل يبدأ بشرطتين:

```text
--order
```

أما الشكل المختصر فيبدأ بشرطة واحدة:

```text
-o
```

---

## 3. أمثلة على تشغيل البرنامج

### إضافة نص وترتيب النتيجة

```bash
go run . --insert=4321 --order asdad
```

النص الأساسي:

```text
asdad
```

النص المطلوب إضافته:

```text
4321
```

يتم الدمج أولًا:

```text
asdad4321
```

ثم يتم ترتيب جميع الأحرف:

```text
1234aadds
```

الناتج النهائي:

```text
1234aadds
```

---

### إضافة نص من دون ترتيب

```bash
go run . --insert=4321 asdad
```

يتم دمج:

```text
asdad
```

مع:

```text
4321
```

فتصبح النتيجة:

```text
asdad4321
```

---

### النص الأساسي فقط

```bash
go run . asdad
```

لا يوجد `insert` ولا `order`.

لذلك يطبع البرنامج النص كما هو:

```text
asdad
```

---

### ترتيب النص فقط

```bash
go run . --order 43a21
```

الأحرف قبل الترتيب:

```text
4 3 a 2 1
```

بعد ترتيبها حسب القيم الرقمية:

```text
1 2 3 4 a
```

الناتج:

```text
1234a
```

---

### تشغيل البرنامج دون arguments

```bash
go run .
```

يطبع البرنامج صفحة المساعدة:

```text
--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.
```

---

## 4. أهمية تنسيق صفحة المساعدة

التمرين لا يهتم بالنص فقط، بل يهتم أيضًا بالمسافات والـ Tab.

قبل:

```text
-i
```

يجب أن توجد مسافتان:

```text
__-i
```

حيث `_` هنا تمثل مسافة.

والأمر نفسه مع:

```text
-o
```

أما سطر الشرح فيجب أن يبدأ بـ:

1. Tab واحد.
2. ثم مسافة واحدة.
3. ثم الجملة.

أي:

```text
<TAB><SPACE>This flag...
```

في Go نكتب ذلك هكذا:

```go
"\t This flag..."
```

الرمز:

```go
\t
```

يمثل Tab.

بعده توجد مسافة عادية قبل كلمة `This`.

---

# 5. الكود الكامل

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

# 6. شرح `package main`

```go
package main
```

هذا الملف يمثل برنامجًا قابلًا للتشغيل.

أي برنامج Go مستقل يحتاج عادةً إلى:

```go
package main
```

ويحتاج أيضًا إلى دالة:

```go
func main()
```

دالة `main` هي نقطة بداية تشغيل البرنامج.

عند تنفيذ:

```bash
go run .
```

يبدأ Go من الدالة:

```go
func main()
```

---

# 7. شرح الاستيرادات

```go
import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)
```

البرنامج يستخدم ثلاث مكتبات.

## مكتبة `fmt`

```go
"fmt"
```

تستخدم هنا لطباعة صفحة المساعدة.

نستخدم:

```go
fmt.Print()
```

لأن صفحة المساعدة تحتوي على نصوص ثابتة وطويلة.

مثال:

```go
fmt.Print("--insert\n")
```

---

## مكتبة `os`

```go
"os"
```

تستخدم للوصول إلى arguments المرسلة إلى البرنامج.

نستخدم:

```go
os.Args
```

---

## مكتبة `z01`

```go
"github.com/01-edu/z01"
```

توفر الدالة:

```go
z01.PrintRune()
```

وهي تطبع `rune` واحدًا في كل مرة.

---

## لماذا توجد مسافة فارغة بين `os` و`z01`؟

تنسيق Go يفصل غالبًا بين:

* مكتبات Go القياسية.
* المكتبات الخارجية.

لذلك نكتب:

```go
import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)
```

هذا هو التنسيق المتوقع من أدوات مثل:

```bash
gofmt
gofumpt
```

---

# 8. شرح دالة `printHelp`

```go
func printHelp() {
	fmt.Print("--insert\n")
	fmt.Print("  -i\n")
	fmt.Print("\t This flag inserts the string into the string passed as argument.\n")
	fmt.Print("--order\n")
	fmt.Print("  -o\n")
	fmt.Print("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}
```

وظيفة هذه الدالة هي طباعة صفحة المساعدة كاملة.

فصلناها في دالة خاصة لأن البرنامج قد يحتاج إليها في أكثر من حالة:

* عندما لا توجد arguments.
* عند استخدام `-h`.
* عند استخدام `--help`.

بدل تكرار نفس أسطر الطباعة ثلاث مرات، نكتبها مرة واحدة داخل `printHelp`.

---

## السطر الأول

```go
fmt.Print("--insert\n")
```

يطبع:

```text
--insert
```

ثم:

```go
\n
```

ينقل المؤشر إلى سطر جديد.

---

## السطر الثاني

```go
fmt.Print("  -i\n")
```

بداية النص تحتوي على مسافتين:

```text
"  -i"
```

لذلك يظهر:

```text
  -i
```

---

## سطر الشرح

```go
fmt.Print("\t This flag inserts the string into the string passed as argument.\n")
```

يبدأ بـ:

```go
\t
```

وهو Tab.

ثم توجد مسافة عادية:

```text
" "
```

ثم تبدأ الجملة:

```text
This flag...
```

---

## لماذا استخدمنا `fmt.Print` وليس `fmt.Println`؟

لأننا نريد التحكم الكامل في نهاية كل سطر باستخدام:

```go
\n
```

يمكن استخدام `fmt.Println`، لكن `fmt.Print` يجعل النص المطلوب واضحًا كما هو تمامًا.

---

# 9. شرح دالة `hasPrefix`

```go
func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}

	return s[:len(prefix)] == prefix
}
```

وظيفة الدالة هي معرفة هل النص `s` يبدأ بالنص `prefix`.

مثال:

```go
hasPrefix("--insert=4321", "--insert=")
```

النتيجة:

```go
true
```

لأن النص الأول يبدأ بـ:

```text
--insert=
```

---

## معاملات الدالة

```go
s string
```

هو النص الكامل.

مثال:

```text
--insert=4321
```

أما:

```go
prefix string
```

فهو الجزء الذي نريد البحث عنه في بداية النص.

مثال:

```text
--insert=
```

---

## نوع الإرجاع

```go
bool
```

يعني أن الدالة تعيد إحدى قيمتين:

```go
true
```

أو:

```go
false
```

---

## فحص الطول

```go
if len(s) < len(prefix) {
	return false
}
```

قبل أخذ بداية النص، يجب التأكد أن `s` ليس أقصر من `prefix`.

مثال:

```text
s      = "-i"
prefix = "--insert="
```

لا يمكن أخذ أول 9 أحرف من نص طوله 2.

لو حاولنا تنفيذ:

```go
s[:len(prefix)]
```

على نص أقصر، فقد يحدث:

```text
panic: slice bounds out of range
```

لذلك نتحقق من الطول أولًا.

---

## مقارنة البداية

```go
return s[:len(prefix)] == prefix
```

إذا كان:

```go
s = "--insert=4321"
```

و:

```go
prefix = "--insert="
```

فإن:

```go
len(prefix)
```

يساوي طول النص:

```text
--insert=
```

ثم:

```go
s[:len(prefix)]
```

يعني:

> خذ النص من بدايته حتى طول `prefix`.

النتيجة:

```text
--insert=
```

ثم نقارن:

```go
"--insert=" == "--insert="
```

فتكون النتيجة:

```go
true
```

---

# 10. شرح الـ Slicing

في هذا السطر:

```go
s[:len(prefix)]
```

نستخدم ما يسمى Slice Expression.

الصيغة العامة:

```go
s[start:end]
```

إذا لم نكتب `start`:

```go
s[:end]
```

فإن البداية تكون من الموقع صفر.

مثال:

```go
word := "abcdef"
part := word[:3]
```

قيمة `part`:

```text
abc
```

لأن المواقع المستخدمة هي:

```text
0
1
2
```

أما `3` فهو الحد النهائي غير المشمول.

---

# 11. شرح دالة `sortASCII`

```go
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
```

وظيفة هذه الدالة هي ترتيب أحرف النص من الأصغر إلى الأكبر.

بالنسبة إلى أحرف ASCII يكون الترتيب وفق القيم الرقمية.

مثال:

```text
1 < 2 < 3 < A < B < a < b
```

---

## تحويل النص إلى `[]rune`

```go
chars := []rune(s)
```

الـ `string` في Go غير قابل للتعديل المباشر.

لا يمكن كتابة:

```go
s[0] = 'A'
```

لأن عناصر النص غير قابلة للتغيير بهذه الطريقة.

لذلك نحول النص إلى Slice من `rune`:

```go
[]rune(s)
```

مثال:

```go
s := "43a21"
chars := []rune(s)
```

تصبح `chars` تقريبًا:

```go
[]rune{'4', '3', 'a', '2', '1'}
```

الـ Slice قابلة للتعديل والتبديل.

---

## ما هو `rune`؟

الـ `rune` يمثل حرفًا أو رمز Unicode واحدًا.

أمثلة:

```go
'A'
'4'
'a'
'م'
'☺'
```

تقنيًا، `rune` هو اسم آخر للنوع:

```go
int32
```

لذلك يمكن مقارنة قيم `rune` باستخدام:

```go
>
<
==
```

---

# 12. شرح Bubble Sort

الكود يستخدم خوارزمية تسمى:

```text
Bubble Sort
```

تعتمد على مقارنة كل عنصر بالعنصر الذي بعده.

إذا كان العنصر الحالي أكبر، يتم تبديلهما.

---

## الحلقة الخارجية

```go
for i := 0; i < len(chars)-1; i++ {
```

تحدد عدد مرات المرور على القائمة.

إذا كان عدد الأحرف 5، فإن الحلقة تعمل عدة مرات حتى تتأكد أن جميع الأحرف وصلت إلى مواقعها الصحيحة.

---

## الحلقة الداخلية

```go
for j := 0; j < len(chars)-1-i; j++ {
```

تستخدم لمقارنة العناصر المتجاورة.

تتم مقارنة:

```go
chars[j]
```

مع:

```go
chars[j+1]
```

مثال:

```text
4 مع 3
3 مع a
a مع 2
2 مع 1
```

---

## لماذا نستخدم `-1`؟

في:

```go
chars[j+1]
```

نصل إلى العنصر التالي.

لذلك يجب أن نتوقف قبل آخر عنصر، حتى لا نحاول الوصول إلى موقع خارج حدود Slice.

لو كان طول القائمة 5، فآخر index هو:

```text
4
```

ولا يجوز أن تصبح قيمة `j` هي 4، لأن:

```go
j + 1
```

ستكون 5، وهذا الموقع غير موجود.

---

## لماذا نطرح `i`؟

```go
len(chars)-1-i
```

بعد كل مرور كامل، أكبر عنصر غير مرتب ينتقل إلى نهاية القائمة.

لذلك لا نحتاج إلى فحص آخر العناصر مرة أخرى.

كلما زادت `i`، يقل عدد المقارنات الداخلية.

---

## شرط التبديل

```go
if chars[j] > chars[j+1] {
```

إذا كان الحرف الحالي أكبر من الحرف التالي، فترتيبهما غير صحيح.

مثال:

```go
'4' > '3'
```

صحيح.

لذلك يجب تبديلهما.

---

## عملية التبديل

```go
chars[j], chars[j+1] = chars[j+1], chars[j]
```

هذه طريقة Go لتبديل قيمتين من دون متغير مؤقت.

لو كانت القيم:

```text
chars[j]   = '4'
chars[j+1] = '3'
```

بعد التبديل تصبح:

```text
chars[j]   = '3'
chars[j+1] = '4'
```

بديل أطول كان سيكون:

```go
temp := chars[j]
chars[j] = chars[j+1]
chars[j+1] = temp
```

لكن Go تسمح بالتبديل المباشر.

---

# 13. تتبع Bubble Sort بمثال

لدينا:

```text
43a21
```

القائمة:

```text
[4, 3, a, 2, 1]
```

## المرور الأول

مقارنة `4` و`3`:

```text
4 > 3
```

نبدلهما:

```text
[3, 4, a, 2, 1]
```

مقارنة `4` و`a`:

```text
4 < a
```

لا يتم التبديل:

```text
[3, 4, a, 2, 1]
```

مقارنة `a` و`2`:

```text
a > 2
```

يتم التبديل:

```text
[3, 4, 2, a, 1]
```

مقارنة `a` و`1`:

```text
a > 1
```

يتم التبديل:

```text
[3, 4, 2, 1, a]
```

أصبح أكبر عنصر في النهاية.

بعد بقية المرور تصبح النتيجة:

```text
[1, 2, 3, 4, a]
```

---

# 14. ترتيب ASCII

الأحرف الإنجليزية والأرقام لها قيم رقمية.

أمثلة:

```text
'0' = 48
'1' = 49
'2' = 50

'A' = 65
'B' = 66

'a' = 97
'b' = 98
```

لهذا السبب يأتي الترتيب عادةً هكذا:

```text
الأرقام
ثم الحروف الكبيرة
ثم الحروف الصغيرة
```

مثال:

```text
2
A
a
```

ترتيبها:

```text
2Aa
```

لأن:

```text
50 < 65 < 97
```

### ملاحظة مهمة

الكود يقارن `rune` بالقيمة الرقمية Unicode.

بالنسبة إلى أحرف ASCII تكون النتيجة مطابقة لترتيب ASCII المطلوب في التمرين.

---

## إعادة `[]rune` إلى `string`

```go
return string(chars)
```

بعد انتهاء الترتيب، تكون النتيجة داخل Slice من runes.

لكن توقيع الدالة يطلب إرجاع:

```go
string
```

لذلك نحول:

```go
[]rune
```

إلى:

```go
string
```

---

# 15. شرح دالة `printString`

```go
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}
```

وظيفتها طباعة النص حرفًا حرفًا، ثم طباعة سطر جديد.

---

## المرور على النص

```go
for _, r := range s {
```

تمر الحلقة على كل `rune` داخل النص.

نستخدم `_` لأننا لا نحتاج index.

---

## طباعة الحرف

```go
z01.PrintRune(r)
```

تطبع الحرف الحالي.

---

## طباعة newline

```go
z01.PrintRune('\n')
```

بعد اكتمال طباعة النص، ينزل البرنامج إلى سطر جديد.

هذا مهم عند الاختبار باستخدام:

```bash
cat -e
```

فتظهر نهاية السطر على شكل:

```text
$
```

---

# 16. شرح دالة `main`

```go
func main() {
```

هذه هي نقطة بداية البرنامج.

كل خطوات تحليل arguments ومعالجة النص تحدث داخلها.

---

# 17. ما هي `os.Args`؟

```go
os.Args
```

هي Slice من النصوص:

```go
[]string
```

تحتوي على اسم البرنامج وكل arguments التي أرسلها المستخدم.

مثال:

```bash
go run . --order 43a21
```

تكون القيم تقريبًا:

```go
os.Args[0] // اسم أو مسار البرنامج
os.Args[1] // "--order"
os.Args[2] // "43a21"
```

ملاحظة: عند استخدام `go run` قد يكون `os.Args[0]` مسارًا لملف تنفيذي مؤقت، لكن هذا لا يؤثر على التمرين.

---

# 18. حالة عدم وجود arguments

```go
if len(os.Args) == 1 {
	printHelp()
	return
}
```

حتى لو لم يكتب المستخدم أي argument، تحتوي `os.Args` على اسم البرنامج.

لذلك:

```go
len(os.Args) == 1
```

تعني أن الموجود هو اسم البرنامج فقط.

في هذه الحالة:

```go
printHelp()
```

تطبع صفحة المساعدة.

ثم:

```go
return
```

تنهي دالة `main` فورًا.

---

# 19. شرح المتغيرات الأساسية

```go
insert := ""
order := false
text := ""
```

## المتغير `insert`

```go
insert := ""
```

يخزن النص الذي سيتم إلحاقه بالنص الأساسي.

القيمة الافتراضية نص فارغ:

```text
""
```

إذا لم يستخدم المستخدم `--insert`، فلن تتم إضافة شيء.

---

## المتغير `order`

```go
order := false
```

متغير Boolean يحدد هل يجب ترتيب النتيجة.

القيمة الافتراضية:

```go
false
```

عند العثور على:

```text
--order
```

أو:

```text
-o
```

تتحول إلى:

```go
true
```

---

## المتغير `text`

```go
text := ""
```

يخزن النص الأساسي الذي سيعالجه البرنامج.

مثال:

```text
asdad
```

---

# 20. المرور على arguments

```go
for i := 1; i < len(os.Args); i++ {
```

نبدأ من:

```go
i := 1
```

لأن:

```go
os.Args[0]
```

هو اسم البرنامج.

نستمر ما دام:

```go
i < len(os.Args)
```

وبعد كل دورة:

```go
i++
```

ننتقل إلى argument التالي.

---

## حفظ argument الحالي

```go
arg := os.Args[i]
```

بدل كتابة:

```go
os.Args[i]
```

مرارًا، نحفظ القيمة الحالية في متغير اسمه:

```go
arg
```

---

# 21. معالجة `--help` و`-h`

```go
if arg == "--help" || arg == "-h" {
	printHelp()
	return
}
```

الرمز:

```go
||
```

يعني **أو**.

إذا كانت قيمة `arg` تساوي:

```text
--help
```

أو:

```text
-h
```

يطبع البرنامج المساعدة وينتهي فورًا.

---

## لماذا للمساعدة أولوية؟

لأن المطلوب عند استخدام `--help` أو `-h` هو عرض الخيارات.

لا نحتاج بعد ذلك إلى:

* إضافة النص.
* ترتيب النص.
* طباعة النص الأساسي.

لذلك نستخدم:

```go
return
```

---

# 22. معالجة `--order` و`-o`

```go
if arg == "--order" || arg == "-o" {
	order = true
	continue
}
```

إذا كان argument الحالي هو علم الترتيب، نغير:

```go
order
```

من:

```go
false
```

إلى:

```go
true
```

---

## لماذا `order` متغير Boolean؟

العلم لا يحتاج إلى قيمة إضافية.

وجوده فقط يعني:

```text
رتب النص
```

لذلك يكفي استخدام:

```go
true
```

أو:

```go
false
```

---

## معنى `continue`

```go
continue
```

تنهي الدورة الحالية من الحلقة، وتنتقل مباشرة إلى argument التالي.

هذا يمنع البرنامج من الوصول إلى:

```go
text = arg
```

وتخزين:

```text
--order
```

على أنه النص الأساسي.

---

# 23. معالجة `--insert=VALUE`

```go
if hasPrefix(arg, "--insert=") {
	insert = arg[len("--insert="):]
	continue
}
```

أولًا نتحقق هل يبدأ argument بـ:

```text
--insert=
```

مثال:

```text
--insert=4321
```

إذا كانت النتيجة صحيحة، نستخرج كل ما يأتي بعد علامة `=`.

---

## استخراج القيمة

```go
insert = arg[len("--insert="):]
```

إذا كان:

```go
arg = "--insert=4321"
```

فإن:

```go
len("--insert=")
```

يمثل موقع بداية القيمة.

الجزء:

```go
arg[len("--insert="):]
```

يعني:

> ابدأ من نهاية `--insert=` وخذ بقية النص.

الناتج:

```text
4321
```

---

## مثال آخر

```text
--insert=hello
```

تصبح قيمة `insert`:

```text
hello
```

---

## قيمة فارغة

إذا كتب المستخدم:

```text
--insert=
```

فإن القيمة بعد `=` فارغة.

تصبح:

```go
insert = ""
```

وهذا لا يضيف شيئًا إلى النص الأساسي.

---

# 24. معالجة `-i=VALUE`

```go
if hasPrefix(arg, "-i=") {
	insert = arg[len("-i="):]
	continue
}
```

هذا هو الشكل المختصر من العلم.

مثال:

```bash
go run . -i=4321 asdad
```

تكون قيمة `insert`:

```text
4321
```

---

# 25. دعم الشكل المنفصل

```go
if arg == "--insert" || arg == "-i" {
	if i+1 < len(os.Args) {
		i++
		insert = os.Args[i]
	}
	continue
}
```

بالإضافة إلى الشكل:

```text
--insert=4321
```

يدعم هذا الكود أيضًا:

```text
--insert 4321
```

وكذلك:

```text
-i 4321
```

---

## فحص وجود argument تالٍ

```go
if i+1 < len(os.Args) {
```

قبل قراءة العنصر التالي، نتأكد أنه موجود.

لو كتب المستخدم:

```bash
go run . --insert
```

فلا توجد قيمة بعد العلم.

من دون فحص الحدود قد يحاول البرنامج قراءة موقع غير موجود، مما يسبب panic.

---

## الانتقال إلى القيمة التالية

```go
i++
```

إذا كان `i` يشير إلى:

```text
--insert
```

فإن زيادته تجعله يشير إلى القيمة التالية:

```text
4321
```

---

## حفظ قيمة الإدخال

```go
insert = os.Args[i]
```

يحفظ النص التالي داخل `insert`.

---

## لماذا سيزيد `i` مرتين؟

داخل الشرط نستخدم:

```go
i++
```

حتى نتجاوز قيمة `insert`.

وفي نهاية دورة الحلقة، تنفذ حلقة `for` أيضًا:

```go
i++
```

وهذا مقصود، لأننا نريد تجاوز:

1. علم `--insert`.
2. القيمة التي بعده.

ثم الانتقال إلى argument التالي.

---

# 26. حفظ النص الأساسي

```go
text = arg
```

إذا لم يكن argument:

* مساعدة.
* علم ترتيب.
* علم إدخال.
* قيمة مرتبطة بعلم الإدخال.

فإن البرنامج يعتبره النص الأساسي.

مثال:

```text
asdad
```

فيصبح:

```go
text = "asdad"
```

---

## ماذا لو وُجد أكثر من نص عادي؟

الكود ينفذ:

```go
text = arg
```

كل مرة يجد argument عاديًا.

لذلك آخر argument عادي سيستبدل القيمة السابقة.

لكن حسب صيغة التمرين، يوجد نص أساسي واحد فقط، لذلك هذا السلوك مناسب للاختبارات المتوقعة.

---

# 27. ترتيب الأعلام غير مهم

لأن البرنامج يمر على جميع arguments، يمكنه التعامل مع:

```bash
go run . --insert=4321 --order asdad
```

وكذلك:

```bash
go run . asdad --order --insert=4321
```

في الحالتين سيحفظ:

```text
text   = asdad
insert = 4321
order  = true
```

ثم ينفذ العمليات في الترتيب الصحيح لاحقًا.

---

# 28. دمج النص

```go
result := text + insert
```

بعد إنهاء تحليل arguments، يتم دمج النص الأساسي مع نص الإضافة.

إذا كانت القيم:

```go
text = "asdad"
```

و:

```go
insert = "4321"
```

تصبح:

```go
result = "asdad4321"
```

---

## ماذا يحدث من دون `insert`؟

إذا لم يستخدم المستخدم علم الإدخال:

```go
insert = ""
```

إذًا:

```go
result := text + ""
```

تكون النتيجة مساوية للنص الأساسي.

---

# 29. تنفيذ الترتيب

```go
if order {
	result = sortASCII(result)
}
```

إذا كانت:

```go
order == true
```

يتم إرسال النتيجة كاملة إلى:

```go
sortASCII
```

ثم حفظ النص المرتب مرة أخرى داخل:

```go
result
```

إذا كانت:

```go
order == false
```

يتم تجاوز هذا الشرط ولا يتغير ترتيب النص.

---

# 30. لماذا الإضافة تحدث قبل الترتيب؟

البرنامج ينفذ:

```go
result := text + insert
```

ثم:

```go
result = sortASCII(result)
```

وهذا مهم لأن المطلوب ترتيب النص الكامل بعد الإضافة.

مثال:

```bash
go run . --insert=4321 --order asdad
```

يجب أولًا إنشاء:

```text
asdad4321
```

ثم ترتيب الكل:

```text
1234aadds
```

لو تم ترتيب `asdad` أولًا ثم إضافة `4321`، قد تكون النتيجة:

```text
aadds4321
```

وهذه نتيجة خاطئة.

---

# 31. طباعة النتيجة

```go
printString(result)
```

في النهاية يرسل البرنامج النص النهائي إلى دالة `printString`.

الدالة تطبع جميع الأحرف، ثم تطبع newline واحدًا.

---

# 32. تتبع البرنامج خطوة بخطوة

لنأخذ المثال:

```bash
go run . --insert=4321 --order asdad
```

تكون arguments تقريبًا:

```go
os.Args[0] = اسم البرنامج
os.Args[1] = "--insert=4321"
os.Args[2] = "--order"
os.Args[3] = "asdad"
```

القيم الأولية:

```text
insert = ""
order  = false
text   = ""
```

---

## الدورة الأولى

```go
arg = "--insert=4321"
```

يبدأ بـ:

```text
--insert=
```

لذلك:

```text
insert = "4321"
```

---

## الدورة الثانية

```go
arg = "--order"
```

لذلك:

```text
order = true
```

---

## الدورة الثالثة

```go
arg = "asdad"
```

ليس علمًا معروفًا.

لذلك:

```text
text = "asdad"
```

---

## بعد الحلقة

```go
result := text + insert
```

تصبح:

```text
result = "asdad4321"
```

ولأن:

```text
order = true
```

يتم ترتيب النتيجة:

```text
result = "1234aadds"
```

ثم تطبع:

```text
1234aadds
```

---

# 33. تتبع مثال من دون ترتيب

```bash
go run . --insert=4321 asdad
```

القيم بعد تحليل arguments:

```text
insert = "4321"
order  = false
text   = "asdad"
```

الدمج:

```text
result = "asdad4321"
```

لأن `order` تساوي `false` لا يتم الترتيب.

الناتج:

```text
asdad4321
```

---

# 34. تتبع مثال الترتيب فقط

```bash
go run . --order 43a21
```

القيم:

```text
insert = ""
order  = true
text   = "43a21"
```

الدمج:

```text
result = "43a21"
```

الترتيب:

```text
result = "1234a"
```

الناتج:

```text
1234a
```

---

# 35. اختبار صفحة المساعدة بدقة

استخدم:

```bash
go run . | cat -eT
```

الخيار:

```text
-e
```

يظهر نهاية كل سطر على شكل:

```text
$
```

والخيار:

```text
-T
```

يظهر Tab على شكل:

```text
^I
```

يجب أن ترى شيئًا قريبًا من:

```text
--insert$
  -i$
^I This flag inserts the string into the string passed as argument.$
--order$
  -o$
^I This flag will behave like a boolean, if it is called it will order the argument.$
```

ظهور:

```text
^I
```

يعني أن بداية السطر تحتوي على Tab حقيقي.

---

# 36. الاختبارات المقترحة

## تنسيق الملف

```bash
gofmt -w main.go
```

أو:

```bash
gofumpt -w main.go
```

ثم تأكد:

```bash
gofumpt -d main.go
```

إذا لم يظهر شيء، فالتنسيق صحيح.

---

## اختبار الإضافة والترتيب

```bash
go run . --insert=4321 --order asdad
```

المتوقع:

```text
1234aadds
```

---

## اختبار الإضافة فقط

```bash
go run . --insert=4321 asdad
```

المتوقع:

```text
asdad4321
```

---

## اختبار النص فقط

```bash
go run . asdad
```

المتوقع:

```text
asdad
```

---

## اختبار الترتيب فقط

```bash
go run . --order 43a21
```

المتوقع:

```text
1234a
```

---

## اختبار الاختصارات

```bash
go run . -i=4321 -o asdad
```

المتوقع:

```text
1234aadds
```

---

## اختبار المساعدة

```bash
go run .
go run . -h
go run . --help
```

يجب أن تطبع الأوامر الثلاثة صفحة المساعدة نفسها.

---

# 37. الأخطاء الشائعة

## الخطأ الأول: ترتيب النص قبل الإضافة

خطأ:

```go
result := sortASCII(text)
result += insert
```

قد ينتج:

```text
aadds4321
```

الصحيح:

```go
result := text + insert

if order {
	result = sortASCII(result)
}
```

---

## الخطأ الثاني: استخدام مسافات بدل Tab

خطأ:

```go
fmt.Print("         This flag...")
```

قد يبدو متشابهًا في بعض الشاشات، لكنه ليس Tab حقيقيًا.

الصحيح:

```go
fmt.Print("\t This flag...")
```

---

## الخطأ الثالث: نسيان المسافتين قبل الاختصار

خطأ:

```go
fmt.Print("-i\n")
```

الصحيح:

```go
fmt.Print("  -i\n")
```

---

## الخطأ الرابع: تخزين `--order` كنص أساسي

لو نسينا:

```go
continue
```

قد يصل التنفيذ إلى:

```go
text = arg
```

فتصبح قيمة `text`:

```text
--order
```

---

## الخطأ الخامس: تعديل `string` مباشرة

هذا غير مسموح:

```go
s[0] = 'A'
```

الصحيح هو التحويل إلى:

```go
chars := []rune(s)
```

---

## الخطأ السادس: أخذ Slice دون فحص الطول

هذا قد يسبب panic:

```go
return s[:len(prefix)] == prefix
```

إذا كان `s` أقصر من `prefix`.

لذلك يجب أولًا:

```go
if len(s) < len(prefix) {
	return false
}
```

---

## الخطأ السابع: نسيان newline

يجب أن تنتهي النتيجة بـ:

```go
z01.PrintRune('\n')
```

وإلا قد يفشل الاختبار الذي يقارن المخرجات بدقة.

---

## الخطأ الثامن: نسيان newline في نهاية ملف Go نفسه

إضافة:

```go
z01.PrintRune('\n')
```

لا تعني أن ملف `main.go` ينتهي بسطر جديد.

بعد آخر:

```go
}
```

يجب أن يوجد newline في الملف.

يمكن إصلاحه باستخدام:

```bash
gofmt -w main.go
```

وتفعيله دائمًا في VS Code:

```json
"files.insertFinalNewline": true
```

---

# 38. التعقيد

## تحليل arguments

إذا كان عدد arguments هو `m`، فالمرور عليها يحتاج تقريبًا:

```text
O(m)
```

مع احتساب أطوال النصوص التي يتم فحص بداياتها.

---

## ترتيب الأحرف

خوارزمية Bubble Sort تحتوي على حلقتين متداخلتين.

إذا كان طول النص النهائي `n`، فإن التعقيد الزمني هو:

```text
O(n²)
```

---

## الذاكرة الإضافية

يتم إنشاء:

```go
[]rune(s)
```

بحجم يعتمد على عدد الأحرف.

لذلك التعقيد المكاني:

```text
O(n)
```

---

# 39. الخلاصة العربية

البرنامج يعمل وفق الخطوات التالية:

1. يقرأ arguments من `os.Args`.
2. يطبع المساعدة إذا لم توجد arguments.
3. يطبع المساعدة عند استخدام `-h` أو `--help`.
4. يحفظ قيمة `--insert` أو `-i`.
5. يفعّل الترتيب عند استخدام `--order` أو `-o`.
6. يحفظ النص الأساسي.
7. يضيف نص `insert` إلى نهاية النص الأساسي.
8. يرتب النتيجة كاملة إذا كان `order` مفعّلًا.
9. يطبع النتيجة باستخدام `z01.PrintRune`.
10. يضيف newline في نهاية المخرجات.

---

---

# Part Two: English Explanation

## 1. Exercise Overview

The program receives a string through command-line arguments and can perform two optional operations:

1. Append another string using:

```text
--insert=VALUE
-i=VALUE
```

2. Sort all characters using ASCII order with:

```text
--order
-o
```

The program must print its help page when:

* No arguments are provided.
* The user passes `--help`.
* The user passes `-h`.

---

## 2. What Is a Flag?

A flag is a special command-line argument that changes how a program behaves.

Examples:

```text
--order
--insert=4321
-h
-o
```

A long flag usually starts with two hyphens:

```text
--order
```

A short flag starts with one hyphen:

```text
-o
```

---

## 3. Usage Examples

### Insert and Sort

```bash
go run . --insert=4321 --order asdad
```

First, the strings are joined:

```text
asdad + 4321 = asdad4321
```

Then all characters are sorted:

```text
1234aadds
```

Final output:

```text
1234aadds
```

---

### Insert Without Sorting

```bash
go run . --insert=4321 asdad
```

Output:

```text
asdad4321
```

---

### Original String Only

```bash
go run . asdad
```

Output:

```text
asdad
```

---

### Sort Only

```bash
go run . --order 43a21
```

Output:

```text
1234a
```

---

### No Arguments

```bash
go run .
```

The program prints the help page.

---

# 4. Exact Help Formatting

The short flags must have exactly two spaces before them:

```text
  -i
  -o
```

Each explanation line must start with:

1. One tab.
2. One normal space.
3. The sentence.

In Go:

```go
"\t This flag..."
```

The `\t` escape sequence represents one tab character.

---

# 5. Complete Code

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

# 6. Package Declaration

```go
package main
```

This declares an executable Go program.

An executable Go program normally requires:

```go
package main
```

and:

```go
func main()
```

The `main` function is the program's entry point.

---

# 7. Imports

```go
import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)
```

## `fmt`

The `fmt` package is used to print the fixed help text.

The program uses:

```go
fmt.Print()
```

---

## `os`

The `os` package provides access to command-line arguments through:

```go
os.Args
```

---

## `z01`

The external `z01` package provides:

```go
z01.PrintRune()
```

It prints one rune at a time.

---

## Import Grouping

The blank line separates standard-library packages from an external package.

This layout matches Go formatting tools such as:

```bash
gofmt
gofumpt
```

---

# 8. The `printHelp` Function

```go
func printHelp() {
	fmt.Print("--insert\n")
	fmt.Print("  -i\n")
	fmt.Print("\t This flag inserts the string into the string passed as argument.\n")
	fmt.Print("--order\n")
	fmt.Print("  -o\n")
	fmt.Print("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}
```

This function prints the complete help page.

It is separated into its own function because the same output is required in several situations:

* No arguments.
* `-h`.
* `--help`.

This avoids repeating the same print statements.

---

## Newline Escape

```go
\n
```

moves the cursor to the next line.

---

## Tab Escape

```go
\t
```

inserts one tab character.

The following string:

```go
"\t This flag..."
```

starts with:

* One tab.
* One space.
* The sentence.

---

# 9. The `hasPrefix` Function

```go
func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}

	return s[:len(prefix)] == prefix
}
```

This function checks whether `s` begins with `prefix`.

Example:

```go
hasPrefix("--insert=4321", "--insert=")
```

returns:

```go
true
```

---

## Length Protection

```go
if len(s) < len(prefix) {
	return false
}
```

The function must ensure that `s` is long enough before slicing it.

Without this check, an expression such as:

```go
s[:len(prefix)]
```

could cause:

```text
panic: slice bounds out of range
```

---

## Prefix Comparison

```go
return s[:len(prefix)] == prefix
```

For:

```text
s      = --insert=4321
prefix = --insert=
```

the slice:

```go
s[:len(prefix)]
```

produces:

```text
--insert=
```

The comparison is therefore true.

---

# 10. String Slicing

The general form is:

```go
s[start:end]
```

The `start` position is included.

The `end` position is not included.

For example:

```go
word := "abcdef"
part := word[:3]
```

`part` contains:

```text
abc
```

---

# 11. The `sortASCII` Function

```go
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
```

This function sorts the characters from the smallest numeric value to the largest.

For ASCII characters, that corresponds to ASCII ordering.

---

## Converting the String to Runes

```go
chars := []rune(s)
```

Go strings are immutable.

This is not allowed:

```go
s[0] = 'A'
```

The string is converted to a rune slice because a slice can be modified.

Example:

```go
s := "43a21"
chars := []rune(s)
```

The resulting slice is conceptually:

```go
[]rune{'4', '3', 'a', '2', '1'}
```

---

## What Is a Rune?

A rune represents one Unicode code point.

Examples:

```go
'A'
'4'
'a'
'م'
'☺'
```

Technically:

```go
rune
```

is an alias for:

```go
int32
```

This means runes can be compared numerically.

---

# 12. Bubble Sort

The function uses the Bubble Sort algorithm.

Bubble Sort repeatedly compares adjacent elements.

When the left element is greater than the right element, their positions are swapped.

---

## Outer Loop

```go
for i := 0; i < len(chars)-1; i++ {
```

The outer loop controls how many passes are made through the slice.

---

## Inner Loop

```go
for j := 0; j < len(chars)-1-i; j++ {
```

The inner loop compares adjacent elements:

```go
chars[j]
```

and:

```go
chars[j+1]
```

---

## Why `-1` Is Required

Because the code accesses:

```go
chars[j+1]
```

`j` must stop before the last index.

Otherwise, the code would try to access an element beyond the end of the slice.

---

## Why `-i` Is Used

After each complete pass, the largest remaining element reaches the end.

That final section is already sorted, so it does not need to be checked again.

Subtracting `i` reduces unnecessary comparisons.

---

## Swap Condition

```go
if chars[j] > chars[j+1] {
```

If the current character has a greater numeric value than the next character, they are in the wrong order.

---

## Swapping Values

```go
chars[j], chars[j+1] = chars[j+1], chars[j]
```

Go supports direct multiple assignment.

This swaps two values without a temporary variable.

---

# 13. Bubble Sort Trace

Input:

```text
43a21
```

Initial slice:

```text
[4, 3, a, 2, 1]
```

Compare `4` and `3`:

```text
4 > 3
```

Swap:

```text
[3, 4, a, 2, 1]
```

Compare `4` and `a`:

```text
4 < a
```

No swap.

Compare `a` and `2`:

```text
a > 2
```

Swap:

```text
[3, 4, 2, a, 1]
```

Compare `a` and `1`:

```text
a > 1
```

Swap:

```text
[3, 4, 2, 1, a]
```

After all passes:

```text
[1, 2, 3, 4, a]
```

Final string:

```text
1234a
```

---

# 14. ASCII Values

ASCII characters have numeric values.

Examples:

```text
'0' = 48
'1' = 49
'2' = 50

'A' = 65
'B' = 66

'a' = 97
'b' = 98
```

Therefore, ASCII sorting generally places:

1. Digits before letters.
2. Uppercase letters before lowercase letters.

Example:

```text
2Aa
```

because:

```text
50 < 65 < 97
```

### Important Note

The code compares Unicode rune values.

For ASCII input, Unicode values and ASCII values are the same, so the result matches the exercise requirement.

---

## Converting Back to String

```go
return string(chars)
```

The function converts the sorted rune slice back into a Go string.

---

# 15. The `printString` Function

```go
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}
```

This function prints every rune in the string and then prints one newline.

---

## Range Loop

```go
for _, r := range s {
```

The loop reads one Unicode rune at a time.

The `_` ignores the byte index because it is not needed.

---

## Printing

```go
z01.PrintRune(r)
```

prints the current rune.

---

## Final Newline

```go
z01.PrintRune('\n')
```

ensures the output ends correctly.

---

# 16. The `main` Function

```go
func main() {
```

The `main` function is where command-line arguments are parsed and the requested operations are performed.

---

# 17. Understanding `os.Args`

```go
os.Args
```

is a slice of strings:

```go
[]string
```

It contains the executable name followed by the user arguments.

For:

```bash
go run . --order 43a21
```

the values are conceptually:

```go
os.Args[0] // executable path
os.Args[1] // "--order"
os.Args[2] // "43a21"
```

When using `go run`, `os.Args[0]` may be a temporary executable path. This does not affect the exercise.

---

# 18. No-Argument Case

```go
if len(os.Args) == 1 {
	printHelp()
	return
}
```

Even when the user provides no arguments, `os.Args` still contains the executable name.

Therefore:

```go
len(os.Args) == 1
```

means that no user arguments were provided.

The program prints help and exits.

---

# 19. Initial State Variables

```go
insert := ""
order := false
text := ""
```

## `insert`

Stores the string that will be appended.

The default value is empty.

---

## `order`

Stores whether sorting was requested.

The default is:

```go
false
```

---

## `text`

Stores the main input string.

---

# 20. Argument Loop

```go
for i := 1; i < len(os.Args); i++ {
```

The loop begins at index `1` to skip the executable name.

---

## Current Argument

```go
arg := os.Args[i]
```

The current argument is stored in `arg` to make the following conditions easier to read.

---

# 21. Help Flag

```go
if arg == "--help" || arg == "-h" {
	printHelp()
	return
}
```

If either help form is found, the program prints help and stops immediately.

The `||` operator means logical OR.

---

# 22. Order Flag

```go
if arg == "--order" || arg == "-o" {
	order = true
	continue
}
```

The presence of the flag sets:

```go
order = true
```

The flag does not require a separate value.

---

## Why `continue` Is Important

`continue` skips the remaining statements in the current loop iteration.

Without it, the program could later execute:

```go
text = arg
```

and mistakenly use `--order` as the main text.

---

# 23. Long Insert Form

```go
if hasPrefix(arg, "--insert=") {
	insert = arg[len("--insert="):]
	continue
}
```

For:

```text
--insert=4321
```

the prefix is:

```text
--insert=
```

The remaining part is:

```text
4321
```

That value is stored in `insert`.

---

# 24. Short Insert Form

```go
if hasPrefix(arg, "-i=") {
	insert = arg[len("-i="):]
	continue
}
```

For:

```text
-i=4321
```

the stored value is:

```text
4321
```

---

# 25. Separate Insert Value

```go
if arg == "--insert" || arg == "-i" {
	if i+1 < len(os.Args) {
		i++
		insert = os.Args[i]
	}
	continue
}
```

This additionally supports:

```bash
--insert 4321
```

and:

```bash
-i 4321
```

The bounds check:

```go
i+1 < len(os.Args)
```

ensures that a next argument exists.

The manual increment:

```go
i++
```

moves from the flag to its value.

---

# 26. Main Text

```go
text = arg
```

Any argument that is not recognized as a flag is treated as the main string.

The exercise is expected to provide one main text argument.

If multiple normal strings are provided, the last one replaces the previous value because of this assignment.

---

# 27. Joining the Strings

```go
result := text + insert
```

The insertion string is appended to the end of the main text.

Example:

```text
text   = asdad
insert = 4321
```

Result:

```text
asdad4321
```

---

# 28. Conditional Sorting

```go
if order {
	result = sortASCII(result)
}
```

Sorting occurs only when `order` is true.

The entire joined string is sorted.

This order of operations is essential.

---

# 29. Why Insertion Happens Before Sorting

Correct flow:

```go
result := text + insert
result = sortASCII(result)
```

For:

```text
asdad + 4321
```

the sorted result is:

```text
1234aadds
```

If only `asdad` were sorted before appending `4321`, the result would be wrong.

---

# 30. Final Output

```go
printString(result)
```

The program prints the processed string followed by one newline.

---

# 31. Full Execution Trace

Command:

```bash
go run . --insert=4321 --order asdad
```

Arguments:

```go
os.Args[1] = "--insert=4321"
os.Args[2] = "--order"
os.Args[3] = "asdad"
```

Initial values:

```text
insert = ""
order  = false
text   = ""
```

After processing `--insert=4321`:

```text
insert = "4321"
```

After processing `--order`:

```text
order = true
```

After processing `asdad`:

```text
text = "asdad"
```

Join:

```text
result = "asdad4321"
```

Sort:

```text
result = "1234aadds"
```

Output:

```text
1234aadds
```

---

# 32. Checking Exact Whitespace

Use:

```bash
go run . | cat -eT
```

`-e` shows line endings with `$`.

`-T` shows tabs as `^I`.

Expected form:

```text
--insert$
  -i$
^I This flag inserts the string into the string passed as argument.$
--order$
  -o$
^I This flag will behave like a boolean, if it is called it will order the argument.$
```

---

# 33. Recommended Tests

```bash
gofmt -w main.go
gofumpt -d main.go
```

Then:

```bash
go run . --insert=4321 --order asdad
go run . --insert=4321 asdad
go run . asdad
go run . --order 43a21
go run . -i=4321 -o asdad
go run .
go run . -h
go run . --help
```

---

# 34. Common Mistakes

1. Sorting before appending.
2. Using spaces instead of a real tab.
3. Forgetting the two spaces before `-i` and `-o`.
4. Treating `--order` as normal text.
5. Trying to modify a Go string directly.
6. Slicing before verifying the string length.
7. Forgetting `continue` after processing a flag.
8. Forgetting the output newline.
9. Forgetting the final newline at the end of the source file.
10. Starting the argument loop at index `0` instead of `1`.

---

# 35. Complexity

Let `n` be the number of characters in the final joined string.

## Time Complexity

Argument parsing is approximately linear in the number of arguments.

Bubble Sort requires:

```text
O(n²)
```

time.

---

## Space Complexity

The conversion:

```go
[]rune(s)
```

creates a rune slice.

Therefore, the additional space usage is:

```text
O(n)
```

---

# 36. English Summary

The program:

1. Reads command-line arguments from `os.Args`.
2. Prints help when no user arguments are present.
3. Prints help for `-h` or `--help`.
4. extracts the value of `--insert` or `-i`.
5. Enables sorting for `--order` or `-o`.
6. Stores the main input string.
7. Appends the insertion value.
8. Sorts the complete result when requested.
9. Prints the result using `z01.PrintRune`.
10. Ends the output with one newline.
