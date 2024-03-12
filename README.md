# Unguess

**Unguess**&nbsp;&mdash; это CLI-приложение для создания надежных паролей с максимальной криптоустойчивостью. Внутренний генератор использует особенности неопределенности конкурентности языка программирования Go, такие как вызовы горутин в непоследовательном порядке.

### Примеры паролей:

```
ORv0AZHd3EU#pxC1n6tM74qGaefjh@uz
Oz47yY0uQgcjP#liRF1qmaJ8GoXhLx9H
0TX$e7cw3NQbJ&Y#RO1%2DKCqzrdvAUG
dV$wzs%8PqWUeSQk6iagK&ZouIG3D1yM
tld70vZGz3LK$a4f9hco#kgPyCbTB8@D
wFe2Vd$kTmP1cv#rM5S8qUjRBtAHEYzo
&2bBODaj1W$Qg6t984pRCq%zi#ElcnZY
LPDCW0n5J#jX@seO&bV3muRIkBMthpZY
xDKfalyMkJoihB&X$NmO1d4@rGWb#2H5
C9Ar#Hj$7S3OcYGmDdl@kZXQqLVWxaRI
wTgRrW5%SG2k71U8VHu4Fv$sm0ob&E3N
9cqhisLvgFaEd6$&l1KZ70DwNWjpPUe%
RLus7APqbyw0ZYg$dUVK3aefJTBHxi8W
5jZu17JFnAMR$x%#eNpDcUtgQP@HOws&
```

# Код для импортирования

В коде этой программы есть также открытый пакет генератора `github.com/mrumyantsev/unguess/pkg/gen`, который может быть импортирован в вашем приложении.

# Требования

### Установленное ПО:

- Средства разработки языка Go &geq;v1.20 (только для локальной сборки)
- Утилита make

# Применение

Для **сборки** приложения **локально**, выполните команду в терминале:

```
make build
```

Для **запуска** приложения **локально**, выполните:

```
make run
```

Чтобы создавать пароли изменяемой длины, необходимо запустить исполняемый файл программы, передав ему параметр целого числа, например такой командой:

```
./build/unguess 67
```

Чтобы запустить приложение в режиме **быстрого запуска** (например, во время разработки), выполните команду:

```
make fast-run
```

Или выполните ее короткую версию:

```
make
```
