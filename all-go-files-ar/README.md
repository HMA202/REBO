# جميع ملفات Go مع شرح عربي

هذه نسخة تعليمية منظمة من المشروع. **المشروع الأصلي بقي في مكانه** حتى لا تتغير بنية ملفات التسليم. داخل هذا المجلد توجد نسخة من كل ملف Go موضوعة في مجلد مستقل مع `README.md` عربي.

## الأقسام

- `piscine-functions/`: جميع ملفات `.go` التي كانت في جذر المشروع، وعددها 47.
- `programs/`: برامج `package main` مثل printparams وflags وrotatevowels، إضافة إلى مشروع quad.
- `practice-exercises/`: جميع التمارين الإضافية من 01 إلى 65، مع `exercise.go` و`main.go` وشرح مستقل.

## دوال piscine الأساسية

- [`alphacount`](piscine-functions/alphacount/README.md) — عدّ الحروف الإنجليزية
- [`atoi`](piscine-functions/atoi/README.md) — تحويل string إلى int مع الإشارة
- [`atoibase`](piscine-functions/atoibase/README.md) — تحويل عدد مكتوب بقاعدة مخصصة إلى int
- [`basicatoi`](piscine-functions/basicatoi/README.md) — تحويل رقمي بسيط من string إلى int
- [`basicatoi2`](piscine-functions/basicatoi2/README.md) — تحويل string رقمي مع التحقق
- [`basicjoin`](piscine-functions/basicjoin/README.md) — دمج عناصر []string من دون فاصل
- [`capitalize`](piscine-functions/capitalize/README.md) — تكبير أول حرف من كل كلمة وتصغير البقية
- [`compare`](piscine-functions/compare/README.md) — مقارنة نصين
- [`concat`](piscine-functions/concat/README.md) — دمج نصين
- [`divmod`](piscine-functions/divmod/README.md) — القسمة والباقي باستخدام المؤشرات
- [`eightqueens`](piscine-functions/eightqueens/README.md) — حل مسألة الملكات الثماني بالـ backtracking
- [`fibonacci`](piscine-functions/fibonacci/README.md) — حساب فيبوناتشي بالاستدعاء الذاتي
- [`findnextprime`](piscine-functions/findnextprime/README.md) — إيجاد أول عدد أولي أكبر من أو يساوي العدد
- [`firstrune`](piscine-functions/firstrune/README.md) — إرجاع أول rune من النص
- [`index`](piscine-functions/index/README.md) — البحث عن نص فرعي
- [`isalpha`](piscine-functions/isalpha/README.md) — التحقق من أن النص أبجدي رقمي
- [`islower`](piscine-functions/islower/README.md) — التحقق من الحروف الصغيرة
- [`isnegative`](piscine-functions/isnegative/README.md) — طباعة هل العدد سالب
- [`isnumeric`](piscine-functions/isnumeric/README.md) — التحقق من أن النص أرقام فقط
- [`isprime`](piscine-functions/isprime/README.md) — فحص العدد الأولي
- [`isprintable`](piscine-functions/isprintable/README.md) — التحقق من محارف ASCII القابلة للطباعة
- [`isupper`](piscine-functions/isupper/README.md) — التحقق من الحروف الكبيرة
- [`iterativefactorial`](piscine-functions/iterativefactorial/README.md) — حساب العامل factorial بحلقة
- [`iterativepower`](piscine-functions/iterativepower/README.md) — حساب القوة بحلقة
- [`join`](piscine-functions/join/README.md) — دمج النصوص مع فاصل
- [`lastrune`](piscine-functions/lastrune/README.md) — إرجاع آخر rune
- [`nrune`](piscine-functions/nrune/README.md) — إرجاع الحرف رقم n
- [`pointone`](piscine-functions/pointone/README.md) — تعديل قيمة عبر pointer
- [`printcomb`](piscine-functions/printcomb/README.md) — طباعة جميع تركيبات ثلاثة أرقام متزايدة
- [`printcomb2`](piscine-functions/printcomb2/README.md) — طباعة أزواج الأعداد من 00 إلى 99
- [`printcombn`](piscine-functions/printcombn/README.md) — طباعة تركيبات n أرقام متزايدة
- [`printnbr`](piscine-functions/printnbr/README.md) — طباعة int باستخدام PrintRune فقط
- [`printnbrbase`](piscine-functions/printnbrbase/README.md) — طباعة عدد بقاعدة مخصصة
- [`printnbrinorder`](piscine-functions/printnbrinorder/README.md) — طباعة خانات العدد مرتبة تصاعديًا
- [`printstr`](piscine-functions/printstr/README.md) — طباعة string حرفًا حرفًا
- [`recursivefactorial`](piscine-functions/recursivefactorial/README.md) — حساب العامل بالاستدعاء الذاتي
- [`recursivepower`](piscine-functions/recursivepower/README.md) — حساب القوة بالاستدعاء الذاتي
- [`sortintegertable`](piscine-functions/sortintegertable/README.md) — ترتيب slice أعداد تصاعديًا
- [`sqrt`](piscine-functions/sqrt/README.md) — الجذر التربيعي الصحيح
- [`strlen`](piscine-functions/strlen/README.md) — حساب عدد حروف Unicode
- [`strrev`](piscine-functions/strrev/README.md) — عكس النص مع دعم Unicode
- [`swap`](piscine-functions/swap/README.md) — تبديل قيمتين باستخدام المؤشرات
- [`tolower`](piscine-functions/tolower/README.md) — تحويل الحروف الإنجليزية إلى صغيرة
- [`toupper`](piscine-functions/toupper/README.md) — تحويل الحروف الإنجليزية إلى كبيرة
- [`trimatoi`](piscine-functions/trimatoi/README.md) — استخراج الأرقام من نص وبناء عدد
- [`ultimatedivmod`](piscine-functions/ultimatedivmod/README.md) — القسمة والباقي داخل نفس المتغيرين
- [`ultimatepointone`](piscine-functions/ultimatepointone/README.md) — تعديل int عبر مؤشر ثلاثي

