# شرح `eightqueens.go` — حل مسألة الملكات الثماني بالـ backtracking

> الدالة الأساسية: `EightQueens`

## ما وظيفة الكود؟

تطبع جميع ترتيبات وضع 8 ملكات على رقعة 8×8 بحيث لا تهدد أي ملكة أخرى في الصف أو القطر. كل سطر يمثل صف الملكة في كل عمود.

## الكود الأصلي

```go
package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var queens [8]int

	placeQueen(&queens, 0)
}

func placeQueen(queens *[8]int, column int) {
	if column == len(queens) {
		printQueens(queens)
		return
	}

	row := 1
	for row <= len(queens) {
		if isSafePosition(queens, column, row) {
			queens[column] = row
			placeQueen(queens, column+1)
		}

		row++
	}
}

func isSafePosition(queens *[8]int, column int, row int) bool {
	previousColumn := 0

	for previousColumn < column {
		if queens[previousColumn] == row {
			return false
		}

		rowDifference := queens[previousColumn] - row
		if rowDifference < 0 {
			rowDifference = -rowDifference
		}

		columnDifference := column - previousColumn

		if rowDifference == columnDifference {
			return false
		}

		previousColumn++
	}

	return true
}

func printQueens(queens *[8]int) {
	digits := [8]rune{'1', '2', '3', '4', '5', '6', '7', '8'}
	index := 0

	for index < len(queens) {
		z01.PrintRune(digits[queens[index]-1])
		index++
	}

	z01.PrintRune('\n')
}
```

## المدخلات والمخرجات

```go
func EightQueens()
```

```go
func placeQueen(queens *[8]int, column int)
```

```go
func isSafePosition(queens *[8]int, column int, row int) bool
```

```go
func printQueens(queens *[8]int)
```

الدالة تستقبل القيم حسب التوقيع أعلاه، ثم إما تعيد نتيجة بواسطة `return` أو تعدّل قيمة عبر pointer أو تطبع باستخدام `z01.PrintRune`.

## تدفق التنفيذ خطوة بخطوة

1. إنشاء مصفوفة queens؛ index يمثل العمود والقيمة تمثل الصف.
2. محاولة وضع ملكة في كل صف من العمود الحالي.
3. التحقق من عدم تكرار الصف وعدم وجود تعارض قطري مع الأعمدة السابقة.
4. عند وضع صالح، الانتقال recursively إلى العمود التالي.
5. عند اكتمال 8 أعمدة، طباعة الحل والعودة لاستكشاف حلول أخرى.

## تتبع بأمثلة

- `أحد الأسطر الممكنة` ⟶ `15863724`
- `عدد الحلول المعروف للمسألة` ⟶ `92 سطرًا`

## المفاهيم المهمة

- **recursion**
- **backtracking**
- **المصفوفات والمؤشرات**
- **فحص الأقطار**

## شرح كل سطر

