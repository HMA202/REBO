# `alphacount.go` — عدّ الحروف الإنجليزية

# `alphacount.go` — Counting English Alphabetic Characters

---

# القسم العربي

## 1. فكرة التمرين

يحتوي الملف `alphacount.go` على دالة اسمها:

```go
AlphaCount
```

وظيفة الدالة هي حساب عدد **الحروف الإنجليزية فقط** الموجودة داخل نص.

الحروف التي يتم حسابها هي:

```text
a إلى z
A إلى Z
```

أما العناصر التالية فلا يتم حسابها:

* الأرقام مثل `1` و`5` و`9`.
* المسافات.
* الرموز مثل `!` و`@` و`#`.
* علامات الترقيم مثل `.` و`,` و`?`.
* الحروف العربية.
* أي أحرف أخرى لا تقع بين `a-z` أو `A-Z`.

---

## 2. الكود الكامل

```go
package piscine

func AlphaCount(s string) int {
	count := 0

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}

	return count
}
```

---

## 3. توقيع الدالة

```go
func AlphaCount(s string) int
```

يمكن تقسيم هذا السطر إلى الأجزاء التالية:

```go
func
```

تعني أننا نعرّف دالة جديدة.

```go
AlphaCount
```

هو اسم الدالة.

```go
s string
```

يعني أن الدالة تستقبل قيمة واحدة:

* اسمها `s`.
* نوعها `string`.
* تمثل النص الذي نريد فحصه.

```go
int
```

يعني أن الدالة تعيد رقمًا صحيحًا من نوع `int`.

مثال:

```go
result := AlphaCount("Hello 123")
```

بعد التنفيذ تصبح قيمة `result`:

```go
5
```

لأن كلمة `Hello` تحتوي على خمسة حروف إنجليزية.

---

## 4. شرح الكود سطرًا سطرًا

### السطر الأول

```go
package piscine
```

يحدد هذا السطر الحزمة التي ينتمي إليها الملف.

الحزمة اسمها:

```go
piscine
```

هذا يعني أن الملف يحتوي على دالة يمكن استدعاؤها من ملفات أخرى داخل الحزمة نفسها أو من برنامج يستورد هذه الحزمة.

هذا الملف ليس برنامجًا مستقلًا، لأنه لا يحتوي على:

```go
package main
```

ولا يحتوي على:

```go
func main()
```

---

### تعريف الدالة

```go
func AlphaCount(s string) int {
```

هذا السطر يعرّف دالة اسمها `AlphaCount`.

تستقبل الدالة نصًا:

```go
s string
```

وتعيد عددًا صحيحًا:

```go
int
```

مثال على قيمة `s`:

```text
"Hello 2026!"
```

---

### إنشاء العداد

```go
count := 0
```

أنشأنا متغيرًا اسمه:

```go
count
```

وظيفته حفظ عدد الحروف الإنجليزية التي وجدناها.

القيمة الابتدائية هي:

```go
0
```

لأننا في بداية الدالة لم نفحص أي حرف بعد.

الرمز:

```go
:=
```

يستخدم لتعريف متغير جديد داخل الدالة.

Go تستنتج نوع المتغير من القيمة الموجودة في الجهة اليمنى.

بما أن القيمة:

```go
0
```

عدد صحيح، يصبح نوع `count` هو:

```go
int
```

هذا السطر يعادل تقريبًا:

```go
var count int = 0
```

---

## 5. حلقة المرور على النص

```go
for _, r := range s {
```

هذه حلقة `for` تستخدم `range` للمرور على النص حرفًا حرفًا.

إذا كانت قيمة `s`:

```text
"Go!"
```

فإن الحلقة تمر على:

```text
G
o
!
```

في كل دورة، يتم وضع الحرف الحالي داخل المتغير:

```go
r
```

نوع `r` هو:

```go
rune
```

---

## 6. ما هو `rune`؟

في Go، النوع `rune` يمثل رمز Unicode واحدًا.

أمثلة:

```go
'A'
'b'
'7'
'!'
'م'
```

كل واحد من هذه العناصر يمكن تمثيله على شكل `rune`.

تقنيًا، `rune` هو اسم آخر للنوع:

```go
int32
```

لكن كلمة `rune` توضح أن القيمة تمثل حرفًا أو رمزًا، وليس مجرد رقم عادي.

في الكود:

```go
for _, r := range s
```

يحصل المتغير `r` على كل حرف من النص على حدة.

---

