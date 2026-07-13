# تمارين Level 13 مع شرح عربي

هذا الدليل يجمع تمارين Level 13 التي تمت إضافتها إلى المشروع. كل تمرين موجود في مجلد مستقل، وداخل كل مجلد:

- `main.go`: ملف الحل القابل للتسليم.
- `README.md`: شرح عربي مفصل ودقيق.

## التمارين

| التمرين | الفكرة |
|---|---|
| [`printprogramname`](./printprogramname/README.md) | طباعة اسم الملف التنفيذي مع دعم Unicode |
| [`printparams`](./printparams/README.md) | طباعة arguments بالترتيب، كل واحد في سطر |
| [`revparams`](./revparams/README.md) | طباعة arguments بترتيب عكسي |
| [`sortparams`](./sortparams/README.md) | ترتيب arguments وفق ASCII باستخدام Bubble Sort |
| [`nbrconvertalpha`](./nbrconvertalpha/README.md) | تحويل الأرقام من 1 إلى 26 إلى حروف مع `--upper` |
| [`flags`](./flags/README.md) | معالجة `--insert` و`--order` و`--help` |
| [`rotatevowels`](./rotatevowels/README.md) | عكس ترتيب حروف العلة عبر جميع arguments |

## تنسيق الملفات قبل التسليم

نفّذ من جذر المشروع:

```bash
gofmt -w printprogramname/main.go \
  printparams/main.go \
  revparams/main.go \
  sortparams/main.go \
  nbrconvertalpha/main.go \
  flags/main.go \
  rotatevowels/main.go
```

ثم افحص وجود newline في نهاية الملفات:

```bash
for file in printprogramname/main.go printparams/main.go revparams/main.go sortparams/main.go nbrconvertalpha/main.go flags/main.go rotatevowels/main.go; do
  printf "%s: " "$file"
  tail -c 1 "$file" | od -An -t u1
done
```

يجب أن تكون القيمة الأخيرة لكل ملف:

```text
10
```

وهي قيمة newline في ASCII.

## ملاحظة مهمة

ملفات `README.md` للشرح فقط. منصة التقييم عادة تحدد ملف التسليم المطلوب، مثل:

```text
rotatevowels/main.go
```

لذلك لا تغيّر اسم `main.go` ولا تنقل الكود خارج مجلد التمرين عند التسليم.
