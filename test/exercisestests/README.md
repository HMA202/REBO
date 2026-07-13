# Go Exercises Tests - Clean Reordered

This version removes duplicated exercises and renumbers the remaining exercises from 01 to 65.

## Removed duplicated folders

- `26_countrepeatedletters`
- `43_only_digits`
- `47_count_words`
- `48_remove_spaces`
- `50_count_char`

## Notes

- `49_run_length_encode` was kept and moved to number 26 because it is the cleaner version of the repeated-letters exercise.
- `26_countrepeatedletters` was removed because it duplicates the same idea but prints directly instead of returning a string.
- `43_only_digits`, `47_count_words`, `48_remove_spaces`, and `50_count_char` were removed because earlier versions already exist.

## Exercises

01. `01_countlen`
02. `02_countchar`
03. `03_countvowels`
04. `04_haschar`
05. `05_firstchar`
06. `06_lastchar`
07. `07_reversestring`
08. `08_ispalindrome`
09. `09_iseven`
10. `10_isodd`
11. `11_max`
12. `12_min`
13. `13_abs`
14. `14_factorial`
15. `15_sumslice`
16. `16_maxslice`
17. `17_minslice`
18. `18_countpositive`
19. `19_reverseslice`
20. `20_contains`
21. `21_printdigits`
22. `22_printreversedigits`
23. `23_printevendigits`
24. `24_printodddigits`
25. `25_printsquare`
26. `26_run_length_encode`
27. `27_countwords`
28. `28_removespaces`
29. `29_onlydigits`
30. `30_touppersimple`
31. `31_greatest_common_divisor_uint`
32. `32_least_common_multiple_uint`
33. `33_is_divisible_by_both`
34. `34_next_divisible_by_both`
35. `35_string_to_int`
36. `36_string_to_uint`
37. `37_int_to_string`
38. `38_parse_bool`
39. `39_bool_to_string`
40. `40_binary_to_decimal`
41. `41_decimal_to_binary`
42. `42_hex_digit_value`
43. `43_sum_digits_in_string`
44. `44_count_upper_lower`
45. `45_string_to_float_simple`
46. `46_sudoku_is_valid`
47. `47_sudoku_find_empty`
48. `48_sudoku_solve`
49. `49_maze_shortest_path`
50. `50_simple_calculator`
51. `51_calculator_with_multiply`
52. `52_csv_parse_line`
53. `53_roman_to_int`
54. `54_int_to_roman`
55. `55_base_convert`
56. `56_balanced_brackets`
57. `57_longest_unique_substring`
58. `58_merge_sort`
59. `59_binary_search`
60. `60_rotate_matrix`
61. `61_word_frequency`
62. `62_group_anagrams`
63. `63_time_to_seconds`
64. `64_caesar_cipher`
65. `65_tic_tac_toe_winner`

## How to run

```bash
cd 01_countlen
go run .
```

Each exercise folder contains:

```txt
exercise.go
main.go
README.md
```

For Piscine submission, copy the function from `exercise.go` and change:

```go
package main
```

to:

```go
package piscine
```
