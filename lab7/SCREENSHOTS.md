# Что заскриншотить для лабораторной №7

1. Проверка Ubuntu-контейнера и инструментов:
   - команда `docker run --rm ubuntu:26.04 cat /etc/os-release`;
   - раздел `1. Ubuntu и установленные инструменты` в выводе `bash scripts/run_in_docker.sh`.

2. AST:
   - раздел `3. AST общего задания`;
   - раздел `6. Индивидуальный вариант 2.10: функции и inline` со строками `FunctionDecl`, `CallExpr`, `BinaryOperator`.

3. LLVM IR и оптимизации:
   - раздел `4. LLVM IR общего задания: -O0, -O2 и opt`;
   - фрагменты `general_O0.ll`, где видны `alloca`, `store`, `load`, `call square`;
   - фрагменты `general_O2.ll`, где видно сокращение IR после оптимизации;
   - diff `artifacts/diff/general_O0_vs_O2.diff`.

4. Индивидуальное задание:
   - раздел `6`, где в `variant_O0.ll` виден вызов `_Z6squarei`;
   - раздел `6`, где в `variant_O2.ll` видно `ret i32 25`;
   - раздел `7. opt -passes=always-inline`, где видно, что при `always_inline` отдельного вызова уже нет, а после полного `O2` остается `ret i32 25`.

5. CFG:
   - разделы `5. CFG общего задания` и `8. CFG индивидуального варианта`;
   - видимые PNG-файлы `cfg_main.png`, `cfg_square.png`, `cfg__Z6squarei.png` из каталогов `artifacts/cfg/general_O0`, `artifacts/cfg/general_O2`, `artifacts/cfg/variant_O0`, `artifacts/cfg/variant_O2`, `artifacts/cfg/variant_always_inline`.

6. Итог:
   - раздел `10. Краткие выводы для скриншота`;
   - файл `README.md` как готовую текстовую часть отчёта.