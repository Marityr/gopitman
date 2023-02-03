
tags: #Go 

---

Пакет Go состоит из Go-файлов, расположенных в одной и той же директории, в начале которых имеется одинаковое выражение <code>package</code>. Пакеты, подключаемые к программам, позволяют расширять их возможности. Некоторые пакеты входят в состав стандартной библиотеки Go. А это значит, что они, если вы пользуетесь Go, уже у вас установлены. Другие пакеты устанавливают с помощью команды <code>go get</code>. Можно, кроме того, писать собственные Go-пакеты, создавая файлы в особых директориях, и придерживаясь правил оформления пакетов.

<a href="https://habr.com/ru/company/ruvds/blog/464289/"><img src="https://habrastorage.org/r/w780q1/webt/ht/dv/4c/htdv4cfnn8nmki3yphorht5eiki.jpeg" data-src="https://habrastorage.org/webt/ht/dv/4c/htdv4cfnn8nmki3yphorht5eiki.jpeg" data-blurred="true"/></a>

Материал, перевод которого мы сегодня публикуем, представляет собой руководство по разработке Go-пакетов, которые можно подключать к другим файлам.
<a name="habracut"></a>

## <font color="#3AC1EF">Предварительные требования</font>



- Настройте программное окружение Go (о том, как это сделать, можно узнать <a href="https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-windows-10">здесь</a>). Создайте рабочее пространство Go (этому посвящён пятый пункт вышеупомянутого материала). В следующем разделе этого материала вы сможете найти примеры, которые рекомендуется воспроизвести у себя. Так вы сможете лучше с ними разобраться.

- Для того чтобы углубить свои знания по <code>GOPATH</code> — взгляните на <a href="https://www.digitalocean.com/community/tutorials/understanding-the-gopath">этот</a> материал.




## <font color="#3AC1EF">Написание и импорт пакетов</font>


Написание кода пакета — это то же самое, что и написание обычного кода на Go. Пакеты могут содержать объявления функций, типов и переменных, которые потом могут быть использованы в других Go-программах.

Прежде чем мы сможем приступить к созданию нового пакета, нам нужно перейти в рабочее пространство Go. Оно находится по пути, задаваемом переменной <code>GOPATH</code>. Например, пусть наша организация называется <code>gopherguides</code>. При этом мы, в качестве репозитория, используем <code>GitHub</code>. Это приводит к тому, что у нас, по пути, задаваемом <code>GOPATH</code>, имеется следующая структура папок:


```go
└── $GOPATH
    └── src
        └── github.com
            └── gopherguides
```

Мы собираемся назвать пакет, который будем разрабатывать в этом руководстве, <code>greet</code>. Для того чтобы это сделать — создадим директорию <code>greet</code> в директории <code>gopherguides</code>. В результате имеющаяся структура папок приобретёт следующий вид:


```go
└── $GOPATH
    └── src
        └── github.com
            └── gopherguides
                └── greet
```

Теперь мы готовы к тому, чтобы добавить в директорию <code>greet</code> первый файл. Обычно файл, который является входной точкой (entry point) пакета, называют так же, как названа директория пакета. В данном случае это означает, что мы, в директории <code>greet</code>, создаём файл <code>greet.go</code>:


```go
└── $GOPATH
    └── src
        └── github.com
            └── gopherguides
                └── greet
                    └── greet.go
```

В этом файле мы можем писать код, который хотим многократно использовать в наших проектах. В данном случае мы создадим функцию <code>Hello</code>, которая выводит текст <code>Hello, World!</code>.

Откройте файл <code>greet.go</code> в текстовом редакторе и добавьте в него следующий код:


```go
package greet

import &#34;fmt&#34;

func Hello() {
    fmt.Println(&#34;Hello, World!&#34;)
}
```

Разберём содержимое этого файла. Первая строка каждого файла должна содержать имя пакета, в котором мы работаем. Так как мы находимся в пакете <code>greet</code> — здесь используется ключевое слово <code>package</code>, за которым следует имя пакета:


```go
package greet
```

Это сообщает компилятору о том, что он должен воспринимать всё, что находится в файле, как часть пакета <code>greet</code>.

Далее выполняется импорт необходимых пакетов с помощью выражения <code>import</code>. В данном случае нам нужен всего один пакет — <code>fmt</code>:


```go
import &#34;fmt&#34;
```

И, наконец, мы создаём функцию <code>Hello</code>. Она будет использовать возможности пакета <code>fmt</code> для вывода на экран строки <code>Hello, World!</code>:


```go
func Hello() {
    fmt.Println(&#34;Hello, World!&#34;)
}
```

Теперь, после того, как создан пакет <code>greet</code>, вы можете воспользоваться им в любом другом пакете. Создадим новый пакет, в котором воспользуемся пакетом <code>greet</code>.

А именно, мы создадим пакет <code>example</code>. Для этого будем исходить из тех же предположений, из которых исходили, создавая пакет <code>greet</code>. Для начала создадим папку <code>example</code> в папке <code>gopherguides</code>:


```go
└── $GOPATH
    └── src
        └── github.com
            └── gopherguides
                    └── example
```

Теперь создаём файл, являющийся входной точкой пакета. Данный пакет мы рассматриваем как выполняемую программу, а не как пакет, код которого планируется использовать в других пакетах. Файлы, являющиеся входными точками программ, принято называть <code>main.go</code>:


```go
└── $GOPATH
    └── src
        └── github.com
            └── gopherguides
                └── example
                    └── main.go
```

Откройте в редакторе файл <code>main.go</code> и внесите в него следующий код, который позволяет воспользоваться возможностями пакета <code>greet</code>:


```go
package main

import &#34;github.com/gopherguides/greet&#34;

func main() {
    greet.Hello()
}
```

Мы импортировали в файле <code>main.go</code> пакет <code>greet</code>, а это значит, что для вызова функции, объявленной в этом пакете, нам понадобится воспользоваться точечной нотацией. Точечная нотация — это конструкция, в которой между именем пакета и именем ресурса этого пакета, который нужно использовать, ставится точка. Например, в пакете <code>greet</code> роль ресурса играет функция <code>Hello</code>. Если нужно вызвать эту функцию — используется точечная нотация<code>: greet.Hello()</code>.

Теперь можно открыть терминал и запустить программу:


```go
go run main.go
```

После того, как вы это сделаете, в терминале будет выведено следующее:


```go
Hello, World!
```

Теперь поговорим о том, как использовать переменные, объявляемые в пакетах. Для этого добавим объявление переменной в файл <code>greet.go</code>:


```go
package greet

import &#34;fmt&#34;

var Shark = &#34;Sammy&#34;

func Hello() {
    fmt.Println(&#34;Hello, World!&#34;)
}
```

Откройте файл <code>main.go</code> и добавьте в него строку, в которой функция <code>fmt.Println()</code> используется для вывода значения переменной <code>Shark</code>, объявленной в пакете <code>greet.go</code>. А именно, приведите <code>main.go</code> к следующему виду:


```go
package main

import (
    &#34;fmt&#34;

    &#34;github.com/gopherguides/greet&#34;
)

func main() {
    greet.Hello()

    fmt.Println(greet.Shark)
}
```

Снова запустите программу:


```go
go run main.go
```

Теперь она выведет следующее:


```go
Hello, World!
Sammy
```

А сейчас поговорим о том, как объявлять в пакетах типы. Создадим тип <code>Octopus</code> с полями <code>Name</code> и <code>Color</code>, а также создадим метод типа. Этот метод, при его вызове, будет возвращать особым образом обработанное содержимое полей типа <code>Octopus</code>. Приведём <code>greet.go</code> к следующему виду:


```go
package greet

import &#34;fmt&#34;

var Shark = &#34;Sammy&#34;

type Octopus struct {
    Name  string
    Color string
}

func (o Octopus) String() string {
    return fmt.Sprintf(&#34;The octopus&#39;s name is %q and is the color %s.&#34;, o.Name, o.Color)
}

func Hello() {
    fmt.Println(&#34;Hello, World!&#34;)
}
```

Теперь откройте <code>main.go</code>, создайте в нём экземпляр структуры нового типа и обратитесь к его методу <code>String()</code>:


```go
package main

import (
    &#34;fmt&#34;

    &#34;github.com/gopherguides/greet&#34;
)

func main() {
    greet.Hello()

    fmt.Println(greet.Shark)

    oct := greet.Octopus{
        Name:  &#34;Jesse&#34;,
        Color: &#34;orange&#34;,
    }

    fmt.Println(oct.String())
}
```

После того, как вы, с помощью конструкции, которая выглядит как <code>oct := greet.Octopus</code>, создали экземпляр <code>Octopus</code>, вы можете обращаться к методам и свойствам типа из пространства имён файла <code>main.go</code>. Это, в частности, позволяет воспользоваться командой <code>oct.String()</code>, расположенной в конце файла <code>main.go</code>, не обращаясь к <code>greet</code>. Кроме того, мы можем, например, обратиться к полю структуры <code>Color</code>, воспользовавшись конструкцией <code>oct.Color</code>. При этом мы, как и тогда, когда вызывали метод, не обращаемся к <code>greet</code>.

Метод <code>String</code> типа <code>Octopus</code> использует функцию <code>fmt.Sprintf</code> для формирования предложения и возвращает, с помощью <code>return</code>, результат, строку, в место вызова метода (в данном случае это место находится в <code>main.go</code>).

Запустим программу снова:


```go
go run main.go
```

Она выведет в консоль следующее:


```go
Hello, World!
Sammy
The octopus&#39;s name is &#34;Jesse&#34; and is the color orange.
```

Теперь, когда мы оснастили Octopus методом <code>String</code>, мы получили механизм вывода сведений о типе, подходящий для многократного использования. Если в будущем понадобится изменить поведение этого метода, который может использоваться во многих проектах, достаточно будет один раз отредактировать его код в <code>greet.go</code>.