## 7. لماذا استخدمنا `_`؟

في حلقة `range` على النص، يمكننا الحصول على قيمتين:

```go
for index, r := range s
```

القيمة الأولى هي موقع الحرف داخل النص:

```go
index
```

والقيمة الثانية هي الحرف نفسه:

```go
r
```

لكن في هذه الدالة لا نحتاج موقع الحرف، بل نحتاج الحرف فقط.

لذلك نكتب:

```go
_
```

بدلًا من اسم المتغير.

```go
for _, r := range s
```

الشرطة السفلية `_` تسمى **Blank Identifier**.

معناها:

> تجاهل هذه القيمة لأنني لا أحتاجها.

---

## 8. شرط التحقق من الحرف

```go
if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
```

هذا الشرط يتحقق هل الحرف الحالي حرف إنجليزي صغير أو كبير.

الشرط يحتوي على جزأين.

### الجزء الأول

```go
r >= 'a' && r <= 'z'
```

يتحقق هل الحرف يقع بين:

```text
a و z
```

مثل:

```text
a
b
c
m
x
z
```

---

### الجزء الثاني

```go
r >= 'A' && r <= 'Z'
```

يتحقق هل الحرف يقع بين:

```text
A و Z
```

مثل:

```text
A
B
G
M
X
Z
```

---

## 9. معنى `&&`

الرمز:

```go
&&
```

يعني **و**.

في هذا الشرط:

```go
r >= 'a' && r <= 'z'
```

يجب أن يكون الشرطان صحيحين معًا:

1. قيمة الحرف أكبر من أو تساوي `'a'`.
2. قيمة الحرف أصغر من أو تساوي `'z'`.

إذا كان الحرف:

```go
'm'
```

فإن:

```go
'm' >= 'a'
```

صحيح.

و:

```go
'm' <= 'z'
```

صحيح.

إذًا النتيجة النهائية:

```go
true
```

---

## 10. معنى `||`

الرمز:

```go
||
```

يعني **أو**.

الشرط الكامل يقول:

```go
الحرف بين a و z
```

أو:

```go
الحرف بين A و Z
```

يكفي أن يكون أحد الجزأين صحيحًا حتى يتم اعتبار الحرف حرفًا إنجليزيًا.

---

## 11. لماذا تعمل المقارنة بين الحروف؟

الحروف في Go لها قيم رقمية في Unicode.

قيم الحروف الإنجليزية متتابعة.

على سبيل المثال:

```text
'A' = 65
'B' = 66
'C' = 67
...
'Z' = 90
```

والحروف الصغيرة:

```text
'a' = 97
'b' = 98
'c' = 99
...
'z' = 122
```

لذلك يمكننا التحقق من أن الحرف يقع داخل مجال معين باستخدام:

```go
r >= 'a' && r <= 'z'
```

هذا لا يحوّل الحرف يدويًا إلى رقم، لأن Go تقوم بالمقارنة باستخدام القيمة الرقمية للحرف تلقائيًا.

---

## 12. زيادة العداد

```go
count++
```

هذا السطر يزيد قيمة `count` بمقدار واحد.

يعادل تقريبًا:

```go
count = count + 1
```

مثال:

إذا كانت قيمة `count`:

```go
3
```

بعد:

```go
count++
```

تصبح:

```go
4
```

لا يتم تنفيذ هذا السطر إلا عندما يكون الحرف الحالي حرفًا إنجليزيًا.

---

## 13. إرجاع النتيجة

```go
return count
```

بعد انتهاء الحلقة وفحص جميع الحروف، تعيد الدالة القيمة الموجودة في `count`.

مثال:

```go
AlphaCount("Go 2026!")
```

الحروف الإنجليزية هي:

```text
G
o
```

إذًا تكون قيمة `count`:

```go
2
```

وتعيد الدالة:

```go
return 2
```

---

# 14. تتبع التنفيذ خطوة بخطوة

لنفترض أننا استدعينا الدالة هكذا:

```go
AlphaCount("Go 2!")
```

في البداية:

```text
count = 0
```

ثم تبدأ الحلقة.

| الدورة | الحرف `r` | هل هو حرف إنجليزي؟ | قيمة `count` |
| -----: | :-------: | :----------------: | -----------: |
|      1 |    `G`    |  نعم، بين `A` و`Z` |            1 |
|      2 |    `o`    |  نعم، بين `a` و`z` |            2 |
|      3 |   مسافة   |         لا         |            2 |
|      4 |    `2`    |         لا         |            2 |
|      5 |    `!`    |         لا         |            2 |