| السطر | التعليمة | الشرح |
|---:|---|---|
| 1 | `package piscine` | يحدد الحزمة التي ينتمي إليها الملف؛ `piscine` تعني أن الدالة جزء من مكتبة التمارين، و`main` تعني برنامجًا قابلًا للتشغيل. |
| 3 | `import "github.com/01-edu/z01"` | يستورد الحزمة المذكورة حتى يمكن استخدام دوالها أو قيمها داخل الملف. |
| 5 | `func EightQueens() {` | تعريف دالة `EightQueens` مع المعاملات ونوع الإرجاع الظاهرين في التوقيع. |
| 6 | `var queens [8]int` | تعليمة Go ضمن الخوارزمية؛ معناها التفصيلي يتضح من قسم تدفق التنفيذ أعلاه ومن المتغيرات المحيطة بها. |
| 8 | `placeQueen(&queens, 0)` | تعليمة Go ضمن الخوارزمية؛ معناها التفصيلي يتضح من قسم تدفق التنفيذ أعلاه ومن المتغيرات المحيطة بها. |
| 9 | `}` | نهاية الكتلة الحالية. |
| 11 | `func placeQueen(queens *[8]int, column int) {` | تعريف دالة `placeQueen` مع المعاملات ونوع الإرجاع الظاهرين في التوقيع. |
| 12 | `if column == len(queens) {` | شرط `if`: ينفذ الكتلة التالية فقط عندما يكون التعبير المنطقي صحيحًا. |
| 13 | `printQueens(queens)` | تعليمة Go ضمن الخوارزمية؛ معناها التفصيلي يتضح من قسم تدفق التنفيذ أعلاه ومن المتغيرات المحيطة بها. |
| 14 | `return` | تنهي تنفيذ الدالة فورًا من دون إرجاع قيمة. |
| 15 | `}` | نهاية الكتلة الحالية. |
| 17 | `row := 1` | إنشاء المتغير `row` لأول مرة واستنتاج نوعه تلقائيًا من القيمة الموجودة يمين `:=`. |
| 18 | `for row <= len(queens) {` | حلقة `for` بعداد أو شرط؛ تتحكم في تكرار الأوامر الموجودة داخلها. |
| 19 | `if isSafePosition(queens, column, row) {` | شرط `if`: ينفذ الكتلة التالية فقط عندما يكون التعبير المنطقي صحيحًا. |
| 20 | `queens[column] = row` | إسناد أو تحديث قيمة `queens[column]` وفق التعبير الموجود في الجهة اليمنى. |
| 21 | `placeQueen(queens, column+1)` | تعليمة Go ضمن الخوارزمية؛ معناها التفصيلي يتضح من قسم تدفق التنفيذ أعلاه ومن المتغيرات المحيطة بها. |
| 22 | `}` | نهاية الكتلة الحالية. |
| 24 | `row++` | يزيد قيمة العداد بمقدار 1. |
| 25 | `}` | نهاية الكتلة الحالية. |
| 26 | `}` | نهاية الكتلة الحالية. |
| 28 | `func isSafePosition(queens *[8]int, column int, row int) bool {` | تعريف دالة `isSafePosition` مع المعاملات ونوع الإرجاع الظاهرين في التوقيع. |
| 29 | `previousColumn := 0` | إنشاء المتغير `previousColumn` لأول مرة واستنتاج نوعه تلقائيًا من القيمة الموجودة يمين `:=`. |
| 31 | `for previousColumn < column {` | حلقة `for` بعداد أو شرط؛ تتحكم في تكرار الأوامر الموجودة داخلها. |
| 32 | `if queens[previousColumn] == row {` | شرط `if`: ينفذ الكتلة التالية فقط عندما يكون التعبير المنطقي صحيحًا. |
| 33 | `return false` | تنهي الدالة فورًا وتعيد القيمة أو التعبير المذكور إلى المستدعي. |
| 34 | `}` | نهاية الكتلة الحالية. |
| 36 | `rowDifference := queens[previousColumn] - row` | إنشاء المتغير `rowDifference` لأول مرة واستنتاج نوعه تلقائيًا من القيمة الموجودة يمين `:=`. |
| 37 | `if rowDifference < 0 {` | شرط `if`: ينفذ الكتلة التالية فقط عندما يكون التعبير المنطقي صحيحًا. |
| 38 | `rowDifference = -rowDifference` | إسناد أو تحديث قيمة `rowDifference` وفق التعبير الموجود في الجهة اليمنى. |
| 39 | `}` | نهاية الكتلة الحالية. |
| 41 | `columnDifference := column - previousColumn` | إنشاء المتغير `columnDifference` لأول مرة واستنتاج نوعه تلقائيًا من القيمة الموجودة يمين `:=`. |
| 43 | `if rowDifference == columnDifference {` | شرط `if`: ينفذ الكتلة التالية فقط عندما يكون التعبير المنطقي صحيحًا. |
| 44 | `return false` | تنهي الدالة فورًا وتعيد القيمة أو التعبير المذكور إلى المستدعي. |
| 45 | `}` | نهاية الكتلة الحالية. |
| 47 | `previousColumn++` | يزيد قيمة العداد بمقدار 1. |
| 48 | `}` | نهاية الكتلة الحالية. |
| 50 | `return true` | تنهي الدالة فورًا وتعيد القيمة أو التعبير المذكور إلى المستدعي. |
| 51 | `}` | نهاية الكتلة الحالية. |
| 53 | `func printQueens(queens *[8]int) {` | تعريف دالة `printQueens` مع المعاملات ونوع الإرجاع الظاهرين في التوقيع. |
| 54 | `digits := [8]rune{'1', '2', '3', '4', '5', '6', '7', '8'}` | إنشاء المتغير `digits` لأول مرة واستنتاج نوعه تلقائيًا من القيمة الموجودة يمين `:=`. |
| 55 | `index := 0` | إنشاء المتغير `index` لأول مرة واستنتاج نوعه تلقائيًا من القيمة الموجودة يمين `:=`. |
| 57 | `for index < len(queens) {` | حلقة `for` بعداد أو شرط؛ تتحكم في تكرار الأوامر الموجودة داخلها. |
| 58 | `z01.PrintRune(digits[queens[index]-1])` | يطبع rune واحدة باستخدام الدالة المسموحة في تمارين piscine. |
| 59 | `index++` | يزيد قيمة العداد بمقدار 1. |
| 60 | `}` | نهاية الكتلة الحالية. |
| 62 | `z01.PrintRune('\n')` | يطبع rune واحدة باستخدام الدالة المسموحة في تمارين piscine. |
| 63 | `}` | نهاية الكتلة الحالية. |

## الحالات الخاصة والأخطاء الشائعة

- الدالة لا تعيد قيمة؛ تطبع كل الحلول.
- عدم تصفير queens بعد الرجوع آمن لأن القيم بعد column الحالي لا تُفحص.
- الصفوف ممثلة من 1 إلى 8.

## التعقيد

البحث أُسي في أسوأ التصور، لكنه مقيد بـ8 أعمدة؛ الذاكرة O(8) للمصفوفة وعمق الاستدعاء.

## طريقة تجربة الدالة

لأن الملف يستخدم `package piscine` فهو ليس برنامجًا مستقلًا. أنشئ ملف `main.go` مؤقتًا في مجلد آخر يستورد الحزمة، أو استخدم ملف اختبار Go. مثال اختبار مبسط:

```go
package piscine

import "testing"

func TestEightQueens(t *testing.T) {
	// ضع هنا حالات الاختبار المناسبة كما في الأمثلة أعلاه.
}
```

> لا تضف ملف `main.go` داخل مجلد التسليم إذا كانت منصة التمرين تطلب ملف الدالة فقط.