## <font color="#3AC1EF">Экспорт сущностей</font>


Возможно, вы обратили внимание на то, что всё, с чем мы работали, обращаясь к пакету <code>greet</code>, имеет имена, начинающиеся с прописной буквы. В Go нет модификаторов доступа наподобие <code>public</code>, <code>private</code> или <code>protected</code>, которые есть в других языках. Видимость сущностей для внешних механизмов контролируется тем, с какой буквы, с маленькой или с большой, начинаются их имена. В результате типы, переменные, функции, имена которых начинаются с прописной буквы, доступны за пределами текущего пакета. Код, который виден за пределами пакета, называется экспортированным.

Если оснастить тип <code>Octopus</code> новым методом с именем <code>reset</code>, то этот метод можно будет вызывать из пакета <code>greet</code>, но не из файла <code>main.go</code>, который находится за пределами пакета <code>greet</code>. Вот обновлённый вариант <code>greet.go</code>:


```go
package greet

import &#34;fmt&#34;

var Shark = &#34;Sammy&#34;

type Octopus struct {
    Name  string
    Color string
}

func (o Octopus) String() string {
    return fmt.Sprintf(&#34;The octopus&#39;s name is %q and is the color %s.&#34;, o.Name, o.Color)
}

func (o Octopus) reset() {
    o.Name = &#34;&#34;
    o.Color = &#34;&#34;
}

func Hello() {
    fmt.Println(&#34;Hello, World!&#34;)
}
```

Попытаемся вызвать <code>reset</code> из файла <code>main.go</code>:


```go
package main

import (
    &#34;fmt&#34;

    &#34;github.com/gopherguides/greet&#34;
)

func main() {
    greet.Hello()

    fmt.Println(greet.Shark)

    oct := greet.Octopus{
        Name:  &#34;Jesse&#34;,
        Color: &#34;orange&#34;,
    }

    fmt.Println(oct.String())

    oct.reset()
}
```

Это приведёт к появлению следующей ошибки компиляции:


```go
oct.reset undefined (cannot refer to unexported field or method greet.Octopus.reset)
```

Для того чтобы экспортировать метод <code>reset</code> типа <code>Octopus</code> нужно его переименовать, заменив первую букву, строчную <code>r</code>, на прописную <code>R</code>. Сделаем это, отредактировав <code>greet.go</code>:


```go
package greet

import &#34;fmt&#34;

var Shark = &#34;Sammy&#34;

type Octopus struct {
    Name  string
    Color string
}

func (o Octopus) String() string {
    return fmt.Sprintf(&#34;The octopus&#39;s name is %q and is the color %s.&#34;, o.Name, o.Color)
}

func (o Octopus) Reset() {
    o.Name = &#34;&#34;
    o.Color = &#34;&#34;
}

func Hello() {
    fmt.Println(&#34;Hello, World!&#34;)
}
```

Это приведёт к тому, что мы сможем вызывать <code>Reset</code> из других пакетов и при этом не сталкиваться с сообщениями об ошибках:


```go
package main

import (
    &#34;fmt&#34;

    &#34;github.com/gopherguides/greet&#34;
)

func main() {
    greet.Hello()

    fmt.Println(greet.Shark)

    oct := greet.Octopus{
        Name:  &#34;Jesse&#34;,
        Color: &#34;orange&#34;,
    }

    fmt.Println(oct.String())

    oct.Reset()

    fmt.Println(oct.String())
}
```

Запустим программу:


```go
go run main.go
```

Вот что попадёт в консоль:


```go
Hello, World!
Sammy
The octopus&#39;s name is &#34;Jesse&#34; and is the color orange
The octopus&#39;s name is &#34;&#34; and is the color .
```

Вызвав метод <code>Reset</code>, мы очистили поля <code>Name</code> и <code>Color</code> нашего экземпляра <code>Octopus</code>. В результате, при вызове <code>String</code>, там, где раньше выводилось содержимое полей <code>Name</code> и <code>Color</code>, теперь не выводится ничего.


## <font color="#3AC1EF">Итоги</font>


Написание пакетов Go ничем не отличается от написания обычного Go-кода. Однако размещение кода пакетов в собственных директориях позволяет изолировать код, которым можно воспользоваться в любых других Go-проектах. Здесь мы поговорили о том, как объявлять в пакетах функции, переменные и типы, рассмотрели порядок использования этих сущностей за пределами пакетов, разобрались с тем, где нужно хранить пакеты, рассчитанные на их многократное использование.

<b>Уважаемые читатели!</b> Какие программы вы обычно пишете на Go? Пользуетесь ли вы в них пакетами собственной разработки?

<a href="https://ruvds.com/vps_start/"><img src="https://habrastorage.org/r/w1560/webt/it/t5/3p/itt53pns2iucwylb3bwn1fmmtnu.png" data-src="https://habrastorage.org/webt/it/t5/3p/itt53pns2iucwylb3bwn1fmmtnu.png"/></a>