بعد انتهاء الحلقة:

```go
return count
```

تعيد:

```text
2
```

---

# 15. مثال تفصيلي آخر

```go
AlphaCount("Hello 2026!")
```

الحروف التي تمر عليها الحلقة:

```text
H
e
l
l
o
 
2
0
2
6
!
```

الحروف التي تحقق الشرط:

```text
H
e
l
l
o
```

عددها:

```text
5
```

الناتج:

```go
5
```

---

# 16. أمثلة إضافية

```go
AlphaCount("Go-Lang")
```

الحروف الإنجليزية:

```text
G o L a n g
```

الناتج:

```go
6
```

---

```go
AlphaCount("123 !")
```

لا توجد حروف إنجليزية.

الناتج:

```go
0
```

---

```go
AlphaCount("")
```

النص فارغ، لذلك الحلقة لا تعمل.

الناتج:

```go
0
```

---

```go
AlphaCount("ABCxyz")
```

جميع العناصر حروف إنجليزية.

الناتج:

```go
6
```

---

```go
AlphaCount("مرحبا")
```

الحروف عربية وليست ضمن المجال:

```text
A-Z
a-z
```

الناتج:

```go
0
```

---

```go
AlphaCount("Hi مرحبا")
```

الحروف الإنجليزية هي:

```text
H
i
```

الناتج:

```go
2
```

---

# 17. لماذا لا تُحسب الحروف العربية؟

الشرط يتحقق فقط من المجالين:

```go
'a' إلى 'z'
```

و:

```go
'A' إلى 'Z'
```

الحروف العربية لها قيم Unicode مختلفة، لذلك لا تقع داخل هذين المجالين.

رغم أن `range` و`rune` يستطيعان قراءة الحروف العربية بشكل صحيح، فإن الشرط يتجاهلها عمدًا.

---

# 18. الفرق بين `byte` و`rune`

كان من الممكن المرور على النص باستخدام المواقع:

```go
for i := 0; i < len(s); i++ {
```

لكن:

```go
s[i]
```

يعيد `byte` وليس `rune`.

هذا قد يسبب مشكلة مع الأحرف التي تحتاج إلى أكثر من byte، مثل:

```text
م
ح
😊
```

استخدام:

```go
for _, r := range s
```

أفضل لأنه يقرأ كل رمز Unicode كاملًا.

مع ذلك، الدالة لا تعد إلا الحروف الإنجليزية بسبب الشرط المستخدم.

---

# 19. الحالات الخاصة

## النص الفارغ

```go
AlphaCount("")
```

الحلقة لا تنفذ أي دورة، وتبقى قيمة `count` تساوي صفرًا.

الناتج:

```go
0
```

---

## نص يحتوي على مسافات فقط

```go
AlphaCount("   ")
```

المسافة ليست حرفًا إنجليزيًا.

الناتج:

```go
0
```

---

## نص يحتوي على أرقام فقط

```go
AlphaCount("2026")
```

الأرقام لا تحقق الشرط.

الناتج:

```go
0
```

---

## نص يحتوي على رموز

```go
AlphaCount("@#$!")
```

الرموز لا تحقق الشرط.

الناتج:

```go
0
```

---

## نص مختلط

```go
AlphaCount("Go123! مرحبا")
```

الحروف الإنجليزية هي:

```text
G
o
```

الناتج:

```go
2
```

---

# 20. الأخطاء الشائعة

## الخطأ الأول: استخدام `||` بدل `&&` داخل المجال

هذا الشرط خاطئ:

```go
if r >= 'a' || r <= 'z'
```

لأن معظم القيم ستحقق أحد الجزأين، مما يؤدي إلى حساب عناصر غير صحيحة.

الصحيح:

```go
if r >= 'a' && r <= 'z'
```

لأن الحرف يجب أن يكون أكبر من أو يساوي `a`، وفي الوقت نفسه أصغر من أو يساوي `z`.

---

## الخطأ الثاني: فحص الحروف الصغيرة فقط

هذا الشرط:

```go
if r >= 'a' && r <= 'z'
```

لن يحسب الحروف الكبيرة مثل:

```text
A
B
C
```

لذلك يجب إضافة:

```go
r >= 'A' && r <= 'Z'
```

---

## الخطأ الثالث: زيادة العداد خارج الشرط

هذا الكود خاطئ:

```go
for _, r := range s {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
	}
	count++
}
```

لأن `count++` خارج `if`، وبالتالي سيتم حساب كل عنصر في النص.

الصحيح:

```go
if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
	count++
}
```

---

## الخطأ الرابع: نسيان `return`

إذا لم نكتب:

```go
return count
```

فلن تعيد الدالة النتيجة، وسيظهر خطأ لأن توقيع الدالة يطلب إرجاع `int`.

---

## الخطأ الخامس: استخدام `len(s)`

هذا الكود:

```go
return len(s)
```

لا يحسب الحروف الإنجليزية.

بل يحسب عدد الـ bytes الموجودة في النص.

مثلًا:

```go
len("Hi 12!")
```

يعطي عدد جميع العناصر، وليس عدد الحروف الإنجليزية فقط.

---

# 21. التعقيد الزمني

## الزمن

```text
O(n)
```

حيث `n` هو عدد الرموز التي تمر عليها الحلقة.

السبب هو أن الدالة تمر على النص مرة واحدة فقط.

---

## الذاكرة الإضافية

```text
O(1)
```

لأن الدالة تستخدم متغيرًا واحدًا رئيسيًا:

```go
count
```

ولا تنشئ slice أو قائمة يزداد حجمها مع طول النص.

---

# 22. طريقة اختبار الدالة

يمكن إنشاء ملف اختبار باسم:

```text
alphacount_test.go
```

ووضعه في الحزمة نفسها:

```go
package piscine

import "testing"

func TestAlphaCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello 2026!", 5},
		{"Go-Lang", 6},
		{"123 !", 0},
		{"", 0},
		{"ABCxyz", 6},
		{"مرحبا", 0},
		{"Hi مرحبا", 2},
	}

	for _, test := range tests {
		result := AlphaCount(test.input)

		if result != test.expected {
			t.Errorf(
				"AlphaCount(%q) = %d; expected %d",
				test.input,
				result,
				test.expected,
			)
		}
	}
}
```

ثم شغّل:

```bash
go test
```

إذا كانت جميع الحالات صحيحة، ستظهر نتيجة مشابهة:

```text
PASS
```

---

# 23. الخلاصة العربية

الدالة `AlphaCount`:

1. تستقبل نصًا من نوع `string`.
2. تنشئ عدادًا يبدأ من صفر.
3. تمر على النص باستخدام `range`.
4. تتعامل مع كل حرف على شكل `rune`.
5. تتحقق هل الحرف بين `a-z` أو `A-Z`.
6. تزيد العداد عند العثور على حرف إنجليزي.
7. تعيد العدد النهائي.

---

---

# English Section

## 1. Exercise Idea

The `alphacount.go` file contains a function named:

```go
AlphaCount
```

The function counts only English alphabetic characters inside a string.

The characters that are counted are:

```text
a through z
A through Z
```

The following values are ignored:

* Digits such as `1`, `5`, and `9`.
* Spaces.
* Symbols such as `!`, `@`, and `#`.
* Punctuation marks.
* Arabic characters.
* Any character outside the English `A-Z` and `a-z` ranges.

---

## 2. Complete Code

```go
package piscine

func AlphaCount(s string) int {
	count := 0

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}

	return count
}
```

---

## 3. Function Signature

```go
func AlphaCount(s string) int
```

The signature can be divided into several parts.

```go
func
```

Declares a function.

```go
AlphaCount
```

Is the function name.

```go
s string
```

Means that the function receives one parameter:

* The parameter is named `s`.
* Its type is `string`.
* It contains the text that will be examined.

```go
int
```

Means that the function returns an integer.

Example:

```go
result := AlphaCount("Hello 123")
```

After execution:

```go
result == 5
```

because `Hello` contains five English letters.

---

## 4. Line-by-Line Explanation

### Package Declaration

```go
package piscine
```

This line declares that the file belongs to the `piscine` package.

The file is not a standalone program because it does not contain:

```go
package main
```

or:

```go
func main()
```

Instead, it defines a reusable function that can be tested or called from another Go file.

---

### Function Definition

```go
func AlphaCount(s string) int {
```

This defines a function named `AlphaCount`.

The function receives a string named `s` and returns an integer.

Example input:

```text
"Hello 2026!"
```

---

### Counter Initialization

```go
count := 0
```

A variable named `count` is created.

Its purpose is to store the number of English letters found in the string.

The initial value is zero because no characters have been examined yet.

The operator:

```go
:=
```

declares a new variable and lets Go infer its type.

Since `0` is an integer value, `count` becomes an `int`.

This is approximately equivalent to:

```go
var count int = 0
```

---

## 5. Iterating Over the String

```go
for _, r := range s {
```

This loop uses `range` to iterate over the string one Unicode character at a time.

For example, if:

```go
s == "Go!"
```

the loop processes:

```text
G
o
!
```

During each iteration, the current character is stored in:

```go
r
```

The type of `r` is:

```go
rune
```

---

## 6. What Is a `rune`?

In Go, a `rune` represents one Unicode code point.

Examples:

```go
'A'
'b'
'7'
'!'
'م'
```

Technically, `rune` is an alias for:

```go
int32
```

However, the name `rune` communicates that the value represents a character rather than a general integer.

In this loop:

```go
for _, r := range s
```

`r` receives each Unicode character from the string.

---

## 7. Why Is `_` Used?

A `range` loop over a string can provide two values:

```go
for index, r := range s
```

The first value is the byte position:

```go
index
```

The second value is the character:

```go
r
```

The function does not need the position of the character.

Therefore, the position is ignored using:

```go
_
```

The underscore is called the **blank identifier**.

It means:

> Receive this value, but intentionally ignore it.

---

## 8. Checking Whether the Character Is an English Letter

```go
if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
```

This condition checks whether the current character is either:

* A lowercase English letter.
* An uppercase English letter.

The condition has two main parts.

### Lowercase Check

```go
r >= 'a' && r <= 'z'
```

This checks whether `r` is between `a` and `z`.

Examples that satisfy this condition:

```text
a
b
m
x
z
```

---

### Uppercase Check

```go
r >= 'A' && r <= 'Z'
```

This checks whether `r` is between `A` and `Z`.

Examples that satisfy this condition:

```text
A
B
M
X
Z
```

---

## 9. Meaning of `&&`

The operator:

```go
&&
```

means logical **AND**.

For this expression:

```go
r >= 'a' && r <= 'z'
```

both conditions must be true:

1. `r` must be greater than or equal to `'a'`.
2. `r` must be less than or equal to `'z'`.

For example, with:

```go
r == 'm'
```

both comparisons are true, so the complete expression evaluates to:

```go
true
```

---

## 10. Meaning of `||`

The operator:

```go
||
```

means logical **OR**.

The full condition means:

```text
The character is lowercase
OR
The character is uppercase
```

Only one of the two groups needs to be true.

---

## 11. Why Character Comparison Works

Characters in Go have numeric Unicode values.

English letters occupy continuous numeric ranges.

Uppercase letters:

```text
'A' = 65
'B' = 66
'C' = 67
...
'Z' = 90
```

Lowercase letters:

```text
'a' = 97
'b' = 98
'c' = 99
...
'z' = 122
```

Because the values are continuous, the program can check whether a character belongs to a range.

For example:

```go
r >= 'a' && r <= 'z'
```

checks whether the numeric value of `r` is inside the lowercase English range.

---

## 12. Incrementing the Counter

```go
count++
```

This increases `count` by one.

It is equivalent to:

```go
count = count + 1
```

Example:

```text
count = 3
```

After:

```go
count++
```

the value becomes:

```text
count = 4
```

This statement only runs when the current character is an English letter.

---

## 13. Returning the Result

```go
return count
```

After all characters have been processed, the function returns the final value stored in `count`.

Example:

```go
AlphaCount("Go 2026!")
```

The English characters are:

```text
G
o
```

Therefore:

```go
count == 2
```

and the function returns:

```go
2
```

---

# 14. Step-by-Step Execution Trace

Consider this call:

```go
AlphaCount("Go 2!")
```

Initially:

```text
count = 0
```

The loop processes each character.

| Iteration | Character `r` | English letter? | `count` |
| --------: | :-----------: | :-------------: | ------: |
|         1 |      `G`      |  Yes, uppercase |       1 |
|         2 |      `o`      |  Yes, lowercase |       2 |
|         3 |     Space     |        No       |       2 |
|         4 |      `2`      |        No       |       2 |
|         5 |      `!`      |        No       |       2 |

After the loop, the function returns:

```go
2
```

---

# 15. More Examples

```go
AlphaCount("Hello 2026!")
```

Result:

```go
5
```

---

```go
AlphaCount("Go-Lang")
```

Counted letters:

```text
G o L a n g
```

Result:

```go
6
```