## البرامج

- [`printalphabet`](programs/printalphabet/README.md)
- [`printdigits`](programs/printdigits/README.md)
- [`printreversealphabet`](programs/printreversealphabet/README.md)
- [`printprogramname`](programs/printprogramname/README.md)
- [`printparams`](programs/printparams/README.md)
- [`revparams`](programs/revparams/README.md)
- [`sortparams`](programs/sortparams/README.md)
- [`nbrconvertalpha`](programs/nbrconvertalpha/README.md)
- [`flags`](programs/flags/README.md)
- [`rotatevowels`](programs/rotatevowels/README.md)
- [`quad`](programs/quad/README.md)

## التمارين الإضافية

- [`01_countlen`](practice-exercises/01_countlen/README.md)
- [`02_countchar`](practice-exercises/02_countchar/README.md)
- [`03_countvowels`](practice-exercises/03_countvowels/README.md)
- [`04_haschar`](practice-exercises/04_haschar/README.md)
- [`05_firstchar`](practice-exercises/05_firstchar/README.md)
- [`06_lastchar`](practice-exercises/06_lastchar/README.md)
- [`07_reversestring`](practice-exercises/07_reversestring/README.md)
- [`08_ispalindrome`](practice-exercises/08_ispalindrome/README.md)
- [`09_iseven`](practice-exercises/09_iseven/README.md)
- [`10_isodd`](practice-exercises/10_isodd/README.md)
- [`11_max`](practice-exercises/11_max/README.md)
- [`12_min`](practice-exercises/12_min/README.md)
- [`13_abs`](practice-exercises/13_abs/README.md)
- [`14_factorial`](practice-exercises/14_factorial/README.md)
- [`15_sumslice`](practice-exercises/15_sumslice/README.md)
- [`16_maxslice`](practice-exercises/16_maxslice/README.md)
- [`17_minslice`](practice-exercises/17_minslice/README.md)
- [`18_countpositive`](practice-exercises/18_countpositive/README.md)
- [`19_reverseslice`](practice-exercises/19_reverseslice/README.md)
- [`20_contains`](practice-exercises/20_contains/README.md)
- [`21_printdigits`](practice-exercises/21_printdigits/README.md)
- [`22_printreversedigits`](practice-exercises/22_printreversedigits/README.md)
- [`23_printevendigits`](practice-exercises/23_printevendigits/README.md)
- [`24_printodddigits`](practice-exercises/24_printodddigits/README.md)
- [`25_printsquare`](practice-exercises/25_printsquare/README.md)
- [`26_run_length_encode`](practice-exercises/26_run_length_encode/README.md)
- [`27_countwords`](practice-exercises/27_countwords/README.md)
- [`28_removespaces`](practice-exercises/28_removespaces/README.md)
- [`29_onlydigits`](practice-exercises/29_onlydigits/README.md)
- [`30_touppersimple`](practice-exercises/30_touppersimple/README.md)
- [`31_greatest_common_divisor_uint`](practice-exercises/31_greatest_common_divisor_uint/README.md)
- [`32_least_common_multiple_uint`](practice-exercises/32_least_common_multiple_uint/README.md)
- [`33_is_divisible_by_both`](practice-exercises/33_is_divisible_by_both/README.md)
- [`34_next_divisible_by_both`](practice-exercises/34_next_divisible_by_both/README.md)
- [`35_string_to_int`](practice-exercises/35_string_to_int/README.md)
- [`36_string_to_uint`](practice-exercises/36_string_to_uint/README.md)
- [`37_int_to_string`](practice-exercises/37_int_to_string/README.md)
- [`38_parse_bool`](practice-exercises/38_parse_bool/README.md)
- [`39_bool_to_string`](practice-exercises/39_bool_to_string/README.md)
- [`40_binary_to_decimal`](practice-exercises/40_binary_to_decimal/README.md)
- [`41_decimal_to_binary`](practice-exercises/41_decimal_to_binary/README.md)
- [`42_hex_digit_value`](practice-exercises/42_hex_digit_value/README.md)
- [`43_sum_digits_in_string`](practice-exercises/43_sum_digits_in_string/README.md)
- [`44_count_upper_lower`](practice-exercises/44_count_upper_lower/README.md)
- [`45_string_to_float_simple`](practice-exercises/45_string_to_float_simple/README.md)
- [`46_sudoku_is_valid`](practice-exercises/46_sudoku_is_valid/README.md)
- [`47_sudoku_find_empty`](practice-exercises/47_sudoku_find_empty/README.md)
- [`48_sudoku_solve`](practice-exercises/48_sudoku_solve/README.md)
- [`49_maze_shortest_path`](practice-exercises/49_maze_shortest_path/README.md)
- [`50_simple_calculator`](practice-exercises/50_simple_calculator/README.md)
- [`51_calculator_with_multiply`](practice-exercises/51_calculator_with_multiply/README.md)
- [`52_csv_parse_line`](practice-exercises/52_csv_parse_line/README.md)
- [`53_roman_to_int`](practice-exercises/53_roman_to_int/README.md)
- [`54_int_to_roman`](practice-exercises/54_int_to_roman/README.md)
- [`55_base_convert`](practice-exercises/55_base_convert/README.md)
- [`56_balanced_brackets`](practice-exercises/56_balanced_brackets/README.md)
- [`57_longest_unique_substring`](practice-exercises/57_longest_unique_substring/README.md)
- [`58_merge_sort`](practice-exercises/58_merge_sort/README.md)
- [`59_binary_search`](practice-exercises/59_binary_search/README.md)
- [`60_rotate_matrix`](practice-exercises/60_rotate_matrix/README.md)
- [`61_word_frequency`](practice-exercises/61_word_frequency/README.md)
- [`62_group_anagrams`](practice-exercises/62_group_anagrams/README.md)
- [`63_time_to_seconds`](practice-exercises/63_time_to_seconds/README.md)
- [`64_caesar_cipher`](practice-exercises/64_caesar_cipher/README.md)
- [`65_tic_tac_toe_winner`](practice-exercises/65_tic_tac_toe_winner/README.md)