---

```go
AlphaCount("123 !")
```

Result:

```go
0
```

---

```go
AlphaCount("")
```

Result:

```go
0
```

---

```go
AlphaCount("ABCxyz")
```

Result:

```go
6
```

---

```go
AlphaCount("مرحبا")
```

Result:

```go
0
```

Arabic letters are valid Unicode runes, but they are outside the English letter ranges checked by the condition.

---

```go
AlphaCount("Hi مرحبا")
```

Counted letters:

```text
H
i
```

Result:

```go
2
```

---

# 16. Why Arabic Letters Are Not Counted

The condition only checks these ranges:

```go
'A' through 'Z'
```

and:

```go
'a' through 'z'
```

Arabic characters have different Unicode values.

The `range` loop reads Arabic characters correctly as runes, but the condition intentionally ignores them.

---

# 17. Difference Between `byte` and `rune`

A string could also be traversed using indexes:

```go
for i := 0; i < len(s); i++ {
```

However:

```go
s[i]
```

returns a `byte`, not a complete Unicode character.

This can cause problems with characters that use multiple UTF-8 bytes, such as:

```text
م
ح
😊
```

Using:

```go
for _, r := range s
```

is safer because it processes complete Unicode code points.

The function still counts only English letters because of its condition.

---

# 18. Edge Cases

## Empty String

```go
AlphaCount("")
```

The loop performs zero iterations.

Result:

```go
0
```

---

## Spaces Only

```go
AlphaCount("   ")
```

Spaces do not satisfy the condition.

Result:

```go
0
```

---

## Digits Only

```go
AlphaCount("2026")
```

Digits are not English alphabetic characters.

Result:

```go
0
```

---

## Symbols Only

```go
AlphaCount("@#$!")
```

Symbols do not satisfy the condition.

Result:

```go
0
```

---

## Mixed String

```go
AlphaCount("Go123! مرحبا")
```

The counted characters are:

```text
G
o
```

Result:

```go
2
```

---

# 19. Common Mistakes

## Mistake 1: Using `||` Inside a Range Check

Incorrect:

```go
if r >= 'a' || r <= 'z'
```

This expression is true for most characters because only one side needs to be true.

Correct:

```go
if r >= 'a' && r <= 'z'
```

The character must satisfy both boundaries.

---

## Mistake 2: Checking Only Lowercase Letters

This condition:

```go
if r >= 'a' && r <= 'z'
```

does not count uppercase letters.

The uppercase range must also be included:

```go
r >= 'A' && r <= 'Z'
```

---

## Mistake 3: Incrementing Outside the `if`

Incorrect:

```go
for _, r := range s {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
	}

	count++
}
```

This counts every character.

Correct:

```go
if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
	count++
}
```

---

## Mistake 4: Forgetting `return`

The function promises to return an `int`.

Without:

```go
return count
```

the program will not compile.

---

## Mistake 5: Returning `len(s)`

Incorrect:

```go
return len(s)
```

`len(s)` returns the number of bytes in the string, not the number of English letters.

---

# 20. Complexity

## Time Complexity

```text
O(n)
```

The function scans the input string once.

`n` represents the number of characters processed.

---

## Extra Space Complexity

```text
O(1)
```

The function uses only a fixed number of variables, mainly:

```go
count
```

It does not create an additional slice or collection whose size grows with the input.

---

# 21. Testing the Function

Create a test file named:

```text
alphacount_test.go
```

Use:

```go
package piscine

import "testing"

func TestAlphaCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello 2026!", 5},
		{"Go-Lang", 6},
		{"123 !", 0},
		{"", 0},
		{"ABCxyz", 6},
		{"مرحبا", 0},
		{"Hi مرحبا", 2},
	}

	for _, test := range tests {
		result := AlphaCount(test.input)

		if result != test.expected {
			t.Errorf(
				"AlphaCount(%q) = %d; expected %d",
				test.input,
				result,
				test.expected,
			)
		}
	}
}
```

Run the tests with:

```bash
go test
```

When all cases pass, Go displays output similar to:

```text
PASS
```

---

# 22. English Summary

The `AlphaCount` function:

1. Receives a string.
2. Creates a counter initialized to zero.
3. Iterates over the string using `range`.
4. Reads each character as a `rune`.
5. Checks whether the character is between `a-z` or `A-Z`.
6. Increments the counter when an English letter is found.
7. Returns the final count.

The solution runs in linear time and uses constant extra memory.
